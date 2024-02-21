package parser

import (
	"strings"

	"github.com/De-Rune/jane/ast"
	"github.com/De-Rune/jane/lexer"
	"github.com/De-Rune/jane/package/jane"
)

const entryPointStandard = `
#pragma region JANE_ENTRY_POINT_STANDARD_CODES
  setlocale(0x0, "");
#pragma endregion
`

type function struct {
	Token      lexer.Token
	Name       string
	ReturnType ast.TypeAST
	Params     []ast.ParameterAST
	Attributes []ast.AttributeAST
	Block      ast.BlockAST
}

func (f function) String() string {
	f.readyCxx()
	var cxx string
	cxx += attributeToString(f.Attributes)
	cxx += jane.CxxTypeNameFromType(f.ReturnType.Type)
	cxx += " "
	cxx += f.Name
	cxx += "("
	cxx += paramsToCxx(f.Params)
	cxx += ") {"
	cxx += getFunctionStandardCode(f.Name)
	cxx += f.Block.String()
	cxx += "\n}"
	return cxx
}

func (f *function) readyCxx() {
	switch f.Name {
	case jane.EntryPoint:
		f.ReturnType.Type = jane.Int32
	}
}

func attributeToString(attributes []ast.AttributeAST) string {
	var cxx strings.Builder
	for _, attribute := range attributes {
		cxx.WriteString(attribute.String())
		cxx.WriteByte(' ')
	}
	return cxx.String()
}

func paramsToCxx(params []ast.ParameterAST) string {
	if len(params) == 0 {
		return ""
	}
	var cxx string
	any := false
	for _, p := range params {
		cxx += p.String()
		cxx += ","
		if !any {
			any = p.Type.Type == jane.Any
		}
	}
	cxx = cxx[:len(cxx)-1]
	if any {
		cxx = "template <typename any>\n" + cxx
	}
	return cxx
}

func getFunctionStandardCode(name string) string {
	switch name {
	case jane.EntryPoint:
		return entryPointStandard
	}
	return ""
}
