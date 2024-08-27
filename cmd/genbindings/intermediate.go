package main

import (
	"regexp"
	"strings"
)

var (
	KnownClassnames map[string]struct{} // Entries of the form QFoo::Bar if it is an inner class
	KnownTypedefs   map[string]CppTypedef
)

func init() {
	KnownClassnames = make(map[string]struct{})
	KnownTypedefs = make(map[string]CppTypedef)

	// Seed well-known typedefs

	KnownTypedefs["QRgb"] = CppTypedef{"QRgb", parseSingleTypeString("unsigned int")}

	KnownTypedefs["WId"] = CppTypedef{"WId", parseSingleTypeString("uintptr_t")}

	// This is a uint64 PID on Linux/mac and a PROCESS_INFORMATION* on Windows
	// A uintptr should be tolerable for both cases until we do better
	// @ref https://doc.qt.io/qt-5/qprocess.html#Q_PID-typedef
	KnownTypedefs["Q_PID"] = CppTypedef{"WId", parseSingleTypeString("uintptr_t")}
}

type CppParameter struct {
	ParameterName string
	ParameterType string
	TypeAlias     string // If we rewrote QStringList->QList<String>, this field contains the original QStringList
	Const         bool
	Pointer       bool
	ByRef         bool
	Optional      bool
}

func (p *CppParameter) AssignAlias(newType string) {
	if p.TypeAlias == "" {
		p.TypeAlias = p.ParameterType // Overwrite once only, at the earliest base type
	}
	p.ParameterType = newType
}

func (p *CppParameter) CopyWithAlias(alias CppParameter) CppParameter {
	ret := *p // copy
	ret.ParameterName = alias.ParameterName
	ret.TypeAlias = alias.ParameterType
	return ret
}

func (p *CppParameter) UnderlyingType() string {
	if p.TypeAlias != "" {
		return p.TypeAlias
	}

	return p.ParameterType
}

func (p CppParameter) QtClassType() bool {
	if p.ParameterType[0] != 'Q' {
		return false
	}

	if strings.Contains(p.ParameterType, `::`) {
		// Maybe if it's an inner class
		if _, ok := KnownClassnames[p.ParameterType]; ok {
			return true
		}
		// Int type
		return false
	}

	// Passed all conditions
	return true
}

func (p CppParameter) QListOf() (CppParameter, bool) {
	if strings.HasPrefix(p.ParameterType, "QList<") && strings.HasSuffix(p.ParameterType, `>`) {
		return parseSingleTypeString(p.ParameterType[6 : len(p.ParameterType)-1]), true
	}

	if strings.HasPrefix(p.ParameterType, "QVector<") && strings.HasSuffix(p.ParameterType, `>`) {
		return parseSingleTypeString(p.ParameterType[8 : len(p.ParameterType)-1]), true
	}

	return CppParameter{}, false
}

func (p CppParameter) QMapOf() bool {
	return strings.HasPrefix(p.ParameterType, `QMap<`) ||
		strings.HasPrefix(p.ParameterType, `QHash<`) // TODO support this
}

func (p CppParameter) QPairOf() bool {
	return strings.HasPrefix(p.ParameterType, `QPair<`) // TODO support this
}

func (p CppParameter) QSetOf() bool {
	return strings.HasPrefix(p.ParameterType, `QSet<`) // TODO support this
}

func (p CppParameter) IntType() bool {

	if strings.Contains(p.ParameterType, `::`) {
		return true
	}

	switch p.ParameterType {
	case "int", "unsigned int", "uint",
		"short", "unsigned short", "ushort", "qint16", "quint16",
		"qint8", "quint8",
		"unsigned char", "uchar",
		"long", "unsigned long", "ulong", "qint32", "quint32",
		"longlong", "ulonglong", "qlonglong", "qulonglong", "qint64", "quint64", "int64_t", "uint64_t", "long long", "unsigned long long",
		"qintptr", "quintptr", "uintptr_t", "intptr_t",
		"qsizetype", "size_t",
		"qptrdiff", "ptrdiff_t",
		"double", "float", "qreal":
		return true

	case "char":
		// Only count char as an integer type with cast assertions if it's
		// not possibly a char* string in disguise
		// (However, unsigned chars are always like ints)
		return !p.Pointer

	default:
		return false
	}
}

type CppProperty struct {
	PropertyName string
	PropertyType string
	Visibility   string
}

type CppMethod struct {
	MethodName         string       // C++ method name, unless OverrideMethodName is set, in which case a nice alternative name
	OverrideMethodName string       // C++ method name, present only if we changed the target
	ReturnType         CppParameter // Name not used
	Parameters         []CppParameter
	IsStatic           bool
	IsSignal           bool
	HasHiddenParams    bool // Set to true if there is an overload with more parameters
}

func IsArgcArgv(params []CppParameter, pos int) bool {
	// Check if the arguments starting at position=pos are the argc/argv pattern.
	// QApplication/QGuiApplication constructors are the only expected example of this.
	return (len(params) > pos+1 &&
		params[pos].ParameterName == "argc" &&
		params[pos].ParameterType == "int" &&
		params[pos].ByRef &&
		params[pos+1].ParameterName == "argv" &&
		params[pos+1].ParameterType == "char **")
}

func IsReceiverMethod(params []CppParameter, pos int) bool {
	// Check if the arguments starting at position=pos are the receiver/member pattern.
	// QMenu->addAction is the main example of this
	return (len(params) > pos+1 &&
		params[pos].ParameterName == "receiver" &&
		params[pos].ParameterType == "QObject" &&
		params[pos].Pointer &&
		params[pos+1].ParameterName == "member" &&
		params[pos+1].ParameterType == "char" &&
		params[pos+1].Pointer)
}

func (nm CppMethod) IsReceiverMethod() bool {
	// Returns true if any of the parameters use the receiever-method pattern
	for i := 0; i < len(nm.Parameters); i++ {
		if IsReceiverMethod(nm.Parameters, i) {
			return true
		}
	}
	return false
}

func (nm CppMethod) SafeMethodName() string {

	tmp := nm.MethodName

	// Strip redundant Qt prefix, we know these are all Qt functions
	if strings.HasPrefix(tmp, "qt_") {
		tmp = tmp[3:]
	}

	// Operator-overload methods have names not representable in binding
	// languages. Replace more specific cases first
	replacer := strings.NewReplacer(

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

	// Also make the first letter uppercase so it becomes public in Go
	tmp = titleCase(tmp)

	// Also replace any underscore_case with CamelCase
	tmp = regexp.MustCompile(`_([a-z])`).ReplaceAllStringFunc(tmp, func(match string) string { return strings.ToUpper(match[1:]) })

	return tmp
}

type CppEnumEntry struct {
	EntryName  string
	EntryValue string
}

type CppEnum struct {
	EnumName string
	Entries  []CppEnumEntry
}

type CppClass struct {
	ClassName string
	Abstract  bool
	Ctors     []CppMethod // only use the parameters
	Inherits  []string    // other class names
	Methods   []CppMethod
	Props     []CppProperty
	CanDelete bool

	ChildTypedefs  []CppTypedef
	ChildClassdefs []CppClass
	ChildEnums     []CppEnum
}

type CppTypedef struct {
	Alias          string
	UnderlyingType CppParameter
}

type CppParsedHeader struct {
	Filename string
	Typedefs []CppTypedef
	Enums    []CppEnum
	Classes  []CppClass
}

func (c CppParsedHeader) Empty() bool {
	return len(c.Typedefs) == 0 &&
		len(c.Classes) == 0
}

func (c *CppParsedHeader) AddContentFrom(other *CppParsedHeader) {
	c.Classes = append(c.Classes, other.Classes...)
	c.Enums = append(c.Enums, other.Enums...)
	c.Typedefs = append(c.Typedefs, other.Typedefs...)
}
