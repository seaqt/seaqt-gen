package main

type lookupResultClass struct {
	PackageName string
	UnitName    string
	Class       CppClass
}

type lookupResultTypedef struct {
	PackageName string
	UnitName    string
	Typedef     CppTypedef
}

type lookupResultEnum struct {
	PackageName string
	UnitName    string
	Enum        CppEnum
}

var (
	KnownClassnames map[string]lookupResultClass // Entries of the form QFoo::Bar if it is an inner class
	KnownTypedefs   map[string]lookupResultTypedef
	KnownEnums      map[string]lookupResultEnum
)

func flushKnownTypes() {
	KnownClassnames = make(map[string]lookupResultClass)
	KnownTypedefs = make(map[string]lookupResultTypedef)
	KnownEnums = make(map[string]lookupResultEnum)
}

func addKnownTypes(packageName string, parsed *CppParsedHeader) {
	for _, c := range parsed.Classes {
		KnownClassnames[c.ClassName] = lookupResultClass{packageName, genUnitName(parsed.Filename), c /* copy */}
	}
	for _, td := range parsed.Typedefs {
		KnownTypedefs[td.Alias] = lookupResultTypedef{packageName, genUnitName(parsed.Filename), td /* copy */}
	}
	for _, en := range parsed.Enums {
		KnownEnums[en.EnumName] = lookupResultEnum{packageName, genUnitName(parsed.Filename), en /* copy */}
	}
}
