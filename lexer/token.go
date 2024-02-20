package lexer

import "github.com/De-Rune/jane/package/io"

type Token struct {
	File   *io.FILE
	Line   int
	Column int
	Value  string
	Type   uint8
}

const (
	NA        uint8 = 0
	Type      uint8 = 1
	Name      uint8 = 2
	Brace     uint8 = 3
	Return    uint8 = 4
	SemiColon uint8 = 5
	Value     uint8 = 6
	Operator  uint8 = 7
	Comma     uint8 = 8
)
