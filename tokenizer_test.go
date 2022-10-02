package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenizer(t *testing.T) {
	testCases := []struct {
		desc	string
		input 	string
		expected []token
	}{
		{
			desc: "simple addition",
			input: "=1+2",
			expected: []token{
				{equal, "="},
				{number, "1"},
				{plus, "+"},
				{number, "2"},
			},
		},
		{
			desc: "more operators",
			input: "=1+2*3-5",
			expected: []token{
				{equal, "="},
				{number, "1"},
				{plus, "+"},
				{number, "2"},
				{multiply, "*"},
				{number, "3"},
				{minus, "-"},
				{number, "5"},
			},
		},
		{
			desc: "more operators with whitespaces",
			input: " = 1+ 2 * 3 -5 ",
			expected: []token{
				{equal, "="},
				{number, "1"},
				{plus, "+"},
				{number, "2"},
				{multiply, "*"},
				{number, "3"},
				{minus, "-"},
				{number, "5"},
			},
		},
		{
			desc: "expression with parenthesis",
			input: "=(1+2)*3-5",
			expected: []token{
				{equal, "="},
				{openParent, "("},
				{number, "1"},
				{plus, "+"},
				{number, "2"},
				{closeParent, ")"},
				{multiply, "*"},
				{number, "3"},
				{minus, "-"},
				{number, "5"},
			},
		},
		{
			desc: "coordinates",
			input: "=A2",
			expected: []token{
				{equal, "="},
				{coord, "A2"},
			},
		},
		{
			desc: "expressions and coordinates",
			input: "=(ABC13+2)*AD53-5",
			expected: []token{
				{equal, "="},
				{openParent, "("},
				{coord, "ABC13"},
				{plus, "+"},
				{number, "2"},
				{closeParent, ")"},
				{multiply, "*"},
				{coord, "AD53"},
				{minus, "-"},
				{number, "5"},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := parseExpression(expressionCell{tC.input})
			assert.Equal(t, tC.expected, got)
		})
	}
}