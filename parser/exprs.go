package parser

import (
	"strings"

	"github.com/DeRuneLabs/jane/ast/models"
	"github.com/DeRuneLabs/jane/lexer/tokens"
	"github.com/DeRuneLabs/jane/package/jnbits"
	"github.com/DeRuneLabs/jane/package/jntype"
)

func arrayExprIsAutoSized(expr models.Expr) bool {
	return len(expr.Processes) == 0 || len(expr.Toks) == 0
}

func isstr(s string) bool {
	return s != "" && (s[0] == '"' || israwstr(s))
}

func israwstr(s string) bool {
	return s != "" && s[0] == '`'
}

func ischar(s string) bool {
	return s != "" && s[0] == '\''
}

func isnil(s string) bool {
	return s == tokens.NIL
}

func isbool(s string) bool {
	return s == tokens.TRUE || s == tokens.FALSE
}

func valIsEnumType(v value) bool {
	return v.isType && typeIsEnum(v.data.Type)
}

func isBoolExpr(v value) bool {
	return typeIsPure(v.data.Type) && v.data.Type.Id == jntype.Bool
}

func isfloat(s string) bool {
	if strings.HasPrefix(s, "0x") {
		return false
	}
	return strings.Contains(s, tokens.DOT) ||
		strings.ContainsAny(s, "eE")
}

func canGetPtr(v value) bool {
	if !v.lvalue {
		return false
	}
	switch v.data.Type.Id {
	case jntype.Func, jntype.Enum:
		return false
	default:
		return v.data.Tok.Id == tokens.Id
	}
}

func valIsStructIns(val value) bool {
	return !val.isType && typeIsStruct(val.data.Type)
}

func valIsTraitIns(val value) bool {
	return !val.isType && typeIsTrait(val.data.Type)
}

func isForeachIterExpr(val value) bool {
	switch {
	case typeIsSlice(val.data.Type),
		typeIsArray(val.data.Type),
		typeIsMap(val.data.Type):
		return true
	case !typeIsPure(val.data.Type):
		return false
	}
	code := val.data.Type.Id
	return code == jntype.Str
}

func isConstNumeric(v string) bool {
	if v == "" {
		return false
	}
	return v[0] == '-' || (v[0] >= '0' && v[0] <= '9')
}

func isConstExpression(v string) bool {
	return isConstNumeric(v) || isstr(v) || ischar(v) || isnil(v) || isbool(v)
}

func checkFloatBit(v models.Data, bit int) bool {
	if bit == 0 {
		return false
	}
	return jnbits.CheckBitFloat(v.Value, bit)
}

func validExprForConst(v value) bool {
	return v.constExpr
}

func okForShifting(v value) bool {
	if !typeIsPure(v.data.Type) ||
		!jntype.IsInteger(v.data.Type.Id) {
		return false
	}
	if !v.constExpr {
		return true
	}
	switch t := v.expr.(type) {
	case int64:
		return t >= 0
	case uint64:
		return true
	}
	return false
}