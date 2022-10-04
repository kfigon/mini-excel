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
			got := newEvaluator().eval(expressionCell{tC.input})
			assert.Equal(t, tC.expected, got)
		})
	}
}