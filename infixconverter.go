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

func (s *stack[T]) peek() (T, bool) {
	if s.empty() {
		var out T
		return out, false
	}
	out := s.tab[len(s.tab)-1]
	return out, true
}

func convertInfix(tokens []token) []token {
	out := []token{}
	stak := newStack[token]()

	for _, cur := range tokens {

		if cur.tokType.isOperand() || cur.tokType == equal {
			out = append(out, cur)
		} else if cur.tokType.isOperator() {

			if cur.tokType == openParent {
				stak.push(cur)
			} else if cur.tokType == closeParent {
				for {
					v, ok := stak.pop()
					if !ok || v.tokType == openParent {
						break
					}
					out = append(out, v)
				}
			} else {
				top, ok := stak.peek()
				if !ok || cur.tokType.predescence() > top.tokType.predescence(){
					stak.push(cur)
					continue
				}
				
				for {
					top, ok := stak.pop()
					if !ok {
						break
					}

					if top.tokType.predescence() >= cur.tokType.predescence() {
						out = append(out, top)
					} else {
						stak.push(top) // get back that popped one
						break
					}
				}
				stak.push(cur)
			}
		}
	}


	for {
		v, ok := stak.pop()
		if !ok {
			break
		}
		out = append(out, v)
	}

	return out
}