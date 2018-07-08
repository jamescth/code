package main

type stack struct {
	s []string
}

func (s *stack) Push(b string) {
	s.s = append(s.s, b)
}

func (s *stack) Pop() string {
	l := len(s.s)
	if l == 0 {
		return ""
	}

	ret := s.s[l-1]
	s.s = s.s[:l-1]
	return ret
}

func isValid(s string) bool {
	l := len(s)
	if l == 0 {
		return true
	}
	if l < 2 {
		return false
	}

	// create a stack FILO
	i := 0
	mys := &stack{}

	for i < l {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			mys.Push(string(s[i]))
			i++
			continue
		}

		// pop
		r := string(mys.Pop())
		// fmt.Printf("%d s[i] %s r %s\n", i, string(s[i]), r)
		switch string(s[i]) {
		case ")":
			if r != "(" {
				return false
			}
		case "}":
			if r != "{" {
				return false
			}
		case "]":
			if r != "[" {
				return false
			}
		}
		i++
	}

	if len(mys.s) != 0 {
		return false
	}
	return true
}
