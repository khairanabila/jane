package parser

import (
	"fmt"
	"strings"

	"github.com/De-Rune/jane/ast"
	"github.com/De-Rune/jane/lexer"
)

type Function struct {
	Token      lexer.Token
	Name       string
	ReturnType uint8
	Block      ast.BlockAST
}

func (f Function) String() string {
	var sb strings.Builder
	sb.WriteString(cxxTypeNameFromType(f.ReturnType))
	sb.WriteByte(' ')
	sb.WriteString(f.Name)
	sb.WriteString("()")
	sb.WriteString(" {")
	for _, s := range f.Block.Content {
		sb.WriteByte('\n')
		sb.WriteString("\t" + fmt.Sprint(s.Value))
		sb.WriteByte(';')
	}
	sb.WriteString("\n}")
	return sb.String()
}
