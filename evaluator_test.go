package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvaluateExpressionWithoutCoordinates(t *testing.T) {
	testCases := []struct {
		input 	 string
		expected int
	}{
		{"=1+2", 3},
		{"=32+13", 45},
		{"=32-13", 19},
		{"=33*2", 66},
		{"=33*2 + 1", 67},
		{"=1+33*2", 67},
		{"=(1+33)*2", 68},
		{"=(1+33)+(5*3)*2", (1+33)+(5*3)*2},
	}
	for _, tC := range testCases {
		t.Run(tC.input, func(t *testing.T) {
			got, err := newEvaluator().eval(expressionCell{tC.input})
			assert.NoError(t, err)
			assert.Equal(t, tC.expected, got)
		})
	}
}

func TestEvaluateWithCoordinates(t *testing.T) {
	t.Fatal("todo")
}

func TestEvaluateExpressionsWithCyclicDependency(t *testing.T) {
	t.Fatal("todo")
}