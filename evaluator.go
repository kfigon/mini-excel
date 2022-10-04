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
	
	return -1, nil
}