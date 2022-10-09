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
}