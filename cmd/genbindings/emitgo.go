package main

import (
	"C"
	"go/format"
	"log"
	"sort"
	"strings"
)

func goReservedWord(s string) bool {
	switch s {
	case "default", "const", "func", "var", "type", "len", "new", "copy", "import":
		return true
	default:
		return false
	}
}

func (p CppParameter) RenderTypeGo() string {
	if p.Pointer && p.ParameterType == "char" {
		return "string"
	}
	if p.ParameterType == "char" {
		return "byte"
	}
	if p.ParameterType == "QString" {
		return "string"
	}
	if p.ParameterType == "uintptr_t" {
		return "uintptr"
	}

	if t, ok := p.QListOf(); ok {
		return "[]" + t.RenderTypeGo()
	}

	ret := ""
	if p.ByRef || p.Pointer {
		ret += "*"
	}

	switch p.ParameterType {
	case "char", "qint8":
		ret += "byte"
	case "unsigned char", "quint8":
		ret += "byte"
	case "short", "qint16":
		ret += "int16"
	case "ushort", "quint16":
		ret += "uint16"
	case "long":
		// Windows ILP32 - 32-bits
		// Linux LP64 - 64-bits
		if C.sizeof_long == 4 {
			ret += "int32"
		} else {
			ret += "int64"
		}
	case "ulong", "unsigned long":
		if C.sizeof_long == 4 {
			ret += "uint32"
		} else {
			ret += "uint64"
		}
	case "QRgb":
		if C.sizeof_int == 4 {
			ret += "uint32"
		} else {
			ret += "uint64"
		}

	case "qint32":
		ret += "int32"
	case "quint32":
		ret += "uint32"
	case "qlonglong", "qint64":
		ret += "int64"
	case "qulonglong", "quint64":
		ret += "uint64"
	case "float":
		ret += "float32"
	case "double":
		ret += "float64"
	default:
		// Do not transform this type
		ret += p.ParameterType

	}

	return ret // ignore const
}

func (p CppParameter) parameterTypeCgo() string {
	tmp := strings.Replace(p.RenderTypeCpp(), `*`, "", -1)
	if strings.HasPrefix(tmp, "unsigned ") {
		tmp = "u" + tmp[9:] // Cgo uses uchar, uint instead of full name
	}
	return "C." + strings.Replace(tmp, " ", "_", -1)
}

func emitParametersGo(params []CppParameter) string {
	tmp := make([]string, 0, len(params))
	for _, p := range params {
		tmp = append(tmp, p.ParameterName+" "+p.RenderTypeGo())
	}
	return strings.Join(tmp, ", ")
}

func emitParametersGo2CABIForwarding(m CppMethod) (preamble string, fowarding string) {
	tmp := make([]string, 0, len(m.Parameters)+2)

	tmp = append(tmp, "this.h")

	for _, p := range m.Parameters {
		if p.ParameterType == "QString" {
			// Go: convert string -> char* and len
			// CABI: convert char* and len -> real QString
			preamble += p.ParameterName + "_Cstring := C.CString(" + p.ParameterName + ")\n"
			preamble += "defer C.free(" + p.ParameterName + "_Cstring)\n"
			tmp = append(tmp, p.ParameterName+"_Cstring, len("+p.ParameterName+")")

		} else if listType, ok := p.QListOf(); ok {
			// QList<T>
			// Go: convert T[] -> t* and len
			// CABI: create a real QList<>

			if listType.ParameterType == "QString" {
				// Combo

				preamble += "// For the C ABI, malloc two C arrays; raw char* pointers and their lengths\n"
				preamble += p.ParameterName + "_CArray := (*[0xffff]*" + listType.parameterTypeCgo() + ")(C.malloc(c.ulong(8 * len(" + p.ParameterName + "))))\n"
				preamble += p.ParameterName + "_Lengths := (*[0xffff]*C.size_t)(C.malloc(c.ulong(8 * len(" + p.ParameterName + "))))\n"
				preamble += "defer C.free(" + p.ParameterName + "_CArray)\n"
				preamble += "defer C.free(" + p.ParameterName + "_Lengths)\n"
				preamble += "for i := range " + p.ParameterName + "{\n"
				preamble += "single_cstring := C.CString(" + p.ParameterName + "[i])\n"
				preamble += "defer C.free(single_cstring)\n"
				preamble += p.ParameterName + "_CArray[i] = single_cstring\n"
				preamble += p.ParameterName + "__Lengths[i] = len(" + p.ParameterName + "[i])\n"
				preamble += "}\n"

				tmp = append(tmp, p.ParameterName+"_CArray, "+p.ParameterName+"_Lengths, len("+p.ParameterName+")")

			} else {

				preamble += "// For the C ABI, malloc a C array of raw pointers\n"
				preamble += p.ParameterName + "_CArray := (*[0xffff]*" + listType.parameterTypeCgo() + ")(C.malloc(c.ulong(8 * len(" + p.ParameterName + "))))\n"
				preamble += "defer C.free(" + p.ParameterName + "_CArray)\n"
				preamble += "for i := range " + p.ParameterName + "{\n"
				preamble += p.ParameterName + "_CArray[i] = " + p.ParameterName + "[i].cPointer()\n"
				preamble += "}\n"

				tmp = append(tmp, p.ParameterName+"_CArray, len("+p.ParameterName+")")
			}

		} else if p.Pointer && p.ParameterType == "char" {
			// Single char* argument
			preamble += p.ParameterName + "_Cstring := C.CString(" + p.ParameterName + ")\n"
			preamble += "defer C.free(" + p.ParameterName + "_Cstring)\n"
			tmp = append(tmp, p.ParameterName+"_Cstring")

		} else if (p.Pointer || p.ByRef) && p.QtClassType() {
			// The C++ type is a pointer to Qt class
			// We want our functions to accept the Go wrapper type, and forward as cPointer()
			tmp = append(tmp, p.ParameterName+".cPointer()")

		} else {
			// Default
			tmp = append(tmp, p.ParameterName)
		}
	}

	if m.ReturnType.ParameterType == "QString" {
		tmp = append(tmp, "&_out, &_out_Strlen")
	} else if _, ok := m.ReturnType.QListOf(); ok {
		tmp = append(tmp, "&_out, &_out_len")
	}

	return preamble, strings.Join(tmp, ", ")
}

func emitGo(src *CppParsedHeader, headerName string) (string, error) {

	ret := strings.Builder{}
	ret.WriteString(`package miqt

/*

#cgo CFLAGS: -fPIC
#cgo pkg-config: Qt5Widgets
#include "gen_` + headerName + `"
#include <stdlib.h>

*/
import "C"

%%_IMPORTLIBS_%%
`)

	imports := map[string]struct{}{}

	for _, c := range src.Classes {

		ret.WriteString(`
		type ` + c.ClassName + ` struct {
			h *C.` + c.ClassName + `
		`)

		// Embed all inherited types to directly allow calling inherited methods
		for _, base := range c.Inherits {
			ret.WriteString(base + "\n")
		}

		ret.WriteString(`
		}
		
		func (this *` + c.ClassName + `) cPointer() *C.` + c.ClassName + ` {
			if this == nil {
				return nil
			}
			return this.h
		}
		
		`)

		localInit := "h: ret"
		for _, base := range c.Inherits {
			localInit += ", " + base + ": " + base + "{h: ret}"
		}

		ret.WriteString(`
			func new` + c.ClassName + `(h *C.` + c.ClassName + `) {
				return &` + c.ClassName + `{` + localInit + `}
			}
			
		`)

		for i, ctor := range c.Ctors {
			preamble, forwarding := emitParametersGo2CABIForwarding(ctor)
			ret.WriteString(`
			// New` + c.ClassName + maybeSuffix(i) + ` constructs a new ` + c.ClassName + ` object.
			func New` + c.ClassName + maybeSuffix(i) + `(` + emitParametersGo(ctor.Parameters) + `) {
				` + preamble + ` ret := C.` + c.ClassName + `_new` + maybeSuffix(i) + `(` + forwarding + `)
				return new` + c.ClassName + `(h)
			}
			
			`)
		}

		for _, m := range c.Methods {
			preamble, forwarding := emitParametersGo2CABIForwarding(m)

			shouldReturn := "return "
			afterword := ""
			returnTypeDecl := m.ReturnType.RenderTypeGo() // FIXME handle byRef/const here too

			if m.ReturnType.ParameterType == "void" && !m.ReturnType.Pointer {
				shouldReturn = ""
				returnTypeDecl = ""

			} else if m.ReturnType.ParameterType == "void" && m.ReturnType.Pointer {
				returnTypeDecl = "interface{}"

			} else if m.ReturnType.ParameterType == "char" && m.ReturnType.Pointer {
				// Qt functions normally return QString - anything returning char*
				// is something like QByteArray.Data() where it returns an unsafe
				// internal pointer
				imports["unsafe"] = struct{}{}
				returnTypeDecl = "unsafe.Pointer"

			} else if m.ReturnType.ParameterType == "QString" {
				shouldReturn = ""
				returnTypeDecl = "string"

				preamble += "var _out *C.char = nil\n"
				preamble += "var _out_Strlen C.size_t = 0\n"
				afterword += "ret := C.GoStringN(_out, _out_Strlen)\n"
				afterword += "C.free(_out)\n"
				afterword += "return ret"

			} else if t, ok := m.ReturnType.QListOf(); ok {
				shouldReturn = ""
				returnTypeDecl = "[]" + t.RenderTypeGo()

				preamble += "var _out **" + t.parameterTypeCgo() + " = nil\n"
				preamble += "var _out_len C.size_t = 0\n"
				afterword += "ret := make([]" + t.RenderTypeGo() + ", _out_Strlen)\n"
				afterword += "for i := 0; i < _out_len; i++ {\n"
				afterword += "ret[i] = new" + t.ParameterType + "(_out[i])\n"
				afterword += "}\n"
				afterword += "C.free(_out)\n"

			} else if m.ReturnType.QtClassType() {
				// Construct our Go type based on this inner CABI type
				shouldReturn = "ret := "

				if m.ReturnType.Pointer {
					afterword = "return " + m.ReturnType.ParameterType + "{h: ret}"
				} else {
					// This is return by value, but CABI has new'd it into a
					// heap type for us
					// To preserve Qt's approximate semantics, add a runtime
					// finalizer to automatically Delete once the type goes out
					// of Go scope
					imports["runtime"] = struct{}{}
					afterword = "// Qt uses pass-by-value semantics for this type. Mimic with finalizer\n"
					afterword += "ret1 := &" + m.ReturnType.ParameterType + "{h: ret}\n"
					afterword += "runtime.SetFinalizer(ret1, func(ret2 *" + m.ReturnType.ParameterType + ") {\n"
					afterword += "ret2.Delete()\n"
					afterword += "runtime.KeepAlive(ret2.h)\n"
					afterword += "})\n"
					afterword += "return ret1"
				}

			}

			ret.WriteString(`
			func (this *` + c.ClassName + `) ` + m.SafeMethodName() + `(` + emitParametersGo(m.Parameters) + `) ` + returnTypeDecl + ` {
				` + preamble +
				shouldReturn + ` C.` + c.ClassName + `_` + m.SafeMethodName() + `(` + forwarding + `)
` + afterword + `}
			
			`)
		}

		ret.WriteString(`
		func (this *` + c.ClassName + `) Delete() {
			C.` + c.ClassName + `_Delete(this.h)
		}
		`)

	}

	goSrc := ret.String()

	// Fixup imports
	if len(imports) > 0 {
		allImports := make([]string, 0, len(imports))
		for k, _ := range imports {
			allImports = append(allImports, `"`+k+`"`)
		}
		sort.Strings(allImports)
		goSrc = strings.Replace(goSrc, `%%_IMPORTLIBS_%%`, "import (\n\t"+strings.Join(allImports, "\n\t")+"\n)", 1)

	} else {
		goSrc = strings.Replace(goSrc, `%%_IMPORTLIBS_%%`, "", 1)

	}

	// Run gofmt over the result
	formattedSrc, err := format.Source([]byte(goSrc))
	if err != nil {
		log.Printf("gofmt failure: %v", err)
		formattedSrc = []byte(goSrc)
	}

	return string(formattedSrc), nil
}
