package models

import (
	"strings"
	"sync/atomic"

	"github.com/DeRuneLabs/jane/package/jn"
)

type Block struct {
	Parent   *Block
	SubIndex int
	Tree     []Statement
	Gotos    *Gotos
	Labels   *Labels
	Func     *Func
}

func (b Block) String() string {
	AddIndent()
	defer func() { DoneIndent() }()
	return ParseBlock(b)
}

func ParseBlock(b Block) string {
	var cpp strings.Builder
	cpp.WriteByte('{')
	for _, s := range b.Tree {
		if s.Data == nil {
			continue
		}
		cpp.WriteByte('\n')
		cpp.WriteString(IndentString())
		cpp.WriteString(s.String())
	}
	cpp.WriteByte('\n')
	indent := strings.Repeat(jn.Set.Indent, int(Indent-1)*jn.Set.IndentCount)
	cpp.WriteString(indent)
	cpp.WriteByte('}')
	return cpp.String()
}

var Indent uint32 = 0

func IndentString() string {
	return strings.Repeat(jn.Set.Indent, int(Indent)*jn.Set.IndentCount)
}

func AddIndent() {
	atomic.AddUint32(&Indent, 1)
}

func DoneIndent() {
	atomic.SwapUint32(&Indent, Indent-1)
}
