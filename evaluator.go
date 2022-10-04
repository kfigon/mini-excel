package main

type evaluator struct{}
func newEvaluator() *evaluator{
	return &evaluator{}
}

func(e *evaluator) eval(exp expressionCell) (int, error) {
	return -1, nil
}