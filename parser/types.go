// Copyright (c) 2024 - DeRuneLabs
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package parser

import (
	"strings"

	"github.com/DeRuneLabs/jane/lexer/tokens"
	"github.com/DeRuneLabs/jane/package/jn"
	"github.com/DeRuneLabs/jane/package/jntype"
)

func findGeneric(id string, generics []*GenericType) *GenericType {
	for _, generic := range generics {
		if generic.Id == id {
			return generic
		}
	}
	return nil
}

func typeIsVoid(t DataType) bool {
	return t.Id == jntype.Void && !t.MultiTyped
}

func typeIsVariadicable(t DataType) bool {
	return typeIsSlice(t)
}

func typeIsAllowForConst(t DataType) bool {
	if !typeIsPure(t) {
		return false
	}
	return t.Id == jntype.Str || jntype.IsNumeric(t.Id)
}

func typeIsStruct(dt DataType) bool {
	return dt.Id == jntype.Struct
}

func typeIsTrait(dt DataType) bool {
	return dt.Id == jntype.Trait
}

func typeIsEnum(dt DataType) bool {
	return dt.Id == jntype.Enum
}

func unptrType(t DataType) DataType {
	t.Kind = t.Kind[1:]
	return t
}

func typeHasThisGeneric(generic *GenericType, t DataType) bool {
	switch {
	case typeIsFunc(t):
		f := t.Tag.(*Func)
		for _, p := range f.Params {
			if typeHasThisGeneric(generic, p.Type) {
				return true
			}
		}
		return typeHasThisGeneric(generic, f.RetType.Type)
	case t.MultiTyped, typeIsMap(t):
		types := t.Tag.([]DataType)
		for _, t := range types {
			if typeHasThisGeneric(generic, t) {
				return true
			}
		}
		return false
	case typeIsSlice(t), typeIsArray(t):
		return typeHasThisGeneric(generic, *t.ComponentType)
	}
	return typeIsThisGeneric(generic, t)
}

func typeHasGenerics(generics []*GenericType, t DataType) bool {
	for _, generic := range generics {
		if typeHasThisGeneric(generic, t) {
			return true
		}
	}
	return false
}

func typeIsThisGeneric(generic *GenericType, t DataType) bool {
	id, _ := t.KindId()
	return id == generic.Id
}

func typeIsGeneric(generics []*GenericType, t DataType) bool {
	if t.Id != jntype.Id {
		return false
	}
	for _, generic := range generics {
		if typeIsThisGeneric(generic, t) {
			return true
		}
	}
	return false
}

func typeIsExplicitPtr(t DataType) bool {
	if t.Kind == "" {
		return false
	}
	return t.Kind[0] == '*'
}

func typeIsPtr(t DataType) bool {
	return typeIsExplicitPtr(t)
}

func typeIsSlice(t DataType) bool {
	return t.Id == jntype.Slice && strings.HasPrefix(t.Kind, jn.Prefix_Slice)
}

func typeIsArray(t DataType) bool {
	return t.Id == jntype.Array && strings.HasPrefix(t.Kind, jn.Prefix_Array)
}

func typeIsMap(t DataType) bool {
	if t.Kind == "" || t.Id != jntype.Map {
		return false
	}
	return t.Kind[0] == '[' && t.Kind[len(t.Kind)-1] == ']'
}

func typeIsFunc(t DataType) bool {
	if t.Id != jntype.Func || t.Kind == "" {
		return false
	}
	return t.Kind[0] == '('
}

func typeIsPure(t DataType) bool {
	return !typeIsPtr(t) &&
		!typeIsSlice(t) &&
		!typeIsArray(t) &&
		!typeIsMap(t) &&
		!typeIsFunc(t)
}

func subIdAccessorOfType(t DataType) string {
	if typeIsPtr(t) {
		return "->"
	}
	return tokens.DOT
}

func typeIsNilCompatible(t DataType) bool {
	return t.Id == jntype.Nil ||
		typeIsFunc(t) ||
		typeIsPtr(t) ||
		typeIsSlice(t) ||
		typeIsTrait(t) ||
		typeIsMap(t)
}

func checkSliceCompatiblity(arrT, t DataType) bool {
	if t.Id == jntype.Nil {
		return true
	}
	return arrT.Kind == t.Kind
}

func checkArrayCompatiblity(arrT, t DataType) bool {
	if !typeIsArray(t) {
		return false
	}
	return arrT.Size.N == t.Size.N
}

func checkMapCompability(mapT, t DataType) bool {
	if t.Id == jntype.Nil {
		return true
	}
	return mapT.Kind == t.Kind
}

func typeIsLvalue(t DataType) bool {
	return typeIsPtr(t) || typeIsSlice(t) || typeIsMap(t)
}

func checkPtrCompability(t1, t2 DataType) bool {
	if t2.Id == jntype.Nil {
		return true
	}
	return t1.Kind == t2.Kind
}

func typesEquals(t1, t2 DataType) bool {
	return t1.Id == t2.Id && t1.Kind == t2.Kind
}

func checkTraitCompability(t1, t2 DataType) bool {
	t := t1.Tag.(*trait)
	switch {
	case typeIsTrait(t2):
		return t == t2.Tag.(*trait)
	case typeIsStruct(t2):
		s := t2.Tag.(*jnstruct)
		return s.hasTrait(t)
	}
	return false
}

func checkStructCompability(t1, t2 DataType) bool {
	s1, s2 := t1.Tag.(*jnstruct), t2.Tag.(*jnstruct)
	switch {
	case s1.Ast.Id != s2.Ast.Id,
		s1.Ast.Tok.File != s2.Ast.Tok.File:
		return false
	}
	if len(s1.Ast.Generics) == 0 {
		return true
	}
	n1, n2 := len(s1.generics), len(s2.generics)
	if n1 != n2 {
		return false
	}
	for i, g1 := range s1.generics {
		g2 := s2.generics[i]
		if !typesEquals(g1, g2) {
			return false
		}
	}
	return true
}

func typesAreCompatible(t1, t2 DataType, ignoreany bool) bool {
	switch {
	case typeIsPtr(t1), typeIsPtr(t2):
		if typeIsPtr(t2) {
			t1, t2 = t2, t1
		}
		return checkPtrCompability(t1, t2)
	case typeIsSlice(t1), typeIsSlice(t2):
		if typeIsSlice(t2) {
			t1, t2 = t2, t1
		}
		return checkSliceCompatiblity(t1, t2)
	case typeIsArray(t1), typeIsArray(t2):
		if typeIsArray(t2) {
			t1, t2 = t2, t1
		}
		return checkArrayCompatiblity(t1, t2)
	case typeIsMap(t1), typeIsMap(t2):
		if typeIsMap(t2) {
			t1, t2 = t2, t1
		}
		return checkMapCompability(t1, t2)
	case typeIsTrait(t1), typeIsTrait(t2):
		if typeIsTrait(t2) {
			t1, t2 = t2, t1
		}
		return checkTraitCompability(t1, t2)
	case typeIsNilCompatible(t1):
		return t2.Id == jntype.Nil
	case typeIsNilCompatible(t2):
		return t1.Id == jntype.Nil
	case typeIsEnum(t1), typeIsEnum(t2):
		return t1.Id == t2.Id && t1.Kind == t2.Kind
	case typeIsStruct(t1), typeIsStruct(t2):
		if t2.Id == jntype.Struct {
			t1, t2 = t2, t1
		}
		return checkStructCompability(t1, t2)
	}
	return jntype.TypesAreCompatible(t1.Id, t2.Id, ignoreany)
}
