
// Package lexer is fenerated by GoGLL. Do not edit.
package lexer

import(
    "io/ioutil"
    "strings"

    "github.com/goccmack/goutil/md"

    "github.com/goccmack/gogll/token"
)

type Lexer struct {
    I []rune
    Tokens []*token.Token
}

func NewFile(fname string) *Lexer {
    if strings.HasSuffix(fname, ".md") {
        src, err := md.GetSource(fname)
        if err != nil {
            panic(err)
        }
        return New([]rune(src))
    }
    buf, err := ioutil.ReadFile(fname)
    if err != nil {
        panic(err)
    }
    return New([]rune(string(buf)))
}

func New(input []rune) *Lexer {
    panic("implement")
}

// GetLineColumn returns the line and column of rune[i] in the input
func (l *Lexer) GetLineColumn(i int) (line, col int) {
	line, col = 1, 1
	for j := 0; j < i; j++ {
		switch l.I[j] {
		case '\n':
			line++
			col = 1
		case '\t':
			col += 4
		default:
			col++
		}
	}
	return
}