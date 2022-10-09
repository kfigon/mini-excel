package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockReader func(string) (cell, bool)
func (m mockReader) read(s string) (cell, bool) {
	return m(s)
}
func emptyMockReader() mockReader {
	return func(s string) (cell, bool) {
		return nil, false
	}
}

func TestEvaluateExpressionWithoutCoordinates(t *testing.T) {
	testCases := []struct {
		input 	 string
		expected int
	}{
		{"=2", 2},
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
			got, err := newEvaluator(emptyMockReader()).eval("", expressionCell{tC.input})
			assert.NoError(t, err)
			assert.Equal(t, tC.expected, got)
		})
	}
}

func TestEvaluateInvalidExpressions(t *testing.T) {
	testCases := []struct {
		input 	 string
	}{
		{"="},
		{"=+"},
		{"+"},
		{"=+3"},
		{"=3+"},
		{`=3/5`},
	}
	for _, tC := range testCases {
		t.Run(tC.input, func(t *testing.T) {
			_, err := newEvaluator(emptyMockReader()).eval("", expressionCell{tC.input})
			assert.Error(t, err)
		})
	}
}

func TestNonExisingCoordinates(t *testing.T) {
	testCases := []struct {
		input 	 string
	}{
		{"=A1"},
		{"=1 + A1"},
	}
	for _, tC := range testCases {
		t.Run(tC.input, func(t *testing.T) {
			_, err := newEvaluator(emptyMockReader()).eval("", expressionCell{tC.input})
			assert.Error(t, err)
		})
	}
}

func TestEvaluateWithCoordinates(t *testing.T) {
	testCases := []struct {
		input 	 string
		exp int
	}{
		{"=A1", 3},
		{"=1 + A1", 4},
		{"=B1 + A1*A2", 36},
		{"=(B1 + A1)*A2", 90},
		{"=C1", 45},
	}

	var mock mockReader = func(s string) (cell, bool) {
		switch s{
		case "A1": return numberCell(3), true
		case "B1": return numberCell(6), true
		case "A2": return numberCell(10), true
		case "C1": return expressionCell{exp: "=32+13"}, true
		}
		return nil, false
	}

	for _, tC := range testCases {
		t.Run(tC.input, func(t *testing.T) {
			_, err := newEvaluator(mock).eval("", expressionCell{tC.input})
			assert.Error(t, err)
		})
	}
}

func TestEvaluateExpressionsWithCyclicDependency(t *testing.T) {
	t.Fatal("todo")
}