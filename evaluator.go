package main

type reader interface{
	read(string, int) (cell, bool)
}

type evaluator struct{
	r reader
}

func newEvaluator(r reader) *evaluator{
	return &evaluator{r: r}
}

func(e *evaluator) eval(currentCoord coordinate, exp expressionCell) (int, error) {
	return -1, nil
}