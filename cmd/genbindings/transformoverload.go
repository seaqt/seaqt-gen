package main

import (
	"fmt"
	"strings"
)

func (p CppParameter) renderTypeForMethod(qualifiers bool) string {
	parameterType := strings.NewReplacer(
		" ", "",
		"::", "_",
		// Generics, ie QMap<A, B>
		"<", "Of",
		">", "",
		",", "_",
		// These appear in generics - putting empty here is not correct if there's
		// an overload for a non-pointer version of the same type but we'd want the
		// p as a prefix to that part of the type.. to be continued!
		"*", "",
	).Replace(p.ParameterType)

	if qualifiers {
		if p.Const {
			parameterType = "c" + parameterType
		}
		if p.Pointer {
			parameterType = "p" + parameterType
		}
	}
	return parameterType
}

// astTransformOverloads renames methods if another method exists with the same
// name.

func renameOverloads(c CppClass, methods *[]CppMethod) bool {
	// Group methods by original name
	anyChange := false
	overloadGroups := map[string][]int{}
	for j, m := range *methods {
		overloadGroups[m.SafeMethodName()] = append(overloadGroups[m.SafeMethodName()], j)
	}

	// For each group, decide naming strategy
	for name, idxs := range overloadGroups {
		if len(idxs) < 2 {
			// Keep original name if there are no overloads
			continue
		}

		useNames := true
		// Check if any method in the group has an empty parameter name
		for _, j := range idxs {
			for _, p := range (*methods)[j].Parameters {
				if strings.HasPrefix(p.ParameterName, "param") {
					// prefer types if there are nameless parameters
					useNames = false
					break
				}
			}
			if !useNames {
				break
			}
		}
		changed := false
		if useNames {
			// Try to disambiguate using parameter names - this works for a lot of
			// of overloads but commonly fails for cases like const/non-const
			// variations or where many "similar" types appear, ie string-likes
			if proposeNames(methods, name, idxs, func(m CppMethod, name string) string {
				return proposeName(m, name, false, true)
			}) {
				changed = true
			}
		}

		if !changed {
			// Use types for disambiguation - try without considering "this" first
			if proposeNames(methods, name, idxs, func(m CppMethod, name string) string {
				return proposeName(m, name, false, false)
			}) {
				changed = true
			}
		}

		if !changed {
			// Differs only by the cv-specifier of "this" - include that too
			if !proposeNames(methods, name, idxs, func(m CppMethod, name string) string {
				return proposeName(m, name, true, false)
			}) {
				panic(fmt.Sprintf("Cannot generate unambiguous name for %v::%v: \n%v", c.ClassName, name, *methods))
			}
		}

		anyChange = true
	}
	return anyChange
}

func proposeName(m CppMethod, basename string, withConst, useNames bool) string {
	proposedName := basename
	if withConst && m.IsConst {
		proposedName = proposedName + "_const"
	}
	if useNames {
		for _, p := range m.Parameters {
			proposedName = proposedName + "_" + p.ParameterName
		}
	} else {
		for _, p := range m.Parameters {
			proposedName = proposedName + "_" + p.renderTypeForMethod(withConst)
		}
	}

	return proposedName
}

func proposeNames(methods *[]CppMethod, name string, idxs []int, strategy func(m CppMethod, name string) string) bool {
	// First pass: assign proposed names using types or parameter names
	proposedNames := make([]string, len(idxs))
	nameCount := map[string]int{}
	for k, j := range idxs {
		m := (*methods)[j]
		proposedName := strategy(m, name)

		proposedNames[k] = proposedName
		nameCount[proposedName]++
	}

	for _, count := range nameCount {
		if count > 1 {
			return false
		}
	}

	for k, j := range idxs {
		m := &(*methods)[j]
		finalName := proposedNames[k]
		m.Rename(finalName)
	}

	return true
}

func astTransformOverloads(parsed *CppParsedHeader) {
	for i, c := range parsed.Classes {
		anyChange := false
		if renameOverloads(c, &c.Ctors) {
			anyChange = true
		}

		if renameOverloads(c, &c.Methods) {
			anyChange = true
		}

		if anyChange {
			parsed.Classes[i] = c
		}
	}
}
