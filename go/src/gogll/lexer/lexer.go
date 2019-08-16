// Code generated by gocc; DO NOT EDIT.

package lexer

import (
	"io/ioutil"
	"unicode/utf8"

	"gogll/token"
)

const (
	NoState    = -1
	NumStates  = 79
	NumSymbols = 88
)

type Lexer struct {
	src    []byte
	pos    int
	line   int
	column int
}

func NewLexer(src []byte) *Lexer {
	lexer := &Lexer{
		src:    src,
		pos:    0,
		line:   1,
		column: 1,
	}
	return lexer
}

func NewLexerFile(fpath string) (*Lexer, error) {
	src, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	return NewLexer(src), nil
}

func (l *Lexer) Scan() (tok *token.Token) {
	tok = new(token.Token)
	if l.pos >= len(l.src) {
		tok.Type = token.EOF
		tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = l.pos, l.line, l.column
		return
	}
	start, startLine, startColumn, end := l.pos, l.line, l.column, 0
	tok.Type = token.INVALID
	state, rune1, size := 0, rune(-1), 0
	for state != -1 {
		if l.pos >= len(l.src) {
			rune1 = -1
		} else {
			rune1, size = utf8.DecodeRune(l.src[l.pos:])
			l.pos += size
		}

		nextState := -1
		if rune1 != -1 {
			nextState = TransTab[state](rune1)
		}
		state = nextState

		if state != -1 {

			switch rune1 {
			case '\n':
				l.line++
				l.column = 1
			case '\r':
				l.column = 1
			case '\t':
				l.column += 4
			default:
				l.column++
			}

			switch {
			case ActTab[state].Accept != -1:
				tok.Type = ActTab[state].Accept
				end = l.pos
			case ActTab[state].Ignore != "":
				start, startLine, startColumn = l.pos, l.line, l.column
				state = 0
				if start >= len(l.src) {
					tok.Type = token.EOF
				}

			}
		} else {
			if tok.Type == token.INVALID {
				end = l.pos
			}
		}
	}
	if end > start {
		l.pos = end
		tok.Lit = l.src[start:end]
	} else {
		tok.Lit = []byte{}
	}
	tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = start, startLine, startColumn

	return
}

func (l *Lexer) Reset() {
	l.pos = 0
}

/*
Lexer symbols:
0: '"'
1: '"'
2: '''
3: '''
4: 'p'
5: 'a'
6: 'c'
7: 'k'
8: 'a'
9: 'g'
10: 'e'
11: ':'
12: ';'
13: '*'
14: '|'
15: 'e'
16: 'm'
17: 'p'
18: 't'
19: 'y'
20: 'A'
21: 'l'
22: 't'
23: 'a'
24: 'n'
25: 'y'
26: 'a'
27: 'n'
28: 'y'
29: 'o'
30: 'f'
31: 'l'
32: 'e'
33: 't'
34: 't'
35: 'e'
36: 'r'
37: 'n'
38: 'u'
39: 'm'
40: 'b'
41: 'e'
42: 'r'
43: 'u'
44: 'p'
45: 'c'
46: 'a'
47: 's'
48: 'e'
49: 'l'
50: 'o'
51: 'w'
52: 'c'
53: 'a'
54: 's'
55: 'e'
56: 'n'
57: 'o'
58: 't'
59: 's'
60: 'p'
61: 'a'
62: 'c'
63: 'e'
64: '\'
65: '"'
66: 'n'
67: 't'
68: 'r'
69: '''
70: '\'
71: '_'
72: '/'
73: '*'
74: '*'
75: '*'
76: '/'
77: '/'
78: '/'
79: '\n'
80: ' '
81: '\t'
82: '\n'
83: '\r'
84: 'A'-'Z'
85: 'a'-'z'
86: '0'-'9'
87: .
*/
