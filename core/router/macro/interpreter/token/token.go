// Copyright 2017 Gerasimos Maropoulos. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package token

type TokenType int

type Token struct {
	Type    TokenType
	Literal string
	Start   int // including the first char
	End     int // including the last char
}

// /about/{fullname:alphabetical}
// /profile/{anySpecialName:string}
// {id:int range(1,5) else 404}
// /admin/{id:int eq(1) else 402}
// /file/{filepath:file else 405}
const (
	EOF = iota // 0
	ILLEGAL

	// Identifiers + literals
	LBRACE // {
	RBRACE // }
	//	PARAM_IDENTIFIER // id
	COLON  // :
	LPAREN // (
	RPAREN // )
	//	PARAM_FUNC_ARG   // 1
	COMMA
	IDENT // string or keyword
	// Keywords
	keywords_start
	ELSE // else
	keywords_end
	INT // 42
)

const eof rune = 0

var keywords = map[string]TokenType{
	"else": ELSE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
