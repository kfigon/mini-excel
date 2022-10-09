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


type stack[T any] struct {
	tab []T
}

func newStack[T any]() *stack[T] {
	return &stack[T]{
		tab: []T{},
	}
}

func (s *stack[T]) pop() (T, bool) {
	if len(s.tab) == 0 {
		var out T
		return out, false
	}
	out := s.tab[len(s.tab)-1]
	s.tab = s.tab[0:len(s.tab)-1]
	return out, true
} 

func (s *stack[T]) push(v T) {
	s.tab = append(s.tab, v)
} 