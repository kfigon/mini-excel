package main

import "fmt"

type reader interface {
	read(string) (cell, bool)
}

type evaluator struct {
	r reader
}

func newEvaluator(r reader) *evaluator {
	return &evaluator{r: r}
}

func (e *evaluator) eval(currentCoord string, exp expressionCell) (int, error) {
	tokens := parseExpression(exp)
	if len(tokens) < 2 {
		return -1, fmt.Errorf("invalid expression provided")
	}

	iter := &tokenIterator{tokens: tokens}
	iter.next() // =

	for iter.hasNext() {
		iter.next()		
	}
	
	return -1, nil
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
