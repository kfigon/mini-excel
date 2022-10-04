package main

type evaluator struct{}
func newEvaluator() *evaluator{
	return &evaluator{}
}

func(e *evaluator) eval(exp expressionCell) int {
	return -1
}