package main

var m = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	l := len(s)

	ret := 0
	for i := 0; i < l; i++ {
		c := s[i]

		switch c {
		case 'I':
			if i+1 < l {
				if s[i+1] == 'V' || s[i+1] == 'X' {
					ret--
					continue
				}
			}
			ret++
		case 'V':
			ret += 5
		case 'X':
			if i+1 < l {
				if s[i+1] == 'L' || s[i+1] == 'C' {
					ret -= 10
					continue
				}
			}
			ret += 10
		case 'L':
			ret += 50
		case 'C':
			if i+1 < l {
				if s[i+1] == 'D' || s[i+1] == 'M' {
					ret -= 100
					continue
				}
			}
			ret += 100
		case 'D':
			ret += 500
		case 'M':
			ret += 1000
		}
	}

	return ret
}
