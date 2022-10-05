package main

import (
	"fmt"
	"strconv"
	"unicode"
)

type tokenType int

const (
	number tokenType = iota
	plus
	minus
	multiply
	equal
	openParent
	closeParent
	coord
)

func (t tokenType) String() string {
	return []string{
		"number",
		"plus",
		"minus",
		"multiply",
		"equal",
		"openParent",
		"closeParent",
		"coord"}[t]
}

type token struct {
	tokType tokenType
	val     string
}

func (t token) String() string {
	return fmt.Sprintf("<%v; %v>", t.tokType, t.val)
}

func parseExpression(ex expressionCell) []token {
	out := []token{}
	singleCharTokens := map[byte]tokenType {
		'+': plus,
		'-': minus,
		'*': multiply,
		'(': openParent,
		')': closeParent,
		'=': equal,
	}

	i := 0
	for i < len(ex.exp) {
		c := ex.exp[i]
		if t, ok := singleCharTokens[c]; ok {
			out = append(out, token{t, string(c)})
			i++
		} else if unicode.IsSpace(rune(c)){
			i++
		} else {
			str := ""
			for i < len(ex.exp) && (unicode.IsDigit(rune(ex.exp[i])) || unicode.IsLetter(rune(ex.exp[i]))) {
				str += string(ex.exp[i])
				i++
			}
			
			if _, err := strconv.Atoi(str); err == nil {
				out = append(out, token{number, str})
			} else if _, ok := parseCoords(str); ok {
				out = append(out, token{coord, str})
			}
		}
	}

	return out
}


type tokenIterator struct {
	tokens []token
	i int
}

func (t *tokenIterator) hasNext() bool {
	return t.i < len(t.tokens)
}

func (t *tokenIterator) next() {
	if t.hasNext() {
		t.i++
	}
}

func (t *tokenIterator) currentToken() (token, bool) {
	if t.hasNext() {
		return t.tokens[t.i], true
	}
	return token{}, false
}

func (t *tokenIterator) peek() (token, bool) {
	if t.i+1 < len(t.tokens) {
		return t.tokens[t.i+1], true
	}
	return token{}, false
}
