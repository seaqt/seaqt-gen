package main

type lookupResultClass struct {
	QtModuleName string
	UnitName     string
	Class        CppClass
}

type lookupResultTypedef struct {
	QtModuleName string
	UnitName     string
	Typedef      CppTypedef
}

type lookupResultEnum struct {
	QtModuleName string
	UnitName     string
	Enum         CppEnum
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

func addKnownTypes(qtModuleName string, parsed *CppParsedHeader) {
	for _, c := range parsed.Classes {
		KnownClassnames[c.ClassName] = lookupResultClass{qtModuleName, genUnitName(parsed.Filename), c /* copy */}
	}
	for _, td := range parsed.Typedefs {
		KnownTypedefs[td.Alias] = lookupResultTypedef{qtModuleName, genUnitName(parsed.Filename), td /* copy */}
	}
	for _, en := range parsed.Enums {
		KnownEnums[en.EnumName] = lookupResultEnum{qtModuleName, genUnitName(parsed.Filename), en /* copy */}
	}
}
