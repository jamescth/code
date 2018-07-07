package main

func romanToInt(s string) int {
	ret := 0

	i, l := 0, len(s)
	for i < l {
		b := s[i]
		switch b {
		case 'I':
			if i+1 < l && (s[i+1] == 'V' || s[i+1] == 'X') {
				ret--
			} else {
				ret++
			}
		case 'X':
			if i+1 < l && (s[i+1] == 'L' || s[i+1] == 'C') {
				ret -= 10
			} else {
				ret += 10
			}
		case 'C':
			if i+1 < l && (s[i+1] == 'D' || s[i+1] == 'M') {
				ret -= 100
			} else {
				ret += 100
			}
		case 'V', 'L', 'D', 'M':
			ret += stoi(b)
		}

		i++
	}
	return ret
}

// symbol to integer
func stoi(b byte) int {
	switch b {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}
	return 0
}
