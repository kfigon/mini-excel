package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := newStack[int]()
	t.Run("empty", func(t *testing.T) {
		_, ok := s.pop()
		assert.False(t, ok)
	})

	t.Run("push 1, pop 2", func(t *testing.T) {
		s.push(3)
		s.push(81)
		
		v, ok := s.pop()
		assert.True(t, ok)
		assert.Equal(t, 81, v)

		v, ok = s.pop()
		assert.True(t, ok)
		assert.Equal(t, 3, v)

		v, ok = s.pop()
		assert.False(t, ok)

		s.push(123)
		v, ok = s.pop()
		assert.True(t, ok)
		assert.Equal(t, 123, v)
	})

	t.Run("peek", func(t *testing.T) {
		s.push(1)
		s.push(2)

		v, ok := s.peek()
		assert.True(t, ok)
		assert.Equal(t, 2, v)

		v, _ = s.pop()
		assert.Equal(t,2,v)
		v, _ = s.pop()
		assert.Equal(t,1,v)
		_, ok = s.pop()
		assert.False(t, ok)
	})
}

func TestInfixConverter(t *testing.T) {
	testCases := []struct {
		input string
		exp string
	}{
		{"=1+2", "= 1 2 +"},
		{"=a1+b1*c1", "= a1 b1 c1 * +"},
		{"=4+4*2*(1-5)", "= 4 4 2 * 1 5 - * +"},
		{"=1 + (2*3-1)-2", "= 1 2 3 * 1 - + 2 -"},
	}
	for _, tC := range testCases {
		t.Run(tC.input, func(t *testing.T) {
			postfix := convertInfix(parseExpression(expressionCell{tC.input}))

			got := ""
			for i,v := range postfix {
				got += v.val
				if i != len(postfix)-1 {
					got += " "
				}
			}

			assert.Equal(t, tC.exp, got)
		})
	}
}