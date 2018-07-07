package main

func isValid(s string) bool {
	l := len(s)

	if l == 0 {
		return true
	}
	if l == 1 {
		return false
	}
	stk := &stack{}

	for i := 0; i < l; i++ {
		b := s[i]

		switch b {
		case '(', '{', '[':
			stk.Push(b)
		case ')':
			r := stk.Pop()
			if r != '(' {
				return false
			}
		case '}':
			r := stk.Pop()
			if r != '{' {
				return false
			}
		case ']':
			r := stk.Pop()
			if r != '[' {
				return false
			}
		}
	}

	return true
}

type stack struct {
	s string
}

func (s *stack) Push(b byte) {
	s.s = s.s + string(b)
}

func (s *stack) Pop() byte {
	b := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return b
}
