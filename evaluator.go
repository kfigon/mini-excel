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

// shunting yard algorithm
// recursive descent parser
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
