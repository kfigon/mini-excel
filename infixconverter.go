package main


type stack[T any] struct {
	tab []T
}

func newStack[T any]() *stack[T] {
	return &stack[T]{
		tab: []T{},
	}
}

func (s *stack[T]) pop() (T, bool) {
	if s.empty() {
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

func (s *stack[T]) empty() bool {
	return len(s.tab) == 0
}

func convertInfix(tokens []token) []token {
	return tokens
}