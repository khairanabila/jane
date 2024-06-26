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

package ast

import (
	"github.com/DeRuneLabs/jane/lexer/tokens"
	"github.com/DeRuneLabs/jane/package/jnlog"
)

func Range(i *int, open, close string, toks Toks) Toks {
	if *i >= len(toks) {
		return nil
	}
	tok := toks[*i]
	if tok.Id != tokens.Brace || tok.Kind != open {
		return nil
	}
	*i++
	braceCount := 1
	start := *i
	for ; braceCount != 0 && *i < len(toks); *i++ {
		tok := toks[*i]
		if tok.Id != tokens.Brace {
			continue
		}
		switch tok.Kind {
		case open:
			braceCount++
		case close:
			braceCount--
		}
	}
	return toks[start : *i-1]
}

func RangeLast(toks Toks) (cutted, cut Toks) {
	if len(toks) == 0 {
		return toks, nil
	} else if toks[len(toks)-1].Id != tokens.Brace {
		return toks, nil
	}
	braceCount := 0
	for i := len(toks) - 1; i >= 0; i-- {
		tok := toks[i]
		if tok.Id == tokens.Brace {
			switch tok.Kind {
			case tokens.RBRACE, tokens.RBRACKET, tokens.RPARENTHESES:
				braceCount++
				continue
			default:
				braceCount--
			}
		}
		if braceCount == 0 {
			return toks[:i], toks[i:]
		}
	}
	return toks, nil
}

func Parts(toks Toks, id uint8, exprMust bool) ([]Toks, []jnlog.CompilerLog) {
	if len(toks) == 0 {
		return nil, nil
	}
	var parts []Toks
	var errs []jnlog.CompilerLog
	braceCount := 0
	last := 0
	for i, tok := range toks {
		if tok.Id == tokens.Brace {
			switch tok.Kind {
			case tokens.LBRACE, tokens.LBRACKET, tokens.LPARENTHESES:
				braceCount++
				continue
			default:
				braceCount--
			}
		}
		if braceCount > 0 {
			continue
		}
		if tok.Id == id {
			if exprMust && i-last <= 0 {
				errs = append(errs, compilerErr(tok, "missing_expr"))
			}
			parts = append(parts, toks[last:i])
			last = i + 1
		}
	}
	if last < len(toks) {
		parts = append(parts, toks[last:])
	} else if !exprMust {
		parts = append(parts, Toks{})
	}
	return parts, errs
}

func SplitColon(toks Toks, i *int) (rangeToks Toks, colon int) {
	rangeToks = nil
	colon = -1
	braceCount := 0
	start := *i
	for ; *i < len(toks); *i++ {
		tok := toks[*i]
		if tok.Id == tokens.Brace {
			switch tok.Kind {
			case tokens.LBRACE, tokens.LBRACKET, tokens.LPARENTHESES:
				braceCount++
				continue
			default:
				braceCount--
			}
		}
		if braceCount == 0 {
			if start+1 > *i {
				return
			}
			rangeToks = toks[start+1 : *i]
			break
		} else if braceCount != 1 {
			continue
		}
		if colon == -1 && tok.Id == tokens.Colon {
			colon = *i - start - 1
		}
	}
	return
}
