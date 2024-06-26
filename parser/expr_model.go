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
	"fmt"
	"strings"

	"github.com/DeRuneLabs/jane/ast/models"
)

type iExpr interface {
	String() string
}

type exprBuildNode struct {
	nodes []iExpr
}

type exprModel struct {
	index int
	nodes []exprBuildNode
}

func newExprModel(processes []Toks) *exprModel {
	m := new(exprModel)
	m.index = 0
	m.nodes = make([]exprBuildNode, len(processes))
	return m
}

func (m *exprModel) appendSubNode(node iExpr) {
	nodes := &m.nodes[m.index].nodes
	*nodes = append(*nodes, node)
}

func (m exprModel) String() string {
	var expr strings.Builder
	for _, node := range m.nodes {
		for _, node := range node.nodes {
			expr.WriteString(node.String())
		}
	}
	return expr.String()
}

func (m *exprModel) Expr() Expr {
	return Expr{Model: m}
}

type exprNode struct {
	value string
}

func (node exprNode) String() string {
	return node.value
}

type anonFuncExpr struct {
	ast  *Func
	vars []*Var
}

func (af anonFuncExpr) String() string {
	var cpp strings.Builder
	t := DataType{
		Tok:  af.ast.Tok,
		Kind: af.ast.DataTypeString(),
		Tag:  af.ast,
	}
	cpp.WriteString(t.FuncString())
	cpp.WriteString("([")
	if len(af.vars) > 0 {
		var vars strings.Builder
		for _, v := range af.vars {
			id := v.OutId()
			vars.WriteString(id)
			if typeIsPtr(v.Type) {
				vars.WriteByte('=')
				vars.WriteString(id)
				vars.WriteString(".__must_heap()")
			}
			vars.WriteByte(',')
		}
		cpp.WriteString(vars.String()[:vars.Len()-1])
	}
	cpp.WriteByte(']')
	cpp.WriteString(paramsToCpp(af.ast.Params))
	cpp.WriteString(" mutable -> ")
	cpp.WriteString(af.ast.RetType.String())
	cpp.WriteByte(' ')
	cpp.WriteString(af.ast.Block.String())
	cpp.WriteByte(')')
	return cpp.String()
}

type sliceExpr struct {
	dataType DataType
	expr     []iExpr
}

func (a sliceExpr) String() string {
	var cpp strings.Builder
	cpp.WriteString(a.dataType.String())
	cpp.WriteString("({")
	if len(a.expr) == 0 {
		cpp.WriteString("})")
		return cpp.String()
	}
	for _, exp := range a.expr {
		cpp.WriteString(exp.String())
		cpp.WriteByte(',')
	}
	return cpp.String()[:cpp.Len()-1] + "})"
}

type mapExpr struct {
	dataType DataType
	keyExprs []iExpr
	valExprs []iExpr
}

func (m mapExpr) String() string {
	var cpp strings.Builder
	cpp.WriteString(m.dataType.String())
	cpp.WriteByte('{')
	for i, k := range m.keyExprs {
		v := m.valExprs[i]
		cpp.WriteByte('{')
		cpp.WriteString(k.String())
		cpp.WriteByte(',')
		cpp.WriteString(v.String())
		cpp.WriteString("},")
	}
	cpp.WriteByte('}')
	return cpp.String()
}

type genericsExpr struct {
	types []DataType
}

func (ge genericsExpr) String() string {
	if len(ge.types) == 0 {
		return ""
	}
	var cpp strings.Builder
	cpp.WriteByte('<')
	for _, generic := range ge.types {
		cpp.WriteString(generic.String())
		cpp.WriteByte(',')
	}
	return cpp.String()[:cpp.Len()-1] + ">"
}

type argsExpr struct {
	args []models.Arg
}

func (a argsExpr) String() string {
	if len(a.args) == 0 {
		return ""
	}
	var cpp strings.Builder
	for _, arg := range a.args {
		cpp.WriteString(arg.String())
		cpp.WriteByte(',')
	}
	return cpp.String()[:cpp.Len()-1]
}

type callExpr struct {
	generics genericsExpr
	args     argsExpr
}

func (ce callExpr) String() string {
	var cpp strings.Builder
	cpp.WriteString(ce.generics.String())
	cpp.WriteByte('(')
	cpp.WriteString(ce.args.String())
	cpp.WriteByte(')')
	return cpp.String()
}

type retExpr struct {
	models []iExpr
	values []value
}

func (re *retExpr) multiRetString() string {
	var cpp strings.Builder
	cpp.WriteString("std::make_tuple(")
	for i, model := range re.models {
		cpp.WriteString(model.String())
		if typeIsPtr(re.values[i].data.Type) {
			cpp.WriteString(".__must_heap()")
		}
		cpp.WriteByte(',')
	}
	return cpp.String()[:cpp.Len()-1] + ")"
}

func (re *retExpr) singleRetString() string {
	var cpp strings.Builder
	v := re.values[0]
	cpp.WriteString(re.models[0].String())
	if typeIsPtr(v.data.Type) {
		cpp.WriteString(".__must_heap()")
	}
	return cpp.String()
}

func (re retExpr) String() string {
	if len(re.values) > 1 {
		return re.multiRetString()
	}
	return re.singleRetString()
}

type serieExpr struct {
	exprs []any
}

func (se serieExpr) String() string {
	var exprs strings.Builder
	for _, expr := range se.exprs {
		exprs.WriteString(fmt.Sprint(expr))
	}
	return exprs.String()
}
