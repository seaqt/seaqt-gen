package main

import (
	"C"
	"fmt"
	"regexp"
	"sort"
	"strings"
)
import (
	"os"
	"path/filepath"
)

type nimFileState struct {
	stdImports   map[string]struct{}
	preImports   map[string]struct{}
	imports      map[string]struct{}
	postImports  map[string]struct{}
	qtModuleName string
	unitName     string
	ind          string
}

func (gfs nimFileState) qualifiedPrefix(qtModuleName string) string {
	// TODO using relative imports works as long as everything is in one giant
	// nimble package, but what if it's not?
	return ifv(qtModuleName == gfs.qtModuleName, "./", "../"+qtModuleName+"/")
}

func (gfs nimFileState) qualifiedTypeImport(v lookupResultClass) string {
	return gfs.qualifiedPrefix(v.QtModuleName) + v.UnitName + "_types"
}

func (gfs nimFileState) qualifiedImplImport(v lookupResultClass) string {
	return gfs.qualifiedPrefix(v.QtModuleName) + v.UnitName
}

func (gfs nimFileState) qualifiedImport(v lookupResultEnum) string {
	return gfs.qualifiedPrefix(v.QtModuleName) + v.UnitName + "_types"
}

func nimReservedWord(s string) bool {
	switch s {
	case "addr", "and", "as", "asm", "bind", "block", "break", "case", "cast",
		"concept", "const", "continue", "converter", "defer", "discard", "distinct",
		"div", "do", "elif", "else", "end", "enum", "except", "export", "finally",
		"for", "from", "func", "if", "import", "in", "include", "interface", "is",
		"isnot", "iterator", "let", "macro", "method", "mixin", "mod", "nil", "not",
		"notin", "object", "of", "or", "out", "proc", "ptr", "raise", "ref", "return",
		"shl", "shr", "static", "template", "try", "tuple", "type", "using", "var",
		"when", "while", "xor", "yield",
		"Exception",
		"super", "ret", "result", "create", "string", "seq": // not language-reserved words, but a binding-reserved words
		return true
	default:
		return false
	}
}

func safeMethodName(tmp string) string {
	// Strip redundant Qt prefix, we know these are all Qt functions
	tmp = strings.TrimPrefix(tmp, "qt_")

	// Operator-overload methods have names not representable in binding
	// languages. Replace more specific cases first
	replacer := strings.NewReplacer(

		// `operator ` with a trailing space only occurs in conversion operators
		// Add a fake _ here, but it will be replaced with camelcase in the regex below
		`operator `, `To `,
		`::`, `__`, // e.g. `operator QCborError::Code`

		`==`, `Equal`,
		`!=`, `NotEqual`,
		`>=`, `GreaterOrEqual`,
		`<=`, `LesserOrEqual`,
		`=`, `Assign`,

		`<<`, `ShiftLeft`, // Qt classes use it more for stream functions e.g. in QDataStream
		`>>`, `ShiftRight`,
		`>`, `Greater`,
		`<`, `Lesser`,

		`+`, `Plus`,
		`-`, `Minus`,
		`*`, `Multiply`,
		`/`, `Divide`,
		`%`, `Modulo`,

		`&&`, `LogicalAnd`,
		`||`, `LogicalOr`,
		`!`, `Not`,
		`&`, `BitwiseAnd`,
		`|`, `BitwiseOr`,
		`~`, `BitwiseXor`,
		`^`, `BitwiseNot`,

		`->`, `PointerDereference`,
		`[]`, `Subscript`,
		`()`, `Call`,
	)
	tmp = replacer.Replace(tmp)

	// Replace spaces (e.g. `operator long long` with CamelCase
	tmp = regexp.MustCompile(` ([a-zA-Z])`).ReplaceAllStringFunc(tmp, func(match string) string { return strings.ToUpper(match[1:]) })

	// Also replace any underscore_case with CamelCase
	// Only catch lowercase letters in this one, not uppercase, as it causes a
	// lot of churn for Scintilla
	tmp = regexp.MustCompile(`_([a-z])`).ReplaceAllStringFunc(tmp, func(match string) string { return strings.ToUpper(match[1:]) })
	if nimReservedWord(tmp) {
		tmp += "X"
	}

	return tmp
}
func (nm CppMethod) rawMethodName() string {
	return safeMethodName(nm.MethodName)
}

func (nm CppMethod) nimMethodName() string {
	return safeMethodName(nm.CppCallTarget())
}

func uniqueName(gfs *nimFileState, sigs map[string]struct{}, m CppMethod) string {
	paramsX := m.nimMethodName()
	for _, p := range m.Parameters {
		paramsX = paramsX + "," + p.renderTypeNim(gfs, false, false)
	}
	orig := paramsX
	j := 0
	for {
		if _, ok := sigs[paramsX]; ok {
			j += 1
			paramsX = maybeSuffix(j) + orig
		} else {
			sigs[paramsX] = struct{}{}
			break
		}
	}
	return m.nimMethodName() + maybeSuffix(j)
}

func (p CppParameter) nimParameterName() string {
	// Also make the first letter uppercase so it becomes public in Go
	parmName := p.ParameterName
	if nimReservedWord(parmName) {
		parmName += "Val"
	}

	if strings.HasPrefix(parmName, "_") {
		parmName = "x" + parmName
	}
	if strings.HasSuffix(parmName, "_") {
		parmName += "x"
	}

	return parmName
}

func cabiClassNameNim(className string, cabi bool) string {
	className = cabiClassName(className)
	className = strings.Replace(className, `__`, ``, -1)
	className = strings.TrimSuffix(className, "_")

	if cabi {
		className = "c" + cabiClassName((className))
	} else {

		className = cabiClassName(className)
	}

	if nimReservedWord(className) {
		className += "X"
	}

	return className
}

func ncabiSlotCallbackName(c CppClass, m CppMethod) string {
	return "f" + cabiClassNameNim(c.ClassName, true) + "_slot_callback_" + m.rawMethodName()
}

func ncabiVtableCallbackName(c CppClass, m CppMethod) string {
	return "f" + cabiClassNameNim(c.ClassName, true) + "_vtable_callback_" + m.rawMethodName()
}

func ncabiMethodCallbackName(c CppClass, m CppMethod) string {
	return "f" + cabiClassNameNim(c.ClassName, true) + "_method_callback_" + m.rawMethodName()
}

func ncabiNewName(c CppClass, i int) string {
	return "f" + cabiClassNameNim(c.ClassName, true) + `_new` + maybeSuffix(i)
}

func ncabiDeleteName(c CppClass) string {
	return "f" + cabiClassNameNim(c.ClassName, true) + `_delete`
}

func ncabiVirtBaseName(c CppClass) string {
	return "f" + cabiClassNameNim(c.ClassName, true) + `_virtbase`
}

func ncabiMethodName(c CppClass, m CppMethod) string {
	return "f" + cabiClassNameNim(c.ClassName, true) + `_` + m.rawMethodName()
}

func ncabiConnectName(c CppClass, m CppMethod) string {
	return "f" + cabiClassNameNim(c.ClassName, true) + `_connect_` + m.rawMethodName()
}

func ncabiVirtualBaseName(c CppClass, m CppMethod) string {
	return "f" + cabiClassNameNim(c.ClassName, true) + `_virtualbase_` + m.rawMethodName()
}
func ncabiProtectedBaseName(c CppClass, m CppMethod) string {
	return "f" + cabiClassNameNim(c.ClassName, true) + `_protectedbase_` + m.SafeMethodName()
}

func ncabiToVdataName(c CppClass) string {
	return "f" + cabiClassNameNim(c.ClassName, true) + `_vdata`
}
func ncabiFromVdataName(c CppClass) string {
	return "fvdata_" + cabiClassNameNim(c.ClassName, true)
}

func nimModulePkgName(qtModuleName string) string {
	return strings.ToLower(qtModuleName) + "_pkg"
}

func isQObject(s string) bool {
	return s == "QObject"
}

func (e CppEnum) nimEnumName() string {
	enumName := cabiClassNameNim(ifv(strings.HasSuffix(e.EnumName, "::"), e.EnumName+"Enum", e.EnumName), false) // Fully qualified name of the enum itself
	// if _, ok := KnownClassnames[enumName]; ok {
	// 	enumName = enumName + "Enum"
	// }
	return enumName + "Enum"
}

func (gfs *nimFileState) indent() {
	gfs.ind += "  "
}
func (gfs *nimFileState) dedent() {
	gfs.ind = gfs.ind[:len(gfs.ind)-2]
}

func (p CppParameter) renderTypeNim(gfs *nimFileState, cabi, isparam bool) string {
	if p.Pointer && p.ParameterType == "char" {
		if cabi {
			return "cstring"
		} else {
			return "cstring"
		}
	}
	if p.ParameterType == "QString" || p.ParameterType == "QAnyStringView" {
		if cabi {
			return "struct_miqt_string"
		} else {
			if isparam {
				return "openArray[char]"
			} else {
				return "string"
			}
		}
	}
	if p.ParameterType == "QByteArray" || p.ParameterType == "QByteArrayView" {
		if cabi {
			return "struct_miqt_string"
		} else {
			if isparam {
				return "openArray[byte]"
			} else {
				return "seq[byte]"
			}
		}
	}

	if t, ok := p.QListOf(); ok {
		if cabi {
			return "struct_miqt_array"
		} else {
			if isparam {
				return "openArray[" + t.renderTypeNim(gfs, cabi, false) + "]"
			}
			return "seq[" + t.renderTypeNim(gfs, cabi, false) + "]"
		}
	}

	if t, ok := p.QSetOf(); ok {
		gfs.stdImports["std/sets"] = struct{}{}
		return "HashSet[" + t.renderTypeNim(gfs, cabi, false) + "]"
	}

	if t1, t2, ok := p.QMapOf(); ok {
		if cabi {
			return "struct_miqt_map"
		} else {
			gfs.stdImports["std/tables"] = struct{}{}
			return "Table[" + t1.renderTypeNim(gfs, cabi, false) + "," + t2.renderTypeNim(gfs, cabi, false) + "]"
		}
	}

	if t1, t2, ok := p.QPairOf(); ok {
		if cabi {
			return "struct_miqt_map"
		} else {

			// Design QPair using capital-named members, in case it gets passed
			// across packages
			return "tuple[first: " + t1.renderTypeNim(gfs, cabi, false) + ", second: " + t2.renderTypeNim(gfs, cabi, false) + "]"
		}
	}

	if p.Pointer {
		if p.ParameterType == "void" || !AllowClass(p.ParameterType) {
			return "pointer"
		}
	}

	if p.ParameterType == "void" {
		return "void"
	}

	ret := ""
	isclass := false
	switch p.ParameterType {
	case "bool", "volatile bool":
		ret += "bool"
	case "char":
		ret += "cchar"
	case "unsigned char", "uchar":
		ret += "uint8" // TODO investigate better options
	case "quint8", "uint8_t":
		ret += "uint8"
	case "signed char":
		ret += "cschar"
	case "qint8", "int8_t":
		ret += "int8"
	case "short":
		ret += "cshort"
	case "qint16", "int16_t":
		ret += "int16"
	case "ushort", "unsigned short":
		ret += "cushort"
	case "quint16", "uint16_t":
		ret += "uint16"
	case "long":
		ret += "clong"
	case "ulong", "unsigned long":
		ret += "culong"
	case "int":
		ret += "cint"
	case "unsigned int":
		ret += "cuint"
	case "qint32":
		ret += "int32"
	case "quint32":
		ret += "uint32"
	case "qlonglong", "long long":
		ret += "clonglong"
	case "qint64":
		ret += "int64"
	case "qulonglong", "unsigned long long":
		ret += "culonglong"
	case "quint64":
		ret += "uint64"
	case "float":
		ret += "float32"
	case "double", "qreal":
		ret += "float64"
	case "size_t", "qsizetype":
		ret += "csize_t"
	case "QIntegerForSizeof<std::size_t>::Signed", "qptrdiff", "ptrdiff_t": // all signed
		if C.sizeof_size_t == 4 {
			ret += "int32"
		} else {
			ret += "int64"
		}
	case "qintptr", "intptr_t":
		ret += "miqt_intptr_t"
	case "quintptr", "uintptr_t":
		ret += "miqt_uintptr_t"
	case "QIntegerForSizeof<void *>::Unsigned", "QIntegerForSizeof<void *>::Signed":
		ret += "uint"
	default:

		if ft, ok := p.QFlagsOf(); ok {
			if cabi || true {
				ret += "cint"
			} else {
				if enumInfo, ok := KnownEnums[ft.ParameterType]; ok && enumInfo.UnitName != gfs.unitName {
					// Cross-package
					ret += enumInfo.UnitName + "." + enumInfo.Enum.nimEnumName()
					gfs.imports[gfs.qualifiedImport(enumInfo)] = struct{}{}
				} else {
					// Same package
					ret += enumInfo.Enum.nimEnumName()
				}
			}
		} else if enumInfo, ok := KnownEnums[p.ParameterType]; ok {
			if cabi || true {
				ret += "cint"
			} else {
				if enumInfo.UnitName != gfs.unitName {
					// Cross-package
					ret += enumInfo.UnitName + "." + enumInfo.Enum.nimEnumName()
					gfs.imports[gfs.qualifiedImport(enumInfo)] = struct{}{}
				} else {
					// Same package
					ret += enumInfo.Enum.nimEnumName()
				}
			}
		} else if strings.Contains(p.ParameterType, `::`) {
			// Inner class
			ret += cabiClassNameNim(p.ParameterType, cabi)

		} else {
			// Do not transform this type
			ret += cabiClassNameNim(p.ParameterType, cabi)
		}

	}

	if pkg, ok := KnownClassnames[p.ParameterType]; ok {
		isclass = true
		ret = pkg.UnitName + "_types." + ret
		if pkg.UnitName != gfs.unitName {
			gfs.imports[gfs.qualifiedTypeImport(pkg)] = struct{}{}
		}
	}

	if cabi && p.QtClassType() {
		ret = "pointer"
	} else if (cabi || !isclass) && (p.ByRef || p.Pointer) {
		ret = "ptr " + ret
	}

	return ret // ignore const
}

func (p CppParameter) parameterTypeNim(gfs *nimFileState) string {
	if p.ParameterType == "QString" || p.ParameterType == "QAnyStringView" {
		return "struct_miqt_string"
	}

	if p.ParameterType == "QByteArray" || p.ParameterType == "QByteArrayView" {
		return "struct_miqt_string"
	}

	if _, ok := p.QListOf(); ok {
		return "struct_miqt_array"
	}

	if _, ok := p.QSetOf(); ok {
		return "struct_miqt_array"
	}

	if _, _, ok := p.QMapOf(); ok {
		return "struct_miqt_map"
	}

	if _, _, ok := p.QPairOf(); ok {
		return "struct_miqt_map"
	}

	return p.renderTypeNim(gfs, true, false)
}

func (gfs *nimFileState) emitParametersNim(params []CppParameter, cabi bool, self string) string {
	tmp := make([]string, 0, len(params))

	skipNext := false

	if len(self) > 0 {
		tmp = append(tmp, self)
	}

	for i, p := range params {
		if IsArgcArgv(params, i) {
			skipNext = true
			if cabi {
				tmp = append(tmp, "argc: ptr cint, argv: cstringArray")
			}

		} else if skipNext {
			// Skip this parameter, already handled
			skipNext = false

		} else {
			// Ordinary parameter
			tmp = append(tmp, p.nimParameterName()+": "+p.renderTypeNim(gfs, cabi, true))

		}
	}
	return strings.Join(tmp, ", ")
}

func (gfs *nimFileState) emitParametersNim2CABIForwarding(m CppMethod, extra string) (preamble string, forwarding string) {
	tmp := make([]string, 0, len(m.Parameters)+2)

	if !m.IsStatic {
		tmp = append(tmp, "self.h")
	}

	if len(extra) > 0 {
		tmp = append(tmp, extra)
	}

	skipNext := false

	for i, p := range m.Parameters {

		if IsArgcArgv(m.Parameters, i) {
			skipNext = true
			// QApplication constructor. Convert 'args' into Qt's wanted types
			// Qt has a warning in the docs saying these pointers must be valid
			// for the entire lifetype of QApplication, so, malloc + never free
			// This transformation only affects the Go side. The CABI side is
			// projected naturally

			preamble += gfs.ind + "# Convert []string to long-lived int& argc, char** argv, never call free()\n"
			preamble += gfs.ind + "var args2 = @[getAppFilename()]\n"
			preamble += gfs.ind + "try:\n"

			gfs.indent()
			preamble += gfs.ind + "args2.add commandLineParams()\n"

			gfs.dedent()
			preamble += gfs.ind + "except OSError:\n"

			gfs.indent()
			preamble += gfs.ind + "echo getCurrentExceptionMsg()\n"

			gfs.dedent()
			preamble += gfs.ind + "var argv: cStringArray = allocCstringArray(args2)\n"
			preamble += gfs.ind + "var argc {.threadvar.}: cint\n"
			preamble += gfs.ind + "argc = args2.len.cint\n"

			tmp = append(tmp, "addr argc, argv")

			gfs.stdImports["std/cmdline"] = struct{}{}
			gfs.stdImports["std/os"] = struct{}{}

		} else if skipNext {
			// Skip this parameter, already handled
			skipNext = false

		} else {
			addPreamble, rvalue := gfs.emitParameterNim2CABIForwarding(p, false)
			preamble += addPreamble
			tmp = append(tmp, rvalue)
		}
	}

	return preamble, strings.Join(tmp, ", ")
}

func (gfs *nimFileState) emitParameterNim2CABIForwarding(p CppParameter, copy bool) (preamble, rvalue string) {
	// If copy is true, memory ownership is given to CABI (this happens when p is a return value)

	nameprefix := makeNamePrefix(p.nimParameterName())

	if p.ParameterType == "QString" || p.ParameterType == "QAnyStringView" {
		if copy {
			preamble += gfs.ind + "var " + nameprefix + "_copy = if len(" + p.nimParameterName() + ") > 0: c_malloc(csize_t(len(" + p.nimParameterName() + "))) else: nil\n"
			preamble += gfs.ind + "if len(" + p.nimParameterName() + ") > 0: copyMem(" + nameprefix + "_copy, addr " + p.nimParameterName() + "[0], csize_t(len(" + p.nimParameterName() + ")))\n"
			rvalue = "struct_miqt_string(data: " + nameprefix + "_copy, len: csize_t(len(" + p.nimParameterName() + ")))"
		} else {
			rvalue = "struct_miqt_string(data: if len(" + p.nimParameterName() + ") > 0: addr " + p.nimParameterName() + "[0] else: nil, len: csize_t(len(" + p.nimParameterName() + ")))"
		}
	} else if p.ParameterType == "QByteArray" || p.ParameterType == "QByteArrayView" {
		if copy {
			preamble += gfs.ind + "var " + nameprefix + "_copy = if len(" + p.nimParameterName() + ") > 0: c_malloc(csize_t(len(" + p.nimParameterName() + "))) else: nil\n"
			preamble += gfs.ind + "if len(" + p.nimParameterName() + ") > 0: copyMem(" + nameprefix + "_copy, addr " + p.nimParameterName() + "[0], csize_t(len(" + p.nimParameterName() + ")))\n"
			rvalue = "struct_miqt_string(data: " + nameprefix + "_copy, len: csize_t(len(" + p.nimParameterName() + ")))"
		} else {
			rvalue = "struct_miqt_string(data: if len(" + p.nimParameterName() + ") > 0: addr " + p.nimParameterName() + "[0] else: nil, len: csize_t(len(" + p.nimParameterName() + ")))"
		}
	} else if listType, ok := p.QListOf(); ok {
		// QList<T>
		// Go: convert T[] -> t* and len
		// CABI: create a real QList<>
		if copy {
			preamble += gfs.ind + "var " + nameprefix + "_CArray = cast[ptr UncheckedArray[" + listType.parameterTypeNim(gfs) + "]](if len(" + p.nimParameterName() + ") > 0: c_malloc(c_sizet(sizeof(" + listType.parameterTypeNim(gfs) + ") * len(" + p.nimParameterName() + "))) else: nil)\n"
		} else {
			preamble += gfs.ind + "var " + nameprefix + "_CArray = newSeq[" + listType.parameterTypeNim(gfs) + "](len(" + p.nimParameterName() + "))\n"
		}

		preamble += gfs.ind + "for i in 0..<len(" + p.nimParameterName() + "):\n"
		gfs.indent()

		listType.ParameterName = p.nimParameterName() + "[i]"
		addPreamble, innerRvalue := gfs.emitParameterNim2CABIForwarding(listType, copy)
		preamble += addPreamble
		preamble += gfs.ind + nameprefix + "_CArray[i] = " + innerRvalue + "\n"
		preamble += "\n"
		gfs.dedent()

		rvalue = "struct_miqt_array(len: csize_t(len(" + p.nimParameterName() + ")), data: if len(" + p.nimParameterName() + ") == 0: nil else: addr(" + nameprefix + "_CArray[0]))"

	} else if _, ok := p.QSetOf(); ok {
		panic("QSet<> arguments are not yet implemented") // n.b. doesn't seem to exist in QtCore/QtGui/QtWidgets at all

	} else if kType, vType, ok := p.QMapOf(); ok {
		// QMap<T>

		if copy {
			preamble += gfs.ind + "var " + nameprefix + "_Keys_CArray = cast[ptr UncheckedArray[" + kType.parameterTypeNim(gfs) + "]](if len(" + p.nimParameterName() + ") > 0: c_malloc(csize_t(sizeof(" + kType.parameterTypeNim(gfs) + ") * len(" + p.nimParameterName() + "))) else: nil)\n"
			preamble += gfs.ind + "var " + nameprefix + "_Values_CArray = cast[ptr UncheckedArray[" + vType.parameterTypeNim(gfs) + "]](if len(" + p.nimParameterName() + ") > 0: c_malloc(csize_t(sizeof(" + vType.parameterTypeNim(gfs) + ") * len(" + p.nimParameterName() + "))) else: nil)\n"
		} else {
			preamble += gfs.ind + "var " + nameprefix + "_Keys_CArray = newSeq[" + kType.parameterTypeNim(gfs) + "](len(" + p.nimParameterName() + "))\n"
			preamble += gfs.ind + "var " + nameprefix + "_Values_CArray = newSeq[" + vType.parameterTypeNim(gfs) + "](len(" + p.nimParameterName() + "))\n"
		}
		preamble += gfs.ind + "var " + nameprefix + "_ctr = 0\n"
		// TODO https://github.com/nim-lang/Nim/issues/24720
		// let's hope iteration order is stable :facepalm:
		preamble += gfs.ind + "for " + nameprefix + "_k in " + p.nimParameterName() + ".keys():\n"
		gfs.indent()
		kType.ParameterName = nameprefix + "_k"
		addPreamble, innerRvalue := gfs.emitParameterNim2CABIForwarding(kType, copy)
		preamble += addPreamble
		preamble += gfs.ind + nameprefix + "_Keys_CArray[" + nameprefix + "_ctr] = " + innerRvalue + "\n"
		preamble += gfs.ind + nameprefix + "_ctr += 1\n"
		gfs.dedent()
		preamble += gfs.ind + nameprefix + "_ctr = 0\n"
		if copy {
			preamble += gfs.ind + "for " + nameprefix + "_v in " + p.nimParameterName() + ".mvalues():\n"
		} else {
			preamble += gfs.ind + "for " + nameprefix + "_v in " + p.nimParameterName() + ".values():\n"
		}
		gfs.indent()
		vType.ParameterName = nameprefix + "_v"
		addPreamble, innerRvalue = gfs.emitParameterNim2CABIForwarding(vType, copy)
		preamble += addPreamble
		preamble += gfs.ind + nameprefix + "_Values_CArray[" + nameprefix + "_ctr] = " + innerRvalue + "\n"

		preamble += gfs.ind + nameprefix + "_ctr += 1\n"

		preamble += "\n"
		gfs.dedent()

		rvalue = "struct_miqt_map(len: csize_t(len(" + p.nimParameterName() + ")),keys: if len(" + p.nimParameterName() + ") == 0: nil else: addr(" + nameprefix + "_Keys_CArray[0]), values: if len(" + p.nimParameterName() + ") == 0: nil else: addr(" + nameprefix + "_Values_CArray[0]),)"

	} else if kType, vType, ok := p.QPairOf(); ok {
		// QPair<T>

		preamble += gfs.ind + "var " + nameprefix + "_CArray_First: " + kType.parameterTypeNim(gfs) + "\n"
		preamble += gfs.ind + "var " + nameprefix + "_CArray_Second: " + vType.parameterTypeNim(gfs) + "\n"

		kType.ParameterName = p.nimParameterName() + ".first"
		addPreamble, innerRvalue := gfs.emitParameterNim2CABIForwarding(kType, copy)
		preamble += addPreamble
		preamble += gfs.ind + nameprefix + "_CArray_First = " + innerRvalue + "\n"

		vType.ParameterName = p.nimParameterName() + ".second"
		addPreamble, innerRvalue = gfs.emitParameterNim2CABIForwarding(vType, copy)
		preamble += addPreamble
		preamble += gfs.ind + nameprefix + "_CArray_Second = " + innerRvalue + "\n"

		rvalue = "struct_miqt_map(len: 1,keys: addr(" + nameprefix + "_CArray_First),values: addr(" + nameprefix + "_CArray_Second),)"

	} else if p.Pointer && p.ParameterType == "char" {
		// Single char* argument
		rvalue = p.nimParameterName()

	} else if p.Pointer && !AllowClass(p.ParameterType) {
		// Single char* argument
		rvalue = p.nimParameterName()
	} else if /*(p.Pointer || p.ByRef) &&*/ p.QtClassType() {
		// The C++ type is a pointer to Qt class
		// We want our functions to accept the Go wrapper type, and forward as cPointer()
		// cPointer() returns the cgo pointer which only works in the same package -
		// anything cross-package needs to go via unsafe.Pointer

		if copy {
			// hack: this is a move, not a copy!
			preamble += gfs.ind + p.nimParameterName() + ".owned = false # TODO move?\n"
			preamble += gfs.ind + "let " + nameprefix + "_h = " + p.nimParameterName() + ".h\n"
			preamble += gfs.ind + p.nimParameterName() + ".h = nil\n"
			rvalue = nameprefix + "_h"
		} else {
			rvalue = p.nimParameterName() + ".h"
		}
	} else if p.IntType() || p.IsFlagType() || p.IsKnownEnum() || p.ParameterType == "bool" {
		if p.Pointer || p.ByRef {
			rvalue = p.nimParameterName() // n.b. This may not work if the integer type conversion was wrong
		} else if p.IsFlagType() || p.IsKnownEnum() {
			rvalue = "cint(" + p.nimParameterName() + ")"

		} else {
			rvalue = p.nimParameterName()
		}

	} else {
		// Default
		rvalue = p.nimParameterName()
	}

	return preamble, rvalue
}

func (gfs *nimFileState) emitCabiToNim(assignExpr string, rt CppParameter, rvalue string) string {

	lines := ""
	namePrefix := "v" + makeNamePrefix(rt.nimParameterName())

	if rt.Void() {
		return gfs.ind + rvalue + "\n"

	} else if rt.Pointer && !AllowClass(rt.ParameterType) {
		return gfs.ind + assignExpr + rvalue + "\n"
	} else if rt.ParameterType == "void" && rt.Pointer {
		return gfs.ind + assignExpr + rvalue + "\n"

	} else if rt.ParameterType == "char" && rt.Pointer {
		// Qt functions normally return QString - anything returning char*
		// is something like QByteArray.Data() where it returns an unsafe
		// internal pointer
		// However in case this is a signal, we need to be able to marshal both
		// forwards and backwards with the same types, this has to be a string
		// in both cases
		// This is not a miqt_string and therefore MIQT did not allocate it,
		// and therefore we don't have to free it either

		lines += gfs.ind + assignExpr + "(" + rvalue + ")\n"

	} else if rt.ParameterType == "QString" || rt.ParameterType == "QAnyStringView" {

		lines += gfs.ind + "let " + namePrefix + "_ms = " + rvalue + "\n"
		lines += gfs.ind + "let " + namePrefix + "x_ret = string.fromBytes(" + namePrefix + "_ms)\n"
		lines += gfs.ind + "c_free(" + namePrefix + "_ms.data)\n"
		lines += gfs.ind + assignExpr + namePrefix + "x_ret\n"

	} else if rt.ParameterType == "QByteArray" || rt.ParameterType == "QByteArrayView" {
		// We receive the CABI type of a miqt_string. Convert it into []byte
		// We must free the miqt_string data pointer - this is a data copy,
		// not an alias

		lines += gfs.ind + "var " + namePrefix + "_bytearray = " + rvalue + "\n"
		lines += gfs.ind + "var " + namePrefix + "x_ret = @(toOpenArray(cast[ptr UncheckedArray[byte]](" + namePrefix + "_bytearray.data), 0, int(" + namePrefix + "_bytearray.len)-1))\n"
		lines += gfs.ind + "c_free(" + namePrefix + "_bytearray.data)\n"
		lines += gfs.ind + assignExpr + namePrefix + "x_ret\n"

	} else if t, ok := rt.QListOf(); ok {
		lines += gfs.ind + "var " + namePrefix + "_ma = " + rvalue + "\n"

		lines += gfs.ind + "var " + namePrefix + "x_ret = newSeq[" + t.renderTypeNim(gfs, false, false) + "](int(" + namePrefix + "_ma.len))\n"
		lines += gfs.ind + "let " + namePrefix + "_outCast = cast[ptr UncheckedArray[" + t.parameterTypeNim(gfs) + "]](" + namePrefix + "_ma.data)\n"
		lines += gfs.ind + "for i in 0 ..< " + namePrefix + "_ma.len:\n"
		gfs.indent()
		lines += gfs.emitCabiToNim(namePrefix+"x_ret[i] = ", t, namePrefix+"_outCast[i]")
		gfs.dedent()
		lines += gfs.ind + "c_free(" + namePrefix + "_ma.data)\n"
		lines += gfs.ind + assignExpr + namePrefix + "x_ret\n"

	} else if t, ok := rt.QSetOf(); ok {

		lines += gfs.ind + "var " + namePrefix + "_ma = " + rvalue + "\n"

		lines += gfs.ind + namePrefix + "x_ret: HashSet[" + t.renderTypeNim(gfs, false, false) + "])\n"
		lines += gfs.ind + namePrefix + "_outCast = cast[ptr UncheckedArray[" + t.parameterTypeNim(gfs) + "](" + namePrefix + "_ma.data))\n"
		lines += gfs.ind + "for i in 0..<" + namePrefix + "_ma.len:\n"
		gfs.indent()
		lines += gfs.emitCabiToNim(namePrefix+"x_ret.incl ", t, namePrefix+"_outCast[i]")
		gfs.dedent()

		lines += gfs.ind + assignExpr + namePrefix + "x_ret\n"

	} else if kType, vType, ok := rt.QMapOf(); ok {
		lines += gfs.ind + "var " + namePrefix + "_mm = " + rvalue + "\n"

		lines += gfs.ind + "var " + namePrefix + "x_ret: Table[" + kType.renderTypeNim(gfs, false, false) + ", " + vType.renderTypeNim(gfs, false, false) + "]\n"
		lines += gfs.ind + "var " + namePrefix + "_Keys = cast[ptr UncheckedArray[" + kType.parameterTypeNim(gfs) + "]](" + namePrefix + "_mm.keys)\n"
		lines += gfs.ind + "var " + namePrefix + "_Values = cast[ptr UncheckedArray[" + vType.parameterTypeNim(gfs) + "]](" + namePrefix + "_mm.values)\n"
		lines += gfs.ind + "for i in 0..<" + namePrefix + "_mm.len:\n"
		gfs.indent()
		lines += gfs.emitCabiToNim("var "+namePrefix+"_entry_Key = ", kType, namePrefix+"_Keys[i]") + "\n"
		lines += gfs.emitCabiToNim("var "+namePrefix+"_entry_Value = ", vType, namePrefix+"_Values[i]") + "\n"
		lines += gfs.ind + namePrefix + "x_ret[" + namePrefix + "_entry_Key] = " + namePrefix + "_entry_Value\n"
		gfs.dedent()
		lines += gfs.ind + "c_free(" + namePrefix + "_mm.keys)\n"
		lines += gfs.ind + "c_free(" + namePrefix + "_mm.values)\n"

		lines += gfs.ind + assignExpr + namePrefix + "x_ret\n"

	} else if kType, vType, ok := rt.QPairOf(); ok {
		lines += gfs.ind + "var " + namePrefix + "_mm = " + rvalue + "\n"

		lines += gfs.ind + "var " + namePrefix + "_First_CArray = cast[ptr UncheckedArray[" + kType.parameterTypeNim(gfs) + "]](" + namePrefix + "_mm.keys)\n"
		lines += gfs.ind + "var " + namePrefix + "_Second_CArray = cast[ptr UncheckedArray[" + vType.parameterTypeNim(gfs) + "]](" + namePrefix + "_mm.values)\n"

		lines += gfs.emitCabiToNim("var "+namePrefix+"_entry_First = ", kType, namePrefix+"_First_CArray[0]") + "\n"
		lines += gfs.emitCabiToNim("var "+namePrefix+"_entry_Second = ", vType, namePrefix+"_Second_CArray[0]") + "\n"

		lines += gfs.ind + "c_free(" + namePrefix + "_mm.keys)\n"
		lines += gfs.ind + "c_free(" + namePrefix + "_mm.values)\n"

		lines += gfs.ind + assignExpr + "(first: " + namePrefix + "_entry_First , second: " + namePrefix + "_entry_Second )\n"

	} else if rt.QtClassType() {
		// Construct our Go type based on this inner CABI type

		pkg, ok := KnownClassnames[rt.ParameterType]
		if !ok {
			panic("emitCabiToNim: Encountered an unknown Qt class")
		}

		crossPackage := pkg.UnitName + "_types."

		if pkg.UnitName != gfs.unitName {
			gfs.imports[gfs.qualifiedTypeImport(pkg)] = struct{}{}
		}

		// We can only reference the rvalue once, in case it is a complex
		// expression

		if !(rt.Pointer || rt.ByRef) {
			// This is return by value, but CABI has new'd it into a
			// heap type for us - assume ownership of the instace
			rvalue = crossPackage + cabiClassNameNim(rt.ParameterType, false) + "(h: " + rvalue + ", owned: true)"
			lines += gfs.ind + assignExpr + rvalue + "\n"
		} else {
			rvalue = crossPackage + cabiClassNameNim(rt.ParameterType, false) + "(h: " + rvalue + ", owned: false)"
			lines += gfs.ind + assignExpr + rvalue + "\n"
		}

		return lines

	} else if rt.IntType() || rt.IsKnownEnum() || rt.IsFlagType() || rt.ParameterType == "bool" || rt.QtCppOriginalType != nil {

		if rt.Pointer || rt.ByRef {
			//return assignExpr + "(" + rt.renderTypeNim(gfs, false) + "(" + rvalue + "))\n"
			lines += gfs.ind + assignExpr + rvalue + "\n"
		} else if rt.IsKnownEnum() || rt.IsFlagType() {
			lines += gfs.ind + assignExpr + rt.renderTypeNim(gfs, false, false) + "(" + rvalue + ")\n"
		} else {

			lines += gfs.ind + assignExpr + rvalue + "\n"
		}
	} else {
		panic(fmt.Sprintf("emitgo::emitCabiToNim missing type handler for parameter %+v", rt))
	}

	return lines
}

func emitNimPkg(qtModuleName, pcName string) string {
	return fmt.Sprintf(`const
  %[1]sCFlags* =
    gorge("pkg-config --cflags %[2]s") &
    (when defined(gcc) or defined(llvm): " -fPIC" else: "")

  %[1]sLibs* = gorge("pkg-config --libs %[2]s")

  %[1]sGenVersion* = "%[3]s"
    ## The version used for generating the bindings

  %[1]sBuildVersion* = gorge("pkg-config --modversion %[2]s")
    ## The version used when compiling the application

{.passl: %[1]sLibs}
`, qtModuleName, pcName, pkgConfig(pcName, "--modversion"))
}

func emitNim(src *CppParsedHeader, headerName string, qtModuleName string, pcName string) (string, string, error) {

	ret := strings.Builder{}
	cabi := strings.Builder{}
	types := strings.Builder{}

	ret.WriteString(`import ./` + nimModulePkgName(qtModuleName) + `

{.push raises: [].}

from system/ansi_c import c_free, c_malloc

type
  struct_miqt_string {.used.} = object
    len: csize_t
    data: pointer

  struct_miqt_array {.used.} = object
    len: csize_t
    data: pointer

  struct_miqt_map {.used.} = object
    len: csize_t
    keys: pointer
    values: pointer

  miqt_uintptr_t {.importc: "uintptr_t", header: "stdint.h", used.} = uint
  miqt_intptr_t {.importc: "intptr_t", header: "stdint.h", used.} = int

func fromBytes(T: type string, v: struct_miqt_string): string {.used.} =
  if v.len > 0:
    let len = cast[int](v.len)
    result = newStringUninit(len)
    when nimvm:
      let d = cast[ptr UncheckedArray[char]](v.data)
      for i in 0..<len:
        result[i] = d[i]
    else:
      copyMem(addr result[0], v.data, len)

`)

	// Type definition
	gfs := nimFileState{
		imports:      map[string]struct{}{},
		preImports:   map[string]struct{}{},
		stdImports:   map[string]struct{}{},
		qtModuleName: qtModuleName,
		unitName:     genUnitName(headerName),
		ind:          "  ",
	}

	// messy: pkg-config flags don't include private headers
	if headerName == "qobject.h" {
		cabi.WriteString(`import std/strutils
const privateDir = block:
  var flag = ""
  for path in QtCoreCFlags.split(" "):
    if "QtCore" in path:
      flag = " " & path & "/" & QtCoreBuildVersion & " " & path & "/" & QtCoreBuildVersion & "/QtCore"
      break
  flag

{.compile("../libseaqt-runtime.cpp", QtCoreCFlags & privateDir).}

type QObjectconnectRawSlot* = proc(args: pointer)

proc QObject_slot_callback_connectRaw(slot: int, args: pointer) {.cdecl.} =
  let slot = cast[ptr QObjectconnectRawSlot](slot)
  slot[](args)

proc QObject_slot_callback_connectRaw_release(slot: int) {.cdecl.} =
  let slot = cast[ref QObjectconnectRawSlot](slot)
  GC_unref(slot)

proc fcQObject_connectRawSlot(
  sender: pointer,
  signal: cstring,
  receiver: pointer,
  slot: int,
  callback: pointer,
  release: pointer,
  typeVal: cint,
  senderMetaObject: pointer,
): pointer {.importc: "QObject_connectRawSlot".}

proc connectRaw*(
    _: type gen_qobject_types.QObject,
    sender: gen_qobject_types.QObject,
    signal: cstring,
    receiver: gen_qobject_types.QObject,
    slot: QObjectconnectRawSlot,
    typeVal: cint,
    senderMetaObject: gen_qobjectdefs_types.QMetaObject,
): gen_qobjectdefs_types.QMetaObjectConnection =
  var tmp = new QObjectconnectRawSLot
  tmp[] = slot
  GC_ref(tmp)
  gen_qobjectdefs_types.QMetaObjectConnection(
    h: fcQObject_connectRawSlot(
      sender.h,
      signal,
      receiver.h,
      cast[int](addr tmp[]),
      QObject_slot_callback_connectRaw,
      QObject_slot_callback_connectRaw_release,
      typeVal,
      senderMetaObject.h,
    ),
    owned: true,
  )

`)
	}

	hasCompile := false
	compileCpp := `
{.compile("gen_` + strings.Replace(headerName, ".h", ".cpp", -1) + `", ` + qtModuleName + `CFlags).}

`

	for _, c := range src.Classes {
		rawClassName := cabiClassNameNim(c.ClassName, true)
		nimClassName := cabiClassNameNim(c.ClassName, false)
		importClassName := cabiClassName(c.ClassName)

		pragmas := " {.inheritable.}"

		inherit := ""
		mi := false
		for _, base := range c.DirectInherits {

			if mi {
				types.WriteString("# TODO Multiple inheritance from " + base + "\n")
			} else {
				if strings.HasPrefix(base, `QList<`) {
					// Can't inherit
					continue
				} else if pkg, ok := KnownClassnames[base]; ok && pkg.UnitName != gfs.unitName {
					// Cross-package parent class
					inherit = " of " + pkg.UnitName + "_types." + cabiClassNameNim(base, false)
					if _, ok := gfs.preImports[gfs.qualifiedTypeImport(pkg)]; !ok {
						types.WriteString(`import ` + gfs.qualifiedTypeImport(pkg) + `
export ` + pkg.UnitName + `_types

`)

						gfs.preImports[gfs.qualifiedTypeImport(pkg)] = struct{}{}
					}
					gfs.imports[gfs.qualifiedImplImport(pkg)] = struct{}{}
				} else {
					// Same-package parent class
					inherit = " of " + cabiClassNameNim(base, false)
				}
				pragmas = ""
				mi = true
			}
		}

		fmt.Fprintf(&cabi, `type %[1]s*{.exportc: "%[2]s", incompleteStruct.} = object
`, rawClassName, importClassName)

		fmt.Fprintf(&types, `type %[1]s*%[2]s = object%[3]s
`, nimClassName, pragmas, inherit)

		if inherit == "" {
			types.WriteString(`  h*: pointer
  owned*: bool

`)

			if c.CanDelete {
				if !hasCompile {
					// The destructor must live in the same module as the type declaration meaning that
					// we need access to the generated code even if we only shuffle pointers arouund -
					// this might have a better solution
					types.WriteString("import ./" + nimModulePkgName(qtModuleName) + "\n")
					types.WriteString(compileCpp)
					hasCompile = true
				}

				types.WriteString(`proc ` + ncabiDeleteName(c) + `(self: pointer) {.importc: "` + cabiDeleteName(c) + `".}
`)
				types.WriteString("proc `=destroy`(self: var " + nimClassName + `) =
  if self.owned: ` + ncabiDeleteName(c) + `(self.h)

`)
				types.WriteString(`proc ` + "`=sink`" + `(dest: var ` + nimClassName + `, source: ` + nimClassName + `) =
  ` + "`=destroy`" + `(dest)
  wasMoved(dest)
  dest.h = source.h
  dest.owned = source.owned

`)
				// TODO copy constructors
				// https://github.com/nim-lang/Nim/issues/24760
				types.WriteString("proc `=copy`(dest: var " + nimClassName + ", source: " + nimClassName + ") {.error.}\n")

				types.WriteString(`proc delete*(self: sink ` + nimClassName + `) =
  let h = self.h
  wasMoved(self)
  ` + ncabiDeleteName(c) + `(h)

`)
				if isQObject(c.ClassName) {
					types.WriteString(`proc fcQObject_deleteLater(self: pointer) {.importc: "QObject_deleteLater".}
proc deleteLater*(self: sink ` + nimClassName + `) =
  let h = self.h
  wasMoved(self)
  fcQObject_deleteLater(h)

`)
				}
			}

		} else {
			// https://github.com/nim-lang/Nim/issues/24760
			types.WriteString("proc `=copy`(dest: var " + nimClassName + ", source: " + nimClassName + ") {.error.}\n")

			// https://github.com/nim-lang/Nim/issues/24762
			// https://github.com/nim-lang/Nim/issues/24764
			types.WriteString(`proc ` + "`=sink`" + `(dest: var ` + nimClassName + `, source: ` + nimClassName + `) =
  ` + "`=destroy`" + `(dest)
  wasMoved(dest)
  dest.h = source.h
  dest.owned = source.owned

`)
		}
	}

	if !hasCompile {
		// ... but when we can, put the compile directive in the implementation
		// file
		ret.WriteString(compileCpp)
		hasCompile = true
	}

	cabi.WriteString("\n")

	// Check if short-named enums are allowed.
	// We only allow short names if there are no conflicts anywhere in the whole
	// file. This doesn't fully defend against cross-file conflicts but those
	// should hopefully be rare enough
	preventShortNames := map[string]struct{}{}
	{
		nameTest := map[string]string{}
	nextEnum:
		for _, e := range src.Enums {
			if e.EnumName == "" {
				continue // Removed by transformRedundant AST pass or empty
			}

			shortEnumName := e.ShortEnumName()

			// Disallow entry<-->entry collisions
			for _, ee := range e.Entries {
				if other, ok := nameTest[shortEnumName+"::"+ee.EntryName]; ok {
					preventShortNames[e.nimEnumName()] = struct{}{} // Our full enum name
					preventShortNames[other] = struct{}{}           // Their full enum name
					continue nextEnum
				}
				nameTest[shortEnumName+"::"+ee.EntryName] = e.nimEnumName()

				if _, ok := KnownClassnames[shortEnumName+"::"+ee.EntryName]; ok {
					preventShortNames[e.nimEnumName()] = struct{}{}
					continue nextEnum
				}
				if _, ok := KnownEnums[shortEnumName+"::"+ee.EntryName]; ok {
					preventShortNames[e.nimEnumName()] = struct{}{}
					continue nextEnum
				}

			}
		}
	}

	for _, e := range src.Enums {
		if e.EnumName == "" {
			continue // Removed by transformRedundant AST pass or empty
		}

		// "::" at the end means an anonymous nested enum
		enumName := e.nimEnumName()

		// enumShortName := enumName // Shorter name, so that enum elements are reachable from the surrounding namespace
		// if _, ok := preventShortNames[enumName]; !ok {
		// 	enumShortName = cabiClassNameNim(e.ShortEnumName(), false)
		// }

		underlyingType := e.UnderlyingType.renderTypeNim(&gfs, false, false)
		ret.WriteString(`
type ` + enumName + `* = distinct ` + underlyingType + `
`)

		if len(e.Entries) > 0 {
			zoo := map[string]struct{}{}
			for _, ee := range e.Entries {
				i := 0

				basename := ee.EntryName
				if nimReservedWord(basename) {
					basename += "Val"
				}
				nimbase := strings.ReplaceAll(strings.ToLower(basename), "_", "")
				nimname := nimbase + maybeSuffix(i)
				for {
					if _, ok := zoo[nimname]; ok {
						i += 1
						nimname = nimbase + maybeSuffix(i)
					} else if _, ok := KnownClassnames[nimname]; ok {
						i += 1
						nimname = nimbase + maybeSuffix(i)
					} else if _, ok := KnownClassnames[basename+maybeSuffix(i)]; ok {
						i += 1
						nimname = nimbase + maybeSuffix(i)
					} else {
						break
					}
				}
				zoo[nimname] = struct{}{}
				entryName := basename + maybeSuffix(i)

				ret.WriteString("template " + entryName + "*(_: type " + enumName + "): untyped = " + ee.EntryValue + "\n")
			}

			ret.WriteString("\n")

		}
	}
	gfs.preImports[gfs.qualifiedPrefix(qtModuleName)+gfs.unitName+`_types`] = struct{}{}
	ret.WriteString(`
import ` + gfs.qualifiedPrefix(qtModuleName) + gfs.unitName + `_types
export ` + gfs.unitName + `_types

%%_IMPORTLIBS_%%
%%_CABI_%%
`)
	for _, c := range src.Classes {
		nimClassName := cabiClassNameNim(c.ClassName, false)
		nimPkgClassName := gfs.unitName + `_types.` + nimClassName
		rawClassName := cabiClassNameNim(c.ClassName, true)
		virtualMethods := c.VirtualMethods()
		protectedMethods := c.ProtectedMethods()

		// Qt has some overloads (const vs non-const, & vs *) that don't result in
		// a distinct parameter set on the nim side
		sigs := map[string]struct{}{}

		for _, m := range c.Methods {
			if m.IsProtected {
				continue // Don't add a direct call for it
			}

			if m.MethodName == "deleteLater" && c.ClassName == "QObject" {
				continue // already in ..._types
			}

			preamble, forwarding := gfs.emitParametersNim2CABIForwarding(m, "")

			returnTypeDecl := m.ReturnType.renderTypeNim(&gfs, false, false)
			rawReturnTypeDecl := m.ReturnType.renderTypeNim(&gfs, true, false)
			rawMethodName := ncabiMethodName(c, m)
			nimMethodName := uniqueName(&gfs, sigs, m)
			rvalue := rawMethodName + `(` + forwarding + `)`

			params := gfs.emitParametersNim(m.Parameters, false, ifv(m.IsStatic, `_: type `, `self: `)+nimPkgClassName)

			// TOOD `this: ptr ` + rawClassName + `, `?
			rawParams := gfs.emitParametersNim(m.Parameters, true, ifv(m.IsStatic, "", "self: pointer"))

			fmt.Fprintf(&cabi, `proc %[1]s(%[2]s): %[3]s {.importc: "%[4]s".}
`, rawMethodName, rawParams, rawReturnTypeDecl, cabiMethodName(c, m))

			fmt.Fprintf(&ret, `proc %[1]s*(%[2]s): %[3]s =
%[4]s%[5]s
`, nimMethodName, params, returnTypeDecl, preamble, gfs.emitCabiToNim("", m.ReturnType, rvalue))

			// Add Connect() wrappers for signal functions
			if m.IsSignal {
				var namedParams []string
				var paramNames []string
				conversion := ""

				namedParams = append(namedParams, "slot: int")
				for i, pp := range m.Parameters {
					namedParams = append(namedParams, pp.nimParameterName()+": "+pp.parameterTypeNim(&gfs))

					paramNames = append(paramNames, fmt.Sprintf("slotval%d", i+1))
					conversion += gfs.emitCabiToNim(fmt.Sprintf("let slotval%d = ", i+1), pp, pp.nimParameterName()) + "\n"
				}

				cbTypeName := nimClassName + m.rawMethodName() + "Slot"
				cbType := `proc(` + gfs.emitParametersNim(m.Parameters, false, "") + `)`

				fmt.Fprintf(&cabi, `proc %[1]s(self: pointer, slot: int, callback: proc (%[3]s) {.cdecl.}, release: proc(slot: int) {.cdecl.}) {.importc: "%[2]s".}
`, ncabiConnectName(c, m), cabiConnectName(c, m), strings.Join(namedParams, ", "))

				fmt.Fprintf(&ret, `type %[1]s* = %[2]s
proc %[3]s(%[4]s) {.cdecl.} =
  let nimfunc = cast[ptr %[1]s](cast[pointer](slot))
%[5]s  nimfunc[](%[6]s)

proc %[3]s_release(slot: int) {.cdecl.} =
  let nimfunc = cast[ref %[1]s](cast[pointer](slot))
  GC_unref(nimfunc)

proc on%[8]s*(self: %[9]s, slot: %[1]s) =
  var tmp = new %[1]s
  tmp[] = slot
  GC_ref(tmp)
  %[7]s(self.h, cast[int](addr tmp[]), %[3]s, %[3]s_release)

`, cbTypeName, cbType, ncabiSlotCallbackName(c, m), strings.Join(namedParams, ", "),
					conversion, strings.Join(paramNames, `, `), ncabiConnectName(c, m),
					titleCase(m.nimMethodName()), nimPkgClassName)
			}
		}

		if len(virtualMethods) > 0 {
			for _, m := range virtualMethods {
				cbTypeName := nimClassName + m.rawMethodName() + "Proc"
				fmt.Fprintf(&ret, "type %s* = proc(%s): %s {.raises: [], gcsafe.}\n",
					cbTypeName, gfs.emitParametersNim(m.Parameters, false, "self: "+nimClassName), m.ReturnType.renderTypeNim(&gfs, false, false))
			}

			// `ptr pointer` here because we always allocate sizeof(pointer) extra mem
			fmt.Fprintf(&cabi, `proc %[2]s(self: pointer): ptr pointer {.importc: "%[3]s".}
proc %[4]s(self: pointer): pointer {.importc: "%[5]s".}

type %[1]sVTable {.pure.} = object
  destructor*: proc(self: pointer) {.cdecl, raises:[], gcsafe.}
`, rawClassName, ncabiToVdataName(c), cabiToVdataName(c),
				ncabiFromVdataName(c), cabiFromVdataName(c))

			fmt.Fprintf(&ret, `
type %[1]sVTable* {.inheritable, pure.} = object
  vtbl: %[2]sVTable
`, nimClassName, rawClassName)

			for _, m := range virtualMethods {
				fmt.Fprintf(&cabi, "  %s*: proc(%s): %s {.cdecl, raises: [], gcsafe.}\n",
					m.rawMethodName(), gfs.emitParametersNim(m.Parameters, true, "self: pointer"), m.ReturnType.renderTypeNim(&gfs, true, false))

				cbTypeName := nimClassName + m.rawMethodName() + "Proc"
				fmt.Fprintf(&ret, "  %s*: %s\n", m.rawMethodName(), cbTypeName)
			}

			ret.WriteString("\n")

			for _, m := range virtualMethods {
				// Add a package-private function to call the C++ base class method
				// QWidget_virtualbase_PaintEvent
				// This is only possible if the function is not pure-virtual

				if !m.IsPureVirtual {
					preamble, forwarding := gfs.emitParametersNim2CABIForwarding(m, "")

					forwarding = "self.h" + strings.TrimPrefix(forwarding, `self.h`) // TODO integrate properly

					returnTypeDecl := m.ReturnType.renderTypeNim(&gfs, false, false)
					rawReturnTypeDecl := m.ReturnType.renderTypeNim(&gfs, true, false)

					fmt.Fprintf(&cabi, `proc %[1]s(%[3]s): %[4]s {.importc: "%[2]s".}
`, ncabiVirtualBaseName(c, m), cabiVirtualBaseName(c, m), gfs.emitParametersNim(m.Parameters, true, "self: pointer"), rawReturnTypeDecl)

					fmt.Fprintf(&ret, `proc %[1]s%[2]s*(%[4]s): %[5]s =
%[6]s%[7]s
`,
						nimClassName, m.nimMethodName(), nimPkgClassName, gfs.emitParametersNim(m.Parameters, false, "self: "+nimPkgClassName), returnTypeDecl,
						preamble,
						gfs.emitCabiToNim("", m.ReturnType, ncabiVirtualBaseName(c, m)+`(`+forwarding+`)`),
					)
				}
			}

			ret.WriteString("\n")

			for _, m := range virtualMethods {

				conversion := ""

				{
					var namedParams []string
					var paramNames []string

					namedParams = append(namedParams, "self: pointer")

					paramNames = append(paramNames, "self")

					for i, pp := range m.Parameters {
						namedParams = append(namedParams, pp.nimParameterName()+": "+pp.parameterTypeNim(&gfs))

						paramNames = append(paramNames, fmt.Sprintf("slotval%d", i+1))
						conversion += gfs.emitCabiToNim(fmt.Sprintf("let slotval%d = ", i+1), pp, pp.nimParameterName())
					}

					cabiReturnType := m.ReturnType.parameterTypeNim(&gfs)

					ret.WriteString(`proc ` + ncabiVtableCallbackName(c, m) + `(` + strings.Join(namedParams, `, `) + `): ` + cabiReturnType + ` {.cdecl.} =
  let vtbl = cast[ptr ` + nimClassName + `VTable](` + ncabiToVdataName(c) + `(self)[])
  let self = ` + nimClassName + `(h: self)
`)
					ret.WriteString(conversion)
					if cabiReturnType == "void" {
						ret.WriteString(gfs.ind + `vtbl[].` + m.rawMethodName() + `(` + strings.Join(paramNames, `, `) + ")\n\n")
					} else {
						ret.WriteString(gfs.ind + `var virtualReturn = vtbl[].` + m.rawMethodName() + `(` + strings.Join(paramNames, `, `) + ")\n")
						virtualRetP := m.ReturnType // copy
						virtualRetP.ParameterName = "virtualReturn"
						binding, rvalue := gfs.emitParameterNim2CABIForwarding(virtualRetP, true)
						ret.WriteString(binding)
						ret.WriteString(gfs.ind + rvalue + "\n\n")
					}
				}
			}

			fmt.Fprintf(&ret, `type Virtual%[1]s* {.inheritable.} = ref object of %[1]s
  vtbl*: %[2]sVTable

`, nimClassName, rawClassName)

			for _, m := range virtualMethods {
				returnTypeDecl := m.ReturnType.renderTypeNim(&gfs, false, false)

				fmt.Fprintf(&ret, `method %[2]s*(%[4]s): %[5]s {.base.} =
`,
					nimClassName, m.nimMethodName(), nimPkgClassName, gfs.emitParametersNim(m.Parameters, false, "self: Virtual"+nimClassName), returnTypeDecl,
				)
				if !m.IsPureVirtual {
					var paramNames []string

					paramNames = append(paramNames, "self[]")

					for _, pp := range m.Parameters {
						paramNames = append(paramNames, pp.nimParameterName())
					}

					fmt.Fprintf(&ret, `  %[1]s%[2]s(%[3]s)
`,
						nimClassName, m.nimMethodName(), strings.Join(paramNames, ", "))

				} else {
					fmt.Fprintf(&ret, "  raiseAssert(\"missing implementation of %s.%s\")\n", nimClassName, m.nimMethodName())
				}
			}

			ret.WriteString("\n")

			for _, m := range virtualMethods {
				conversion := ""

				{
					var namedParams []string
					var paramNames []string

					namedParams = append(namedParams, "self: pointer")

					for i, pp := range m.Parameters {
						namedParams = append(namedParams, pp.nimParameterName()+": "+pp.parameterTypeNim(&gfs))

						paramNames = append(paramNames, fmt.Sprintf("slotval%d", i+1))
						conversion += gfs.emitCabiToNim(fmt.Sprintf("let slotval%d = ", i+1), pp, pp.nimParameterName())
					}

					cabiReturnType := m.ReturnType.parameterTypeNim(&gfs)

					ret.WriteString(`proc ` + ncabiMethodCallbackName(c, m) + `(` + strings.Join(namedParams, `, `) + `): ` + cabiReturnType + ` {.cdecl.} =
  let inst = cast[Virtual` + nimClassName + `](` + ncabiToVdataName(c) + `(self)[])
`)
					ret.WriteString(conversion)
					if cabiReturnType == "void" {
						ret.WriteString(gfs.ind + `inst.` + m.nimMethodName() + `(` + strings.Join(paramNames, `, `) + ")\n\n")
					} else {
						ret.WriteString(gfs.ind + `var virtualReturn = inst.` + m.nimMethodName() + `(` + strings.Join(paramNames, `, `) + ")\n")
						virtualRetP := m.ReturnType // copy
						virtualRetP.ParameterName = "virtualReturn"
						binding, rvalue := gfs.emitParameterNim2CABIForwarding(virtualRetP, true)
						ret.WriteString(binding)
						ret.WriteString(gfs.ind + rvalue + "\n\n")
					}
				}
			}
			if len(virtualMethods) > 0 {
				ret.WriteString("\n")
			}
		}

		for _, m := range protectedMethods {
			// Add a package-private function to call the C++ base class method
			// QWidget_virtualbase_PaintEvent
			// This is only possible if the function is not pure-virtual

			if !m.IsPureVirtual {
				preamble, forwarding := gfs.emitParametersNim2CABIForwarding(m, "")

				forwarding = "self.h" + strings.TrimPrefix(forwarding, `self.h`) // TODO integrate properly

				returnTypeDecl := m.ReturnType.renderTypeNim(&gfs, false, false)
				rawReturnTypeDecl := m.ReturnType.renderTypeNim(&gfs, true, false)

				fmt.Fprintf(&cabi, `proc %[1]s(%[3]s): %[4]s {.importc: "%[2]s".}
`, ncabiProtectedBaseName(c, m), cabiProtectedBaseName(c, m), gfs.emitParametersNim(m.Parameters, true, "self: pointer"), rawReturnTypeDecl)

				fmt.Fprintf(&ret, `proc %[2]s*(%[4]s): %[5]s =
%[6]s%[7]s
`,
					nimClassName, m.nimMethodName(), nimPkgClassName, gfs.emitParametersNim(m.Parameters, false, "self: "+nimPkgClassName), returnTypeDecl,
					preamble,
					gfs.emitCabiToNim("", m.ReturnType, ncabiProtectedBaseName(c, m)+`(`+forwarding+`)`),
				)
			}
		}

		if len(c.Ctors) > 0 {
			for i, ctor := range c.Ctors {
				preamble, forwarding := gfs.emitParametersNim2CABIForwarding(ctor, ifv(len(virtualMethods) > 0, "addr(vtbl[].vtbl), csize_t(sizeof(pointer))", ""))
				cabiParams := gfs.emitParametersNim(ctor.Parameters, true, ifv(len(virtualMethods) > 0, "vtbl: pointer, vdata: csize_t", ""))
				fmt.Fprintf(&cabi, `proc %[1]s(%[2]s): ptr %[3]s {.importc: "%[4]s".}
`, ncabiNewName(c, i), cabiParams, rawClassName, cabiNewName(c, i))

				nimParams := gfs.emitParametersNim(ctor.Parameters, false, "")
				paramsX := ""
				for _, p := range ctor.Parameters {
					paramsX = paramsX + "," + p.renderTypeNim(&gfs, false, true)
				}
				orig := paramsX
				j := 0
				for {
					if _, ok := sigs[paramsX]; ok {
						j += 1
						paramsX = maybeSuffix(j) + orig
					} else {
						sigs[paramsX] = struct{}{}
						break
					}
				}

				vparams := []string{}
				vparams = append(vparams, "T: type "+nimPkgClassName)
				if len(nimParams) > 0 {
					vparams = append(vparams, nimParams)
				}
				if len(virtualMethods) > 0 {
					preamble = preamble + `  let vtbl = if vtbl == nil: new ` + nimClassName + `VTable else: vtbl
  GC_ref(vtbl)
  vtbl[].vtbl.destructor = proc(self: pointer) {.cdecl.} =
    let vtbl = cast[ref ` + nimClassName + `VTable](` + ncabiToVdataName(c) + `(self)[])
    GC_unref(vtbl)
`
					vparams = append(vparams, "vtbl: ref "+nimClassName+"VTable = nil")
					for _, m := range virtualMethods {
						preamble = preamble + fmt.Sprintf(`  if not isNil(vtbl[].%[1]s):
    vtbl[].vtbl.%[1]s = %[2]s
`, m.rawMethodName(), ncabiVtableCallbackName(c, m))
					}
				}

				fmt.Fprintf(&ret, `proc create%[1]s*(%[3]s): %[2]s =
%[4]s  let tmp = %[2]s(h: %[5]s(%[6]s), owned: true)
`, maybeSuffix(j), nimPkgClassName, strings.Join(vparams, ",\n    "),
					preamble, ncabiNewName(c, i), forwarding)

				if len(virtualMethods) > 0 {
					fmt.Fprintf(&ret, `  %[1]s(tmp.h)[] = addr(vtbl[])
`, ncabiToVdataName(c))
				}
				ret.WriteString("  tmp\n")
			}
		}

		if len(virtualMethods) > 0 {

			fmt.Fprintf(&ret, `const %[1]s_mvtbl = %[1]sVTable(
  destructor: proc(self: pointer) {.cdecl.} =
    let inst = cast[ptr typeof(Virtual%[2]s()[])](self.%[3]s()[])
    inst[].h = nil
    inst[].owned = false,

`, rawClassName, nimClassName, ncabiToVdataName(c))
			for _, m := range virtualMethods {
				fmt.Fprintf(&ret, `  %[1]s: %[2]s,
`, m.rawMethodName(), ncabiMethodCallbackName(c, m))
			}
			ret.WriteString(")\n")

			sigs = map[string]struct{}{}
			for i, ctor := range c.Ctors {
				preamble, forwarding := gfs.emitParametersNim2CABIForwarding(ctor, "addr("+rawClassName+"_mvtbl), csize_t(sizeof(pointer))")

				nimParams := gfs.emitParametersNim(ctor.Parameters, false, "")
				paramsX := ""
				for _, p := range ctor.Parameters {
					paramsX = paramsX + "," + p.renderTypeNim(&gfs, false, true)
				}
				orig := paramsX
				j := 0
				for {
					if _, ok := sigs[paramsX]; ok {
						j += 1
						paramsX = maybeSuffix(j) + orig
					} else {
						sigs[paramsX] = struct{}{}
						break
					}
				}

				vparams := []string{}
				vparams = append(vparams, "T: type "+nimPkgClassName)
				if len(nimParams) > 0 {
					vparams = append(vparams, nimParams)
				}
				vparams = append(vparams, "inst: Virtual"+nimClassName)

				// TODO https://github.com/nim-lang/Nim/issues/24725
				fmt.Fprintf(&ret, `proc create%[1]s*(%[3]s) =
%[4]s  if inst[].h != nil: delete(move(inst[]))
  inst[].h = %[5]s(%[6]s)
  %[7]s(inst[].h)[] = addr inst[]
  inst[].owned = true

`, maybeSuffix(j), nimPkgClassName, strings.Join(vparams, ",\n    "),
					preamble, ncabiNewName(c, i), forwarding, ncabiToVdataName(c))
			}
		}

		for _, p := range c.Props {
			gfs.imports[gfs.qualifiedPrefix("QtCore")+"gen_qobjectdefs_types"] = struct{}{}

			if p.PropertyName == "staticMetaObject" {
				fmt.Fprintf(&cabi, `proc fc%[1]s(): pointer {.importc: "%[1]s".}
`, cabiStaticMetaObjectName(c))

				fmt.Fprintf(&ret, `proc staticMetaObject*(_: type %s): gen_qobjectdefs_types.QMetaObject =
  gen_qobjectdefs_types.QMetaObject(h: fc%s())
`, nimPkgClassName, cabiStaticMetaObjectName(c))
			}
		}
	}

	nimSrc := ret.String()

	// Fixup imports
	for k := range gfs.preImports {
		delete(gfs.imports, k)
	}
	if len(gfs.imports) > 0 {
		allImports := make([]string, 0, len(gfs.imports))
		imports := make([]string, 0, len(gfs.imports))
		for k := range gfs.imports {
			allImports = append(allImports, k)
			imports = append(imports, k)
		}
		for k := range gfs.stdImports {
			allImports = append(allImports, k)
		}

		sort.Strings(imports)
		sort.Strings(allImports)
		exports := []string{}
		for _, s := range imports {
			exports = append(exports, filepath.Base(s))
		}
		nimSrc = strings.Replace(nimSrc, `%%_IMPORTLIBS_%%`, "import\n  "+strings.Join(allImports, ",\n  ")+"\n"+"export\n  "+strings.Join(exports, ",\n  ")+"\n", 1)
	} else {
		nimSrc = strings.Replace(nimSrc, `%%_IMPORTLIBS_%%`, "", 1)
	}

	nimSrc = strings.Replace(nimSrc, `%%_CABI_%%`, cabi.String(), 1)

	typesSrc := types.String()

	return nimSrc, typesSrc, nil
}

func writeNimbleFile(outdir string, version string) {
	src := fmt.Sprintf(`version = "0.%s.0"
license = "MIT"
author = "seaqt"
description = "Generator-based bindings for Qt/QML"
requires "nim >= 2.0.0"
`, version)
	os.WriteFile(filepath.Join(outdir, `seaqt.nimble`), []byte(src), 0644)
}
