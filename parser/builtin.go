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
	"math"
	"strconv"

	"github.com/DeRuneLabs/jane/ast/models"
	"github.com/DeRuneLabs/jane/lexer/tokens"
	"github.com/DeRuneLabs/jane/package/jn"
	"github.com/DeRuneLabs/jane/package/jntype"
)

var i8statics = &Defmap{
	Globals: []*Var{
		{
			Pub:     true,
			Const:   true,
			Id:      "max",
			Type:    DataType{Id: jntype.I8, Kind: tokens.I8},
			ExprTag: int64(math.MaxInt8),
			Expr: models.Expr{
				Model: exprNode{
					jntype.CppId(jntype.I8) + "{" + strconv.FormatInt(math.MaxInt8, 10) + "}",
				},
			},
		},
		{
			Pub:     true,
			Const:   true,
			Id:      "min",
			Type:    DataType{Id: jntype.I8, Kind: tokens.I8},
			ExprTag: int64(math.MinInt8),
			Expr: models.Expr{
				Model: exprNode{
					jntype.CppId(jntype.I8) + "{" + strconv.FormatInt(math.MinInt8, 10) + "}",
				},
			},
		},
	},
}

var i16statics = &Defmap{
	Globals: []*Var{
		{
			Pub:     true,
			Const:   true,
			Id:      "max",
			Type:    DataType{Id: jntype.I16, Kind: tokens.I16},
			ExprTag: int64(math.MaxInt16),
			Expr: models.Expr{
				Model: exprNode{
					jntype.CppId(jntype.I16) + "{" + strconv.FormatInt(math.MaxInt16, 10) + "}",
				},
			},
		},
		{
			Pub:     true,
			Const:   true,
			Id:      "min",
			Type:    DataType{Id: jntype.I16, Kind: tokens.I16},
			ExprTag: int64(math.MinInt16),
			Expr: models.Expr{
				Model: exprNode{
					jntype.CppId(jntype.I16) + "{" + strconv.FormatInt(math.MinInt16, 10) + "}",
				},
			},
		},
	},
}

var i32statics = &Defmap{
	Globals: []*Var{
		{
			Pub:     true,
			Const:   true,
			Id:      "max",
			Type:    DataType{Id: jntype.I32, Kind: tokens.I32},
			ExprTag: int64(math.MaxInt32),
			Expr: models.Expr{
				Model: exprNode{
					jntype.CppId(jntype.I32) + "{" + strconv.FormatInt(math.MaxInt32, 10) + "}",
				},
			},
		},
		{
			Pub:     true,
			Const:   true,
			Id:      "min",
			Type:    DataType{Id: jntype.I32, Kind: tokens.I32},
			ExprTag: int64(math.MinInt32),
			Expr: models.Expr{
				Model: exprNode{
					jntype.CppId(jntype.I32) + "{" + strconv.FormatInt(math.MinInt32, 10) + "}",
				},
			},
		},
	},
}

var i64statics = &Defmap{
	Globals: []*Var{
		{
			Pub:     true,
			Const:   true,
			Id:      "max",
			Type:    DataType{Id: jntype.I64, Kind: tokens.I64},
			ExprTag: int64(math.MaxInt64),
			Expr: models.Expr{
				Model: exprNode{
					jntype.CppId(jntype.I64) + "{" + strconv.FormatInt(math.MaxInt64, 10) + "}",
				},
			},
		},
		{
			Pub:     true,
			Const:   true,
			Id:      "min",
			Type:    DataType{Id: jntype.I64, Kind: tokens.I64},
			ExprTag: int64(math.MinInt64),
			Expr: models.Expr{
				Model: exprNode{
					jntype.CppId(jntype.I64) + "{" + strconv.FormatInt(math.MinInt64, 10) + "}",
				},
			},
		},
	},
}

var u8statics = &Defmap{
	Globals: []*Var{
		{
			Pub:     true,
			Const:   true,
			Id:      "max",
			Type:    DataType{Id: jntype.U8, Kind: tokens.U8},
			ExprTag: uint64(math.MaxUint8),
			Expr: models.Expr{
				Model: exprNode{
					jntype.CppId(jntype.U8) + "{" + strconv.FormatUint(math.MaxUint8, 10) + "}",
				},
			},
		},
	},
}

var u16statics = &Defmap{
	Globals: []*Var{
		{
			Pub:     true,
			Const:   true,
			Id:      "max",
			Type:    DataType{Id: jntype.U16, Kind: tokens.U16},
			ExprTag: uint64(math.MaxUint16),
			Expr: models.Expr{
				Model: exprNode{
					jntype.CppId(jntype.U16) + "{" + strconv.FormatUint(math.MaxUint16, 10) + "}",
				},
			},
		},
	},
}

var u32statics = &Defmap{
	Globals: []*Var{
		{
			Pub:     true,
			Const:   true,
			Id:      "max",
			Type:    DataType{Id: jntype.U32, Kind: tokens.U32},
			ExprTag: uint64(math.MaxUint32),
			Expr: models.Expr{
				Model: exprNode{
					jntype.CppId(jntype.U32) + "{" + strconv.FormatUint(math.MaxUint32, 10) + "}",
				},
			},
		},
	},
}

var u64statics = &Defmap{
	Globals: []*Var{
		{
			Pub:     true,
			Const:   true,
			Id:      "max",
			Type:    DataType{Id: jntype.U64, Kind: tokens.U64},
			ExprTag: uint64(math.MaxUint64),
			Expr: models.Expr{
				Model: exprNode{
					jntype.CppId(jntype.U64) + "{" + strconv.FormatUint(math.MaxUint64, 10) + "}",
				},
			},
		},
	},
}

var uintStatics = &Defmap{
	Globals: []*Var{
		{
			Pub:   true,
			Const: true,
			Id:    "max",
			Type:  DataType{Id: jntype.UInt, Kind: tokens.UINT},
		},
	},
}

var intStatics = &Defmap{
	Globals: []*Var{
		{
			Const: true,
			Id:    "max",
			Type:  DataType{Id: jntype.Int, Kind: tokens.INT},
		},
		{
			Const: true,
			Id:    "min",
			Type:  DataType{Id: jntype.Int, Kind: tokens.INT},
		},
	},
}

const f32min = float64(1.17549435082228750796873653722224568e-38)

var f32min_model = exprNode{
	jntype.CppId(jntype.F32) + "{1.17549435082228750796873653722224568e-38F}",
}

var f32statics = &Defmap{
	Globals: []*Var{
		{
			Pub:     true,
			Const:   true,
			Id:      "max",
			Type:    DataType{Id: jntype.F32, Kind: tokens.F32},
			ExprTag: float64(math.MaxFloat32),
			Expr: models.Expr{
				Model: exprNode{strconv.FormatFloat(math.MaxFloat32, 'e', -1, 32) + "F"},
			},
		},
		{
			Pub:     true,
			Const:   true,
			Id:      "min",
			Type:    DataType{Id: jntype.F32, Kind: tokens.F32},
			ExprTag: f32min,
			Expr:    models.Expr{Model: f32min_model},
		},
	},
}

const f64min = float64(2.22507385850720138309023271733240406e-308)

var f64min_model = exprNode{
	jntype.CppId(jntype.F64) + "{2.22507385850720138309023271733240406e-308}",
}

var f64statics = &Defmap{
	Globals: []*Var{
		{
			Pub:     true,
			Const:   true,
			Id:      "max",
			Type:    DataType{Id: jntype.F64, Kind: tokens.F64},
			ExprTag: float64(math.MaxFloat64),
			Expr: models.Expr{
				Model: exprNode{strconv.FormatFloat(math.MaxFloat64, 'e', -1, 64)},
			},
		},
		{
			Pub:     true,
			Const:   true,
			Id:      "min",
			Type:    DataType{Id: jntype.F64, Kind: tokens.F64},
			ExprTag: f64min,
			Expr:    models.Expr{Model: f64min_model},
		},
	},
}

var strDefaultFunc = Func{
	Pub:     true,
	Id:      "str",
	Params:  []Param{{Id: "obj", Type: DataType{Id: jntype.Any, Kind: tokens.ANY}}},
	RetType: RetType{Type: DataType{Id: jntype.Str, Kind: tokens.STR}},
}

var errorTrait = &trait{
	Ast: &models.Trait{
		Id: "Error",
	},
	Defs: &Defmap{
		Funcs: []*function{
			{Ast: &models.Func{
				Pub:     true,
				Id:      "error",
				RetType: models.RetType{Type: DataType{Id: jntype.Str, Kind: tokens.STR}},
			}},
		},
	},
}

var errorType = DataType{
	Id:   jntype.Trait,
	Kind: errorTrait.Ast.Id,
	Tag:  errorTrait,
	Pure: true,
}

var panicFunc = &function{
	Ast: &models.Func{
		Pub: true,
		Id:  "panic",
		Params: []models.Param{
			{
				Id:   "error",
				Type: errorType,
			},
		},
	},
}

var errorHandlerFunc = &models.Func{
	Id: "handler",
	Params: []models.Param{
		{
			Id:   "error",
			Type: errorType,
		},
	},
	RetType: models.RetType{
		Type: models.DataType{
			Id:   jntype.Void,
			Kind: jntype.TypeMap[jntype.Void],
		},
	},
}

var recoverFunc = &function{
	Ast: &models.Func{
		Pub: true,
		Id:  "recover",
		Params: []models.Param{
			{
				Id: "handler",
				Type: models.DataType{
					Id:   jntype.Func,
					Kind: errorHandlerFunc.DataTypeString(),
					Tag:  errorHandlerFunc,
				},
			},
		},
	},
}

var genericFile = &Parser{}

var Builtin = &Defmap{
	Types: []*models.Type{
		{
			Pub:  true,
			Id:   "byte",
			Type: DataType{Id: jntype.U8, Kind: jntype.TypeMap[jntype.U8]},
		},
		{
			Pub:  true,
			Id:   "rune",
			Type: DataType{Id: jntype.I32, Kind: jntype.TypeMap[jntype.I32]},
		},
	},
	Funcs: []*function{
		panicFunc,
		recoverFunc,
		{Ast: &Func{
			Pub: true,
			Id:  "print",
			RetType: RetType{
				Type: DataType{Id: jntype.Void, Kind: jntype.TypeMap[jntype.Void]},
			},
			Params: []Param{{
				Id:   "expr",
				Type: DataType{Id: jntype.Any, Kind: tokens.ANY},
			}},
		}},
		{Ast: &Func{
			Pub:      true,
			Id:       "new",
			Owner:    genericFile,
			Generics: []*GenericType{{Id: "T"}},
			Attributes: []Attribute{
				models.Attribute{Tag: jn.Attribute_TypeArg},
			},
			RetType: RetType{Type: DataType{Id: jntype.Id, Kind: tokens.STAR + "T"}},
		}},
		{Ast: &Func{
			Pub:      true,
			Id:       "make",
			Owner:    genericFile,
			Generics: []*GenericType{{Id: "Item"}},
			RetType: models.RetType{
				Type: DataType{
					Id:            jntype.Slice,
					Kind:          jn.Prefix_Slice + "Item",
					ComponentType: &DataType{Id: jntype.Id, Kind: "Item"},
				},
			},
			Params: []models.Param{
				{
					Id:   "n",
					Type: DataType{Id: jntype.Int, Kind: jntype.TypeMap[jntype.Int]},
				},
			},
		}},
		{Ast: &Func{
			Pub:      true,
			Id:       "copy",
			Owner:    genericFile,
			Generics: []*GenericType{{Id: "Item"}},
			RetType: models.RetType{
				Type: DataType{Id: jntype.Int, Kind: jntype.TypeMap[jntype.Int]},
			},
			Params: []models.Param{
				{
					Id: "dest",
					Type: DataType{
						Id:            jntype.Slice,
						Kind:          jn.Prefix_Slice + "Item",
						ComponentType: &DataType{Id: jntype.Id, Kind: "Item"},
					},
				},
				{
					Id: "src",
					Type: DataType{
						Id:            jntype.Slice,
						Kind:          jn.Prefix_Slice + "Item",
						ComponentType: &DataType{Id: jntype.Id, Kind: "Item"},
					},
				},
			},
		}},
		{Ast: &Func{
			Pub:      true,
			Id:       "append",
			Owner:    genericFile,
			Generics: []*GenericType{{Id: "Item"}},
			RetType: models.RetType{
				Type: DataType{
					Id:            jntype.Slice,
					Kind:          jn.Prefix_Slice + "Item",
					ComponentType: &DataType{Id: jntype.Id, Kind: "Item"},
				},
			},
			Params: []models.Param{
				{
					Id: "src",
					Type: DataType{
						Id:            jntype.Slice,
						Kind:          jn.Prefix_Slice + "Item",
						ComponentType: &DataType{Id: jntype.Id, Kind: "Item"},
					},
				},
				{
					Id:       "components",
					Type:     DataType{Id: jntype.Id, Kind: "Item"},
					Variadic: true,
				},
			},
		}},
	},
	Traits: []*trait{
		errorTrait,
	},
}

var strDefs = &Defmap{
	Globals: []*Var{
		{
			Pub:  true,
			Id:   "len",
			Type: DataType{Id: jntype.Int, Kind: tokens.INT},
			Tag:  "len()",
		},
	},
	Funcs: []*function{
		{Ast: &Func{
			Pub:     true,
			Id:      "empty",
			RetType: RetType{Type: DataType{Id: jntype.Bool, Kind: tokens.BOOL}},
		}},
		{Ast: &Func{
			Pub:     true,
			Id:      "has_prefix",
			Params:  []Param{{Id: "sub", Type: DataType{Id: jntype.Str, Kind: tokens.STR}}},
			RetType: RetType{Type: DataType{Id: jntype.Bool, Kind: tokens.BOOL}},
		}},
		{Ast: &Func{
			Pub:     true,
			Id:      "has_suffix",
			Params:  []Param{{Id: "sub", Type: DataType{Id: jntype.Str, Kind: tokens.STR}}},
			RetType: RetType{Type: DataType{Id: jntype.Bool, Kind: tokens.BOOL}},
		}},
		{Ast: &Func{
			Pub:     true,
			Id:      "find",
			Params:  []Param{{Id: "sub", Type: DataType{Id: jntype.Str, Kind: tokens.STR}}},
			RetType: RetType{Type: DataType{Id: jntype.Int, Kind: tokens.INT}},
		}},
		{Ast: &Func{
			Pub:     true,
			Id:      "rfind",
			Params:  []Param{{Id: "sub", Type: DataType{Id: jntype.Str, Kind: tokens.STR}}},
			RetType: RetType{Type: DataType{Id: jntype.Int, Kind: tokens.INT}},
		}},
		{Ast: &Func{
			Pub:     true,
			Id:      "trim",
			Params:  []Param{{Id: "bytes", Type: DataType{Id: jntype.Str, Kind: tokens.STR}}},
			RetType: RetType{Type: DataType{Id: jntype.Str, Kind: tokens.STR}},
		}},
		{Ast: &Func{
			Pub:     true,
			Id:      "rtrim",
			Params:  []Param{{Id: "bytes", Type: DataType{Id: jntype.Str, Kind: tokens.STR}}},
			RetType: RetType{Type: DataType{Id: jntype.Str, Kind: tokens.STR}},
		}},
		{Ast: &Func{
			Pub: true,
			Id:  "split",
			Params: []Param{
				{Id: "sub", Type: DataType{Id: jntype.Str, Kind: tokens.STR}},
				{
					Id:   "n",
					Type: DataType{Id: jntype.Int, Kind: tokens.INT},
				},
			},
			RetType: RetType{Type: DataType{Id: jntype.Str, Kind: jn.Prefix_Slice + tokens.STR}},
		}},
		{Ast: &Func{
			Pub: true,
			Id:  "replace",
			Params: []Param{
				{Id: "sub", Type: DataType{Id: jntype.Str, Kind: tokens.STR}},
				{Id: "new", Type: DataType{Id: jntype.Str, Kind: tokens.STR}},
				{
					Id:   "n",
					Type: DataType{Id: jntype.Int, Kind: tokens.INT},
				},
			},
			RetType: RetType{Type: DataType{Id: jntype.Str, Kind: tokens.STR}},
		}},
	},
}

var sliceDefs = &Defmap{
	Globals: []*Var{
		{
			Pub:  true,
			Id:   "len",
			Type: DataType{Id: jntype.Int, Kind: tokens.INT},
			Tag:  "len()",
		},
	},
	Funcs: []*function{
		{Ast: &Func{
			Pub:     true,
			Id:      "empty",
			RetType: RetType{Type: DataType{Id: jntype.Bool, Kind: tokens.BOOL}},
		}},
	},
}

var arrayDefs = &Defmap{
	Globals: []*Var{
		{
			Pub:  true,
			Id:   "len",
			Type: DataType{Id: jntype.Int, Kind: tokens.INT},
			Tag:  "len()",
		},
	},
	Funcs: []*function{
		{Ast: &Func{
			Pub:     true,
			Id:      "empty",
			RetType: RetType{Type: DataType{Id: jntype.Bool, Kind: tokens.BOOL}},
		}},
	},
}

var mapDefs = &Defmap{
	Globals: []*Var{
		{
			Pub:  true,
			Id:   "len",
			Type: DataType{Id: jntype.Int, Kind: tokens.INT},
			Tag:  "len()",
		},
	},
	Funcs: []*function{
		{Ast: &Func{
			Pub: true,
			Id:  "clear",
		}},
		{Ast: &Func{
			Pub: true,
			Id:  "keys",
		}},
		{Ast: &Func{
			Pub: true,
			Id:  "values",
		}},
		{Ast: &Func{
			Pub:     true,
			Id:      "empty",
			RetType: RetType{Type: DataType{Id: jntype.Bool, Kind: tokens.BOOL}},
		}},
		{Ast: &Func{
			Pub:     true,
			Id:      "has",
			Params:  []Param{{Id: "key"}},
			RetType: RetType{Type: DataType{Id: jntype.Bool, Kind: tokens.BOOL}},
		}},
		{Ast: &Func{
			Pub:    true,
			Id:     "del",
			Params: []Param{{Id: "key"}},
		}},
	},
}

func readyMapDefs(mapt DataType) {
	types := mapt.Tag.([]DataType)
	keyt := types[0]
	valt := types[1]

	keysFunc, _, _ := mapDefs.funcById("keys", nil)
	keysFunc.Ast.RetType.Type = keyt
	keysFunc.Ast.RetType.Type.Kind = jn.Prefix_Slice + keysFunc.Ast.RetType.Type.Kind

	valuesFunc, _, _ := mapDefs.funcById("values", nil)
	valuesFunc.Ast.RetType.Type = valt
	valuesFunc.Ast.RetType.Type.Kind = jn.Prefix_Slice + valuesFunc.Ast.RetType.Type.Kind

	hasFunc, _, _ := mapDefs.funcById("has", nil)
	hasFunc.Ast.Params[0].Type = keyt

	delFunc, _, _ := mapDefs.funcById("del", nil)
	delFunc.Ast.Params[0].Type = keyt
}

func init() {
	printFunc, _, _ := Builtin.funcById("print", nil)
	printlnFunc := new(function)
	*printlnFunc = *printFunc
	printlnFunc.Ast = new(models.Func)
	*printlnFunc.Ast = *printFunc.Ast
	printlnFunc.Ast.Id = "println"
	Builtin.Funcs = append(Builtin.Funcs, printlnFunc)

	intMax := intStatics.Globals[0]
	intMin := intStatics.Globals[1]
	uintMax := uintStatics.Globals[0]
	switch jntype.BitSize {
	case 8:
		intMax.Expr = i8statics.Globals[0].Expr
		intMax.ExprTag = i8statics.Globals[0].ExprTag
		intMin.Expr = i8statics.Globals[1].Expr
		intMin.ExprTag = i8statics.Globals[1].ExprTag

		uintMax.Expr = u8statics.Globals[0].Expr
		uintMax.ExprTag = u8statics.Globals[0].ExprTag
	case 16:
		intMax.Expr = i16statics.Globals[0].Expr
		intMax.ExprTag = i16statics.Globals[0].ExprTag
		intMin.Expr = i16statics.Globals[1].Expr
		intMin.ExprTag = i16statics.Globals[1].ExprTag

		uintMax.Expr = u16statics.Globals[0].Expr
		uintMax.ExprTag = u16statics.Globals[0].ExprTag
	case 32:
		intMax.Expr = i32statics.Globals[0].Expr
		intMax.ExprTag = i32statics.Globals[0].ExprTag
		intMin.Expr = i32statics.Globals[1].Expr
		intMin.ExprTag = i32statics.Globals[1].ExprTag

		uintMax.Expr = u32statics.Globals[0].Expr
		uintMax.ExprTag = u32statics.Globals[0].ExprTag
	case 64:
		intMax.Expr = i64statics.Globals[0].Expr
		intMax.ExprTag = i64statics.Globals[0].ExprTag
		intMin.Expr = i64statics.Globals[1].Expr
		intMin.ExprTag = i64statics.Globals[1].ExprTag

		uintMax.Expr = u64statics.Globals[0].Expr
		uintMax.ExprTag = u64statics.Globals[0].ExprTag
	}
}
