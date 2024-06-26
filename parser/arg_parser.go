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
	"github.com/DeRuneLabs/jane/ast/models"
	"github.com/DeRuneLabs/jane/lexer/tokens"
	"github.com/DeRuneLabs/jane/package/jn"
	"github.com/DeRuneLabs/jane/package/jntype"
)

func getParamMap(params []Param) *paramMap {
	pmap := new(paramMap)
	*pmap = make(paramMap, len(params))
	for i := range params {
		param := &params[i]
		(*pmap)[param.Id] = &paramMapPair{param, nil}
	}
	return pmap
}

type pureArgParser struct {
	p       *Parser
	pmap    *paramMap
	f       *Func
	args    *models.Args
	i       int
	arg     Arg
	errTok  Tok
	m       *exprModel
	paramId string
}

func (pap *pureArgParser) buildArgs() {
	pap.args.Src = make([]Arg, len(*pap.pmap))
	for i, p := range pap.f.Params {
		pair := (*pap.pmap)[p.Id]
		switch {
		case pair.arg != nil:
			pap.args.Src[i] = *pair.arg
		case pair.param.Variadic:
			t := DataType{
				Id:            jntype.Slice,
				Tok:           pair.param.Type.Tok,
				Kind:          jn.Prefix_Slice + pair.param.Type.Kind,
				Pure:          true,
				ComponentType: new(DataType),
			}
			*t.ComponentType = pair.param.Type
			model := sliceExpr{t, nil}
			arg := Arg{Expr: Expr{Model: model}}
			pap.args.Src[i] = arg
		}
	}
}

func (pap *pureArgParser) pushVariadicArgs(pair *paramMapPair) {
	var model serieExpr
	variadiced := false
	pap.p.parseArg(pap.f, pair, pap.args, &variadiced)
	model.exprs = append(model.exprs, pair.arg.Expr.Model.(iExpr))
	once := false
	for pap.i++; pap.i < len(pap.args.Src); pap.i++ {
		pair.arg = &pap.args.Src[pap.i]
		once = true
		pap.p.parseArg(pap.f, pair, pap.args, &variadiced)
		model.exprs = append(model.exprs, exprNode{tokens.COMMA})
		model.exprs = append(model.exprs, pair.arg.Expr.Model.(iExpr))
	}
	model.exprs = append(model.exprs, exprNode{tokens.RBRACE})
	pair.arg.Expr.Model = model
	if !once {
		return
	}
	if variadiced {
		pap.p.pusherrtok(pap.errTok, "more_args_with_variadiced")
	}
}

func (pap *pureArgParser) checkPasses() {
	for _, pair := range *pap.pmap {
		if pair.arg == nil && !pair.param.Variadic {
			pap.p.pusherrtok(pap.errTok, "missing_expr_for", pair.param.Id)
		}
	}
}

func (pap *pureArgParser) pushArg() {
	defer func() { pap.i++ }()
	pair := (*pap.pmap)[pap.paramId]
	arg := pap.arg
	pair.arg = &arg
	if pair.param.Variadic {
		pap.pushVariadicArgs(pair)
	} else {
		pap.p.parseArg(pap.f, pair, pap.args, nil)
	}
}

func (pap *pureArgParser) parse() {
	if len(pap.args.Src) < len(pap.f.Params) {
		if len(pap.args.Src) == 1 {
			if pap.tryFuncMultiRetAsArgs() {
				return
			}
		}
	}
	pap.pmap = getParamMap(pap.f.Params)
	argCount := 0
	for pap.i < len(pap.args.Src) {
		if argCount >= len(pap.f.Params) {
			pap.p.pusherrtok(pap.errTok, "argument_overflow")
			return
		}
		argCount++
		pap.arg = pap.args.Src[pap.i]
		pap.paramId = pap.f.Params[pap.i].Id
		pap.pushArg()
	}
	pap.checkPasses()
	pap.buildArgs()
}

func (pap *pureArgParser) tryFuncMultiRetAsArgs() bool {
	arg := pap.args.Src[0]
	val, model := pap.p.evalExpr(arg.Expr)
	arg.Expr.Model = model
	if !val.data.Type.MultiTyped {
		return false
	}
	types := val.data.Type.Tag.([]DataType)
	if len(types) < len(pap.f.Params) {
		return false
	} else if len(types) > len(pap.f.Params) {
		return false
	}
	if pap.m != nil {
		fname := pap.m.nodes[pap.m.index].nodes[0]
		pap.m.nodes[pap.m.index].nodes[0] = exprNode{"tuple_as_args"}
		pap.args.Src = make([]Arg, 2)
		pap.args.Src[0] = Arg{Expr: Expr{Model: fname}}
		pap.args.Src[1] = arg
	}
	for i, param := range pap.f.Params {
		rt := types[i]
		pap.p.wg.Add(1)
		val := value{data: models.Data{Type: rt}}
		pap.p.checkArgType(&param, val, arg.Tok)
	}
	return true
}
