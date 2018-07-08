package main

func canConstruct(ransomNote string, magazine string) bool {
	lr, lm := len(ransomNote), len(magazine)

	if lr == 0 && lm == 0 {
		return true
	}

	m := make(map[rune]int)
	for _, b := range ransomNote {
		if v, ok := m[b]; !ok {
			m[b] = 1
		} else {
			m[b] = v + 1
		}
	}

	for _, b := range magazine {
		if v, ok := m[b]; ok {
			if v == 1 {
				delete(m, b)
			} else {
				m[b] = v - 1
			}
		}
	}

	if len(m) > 0 {
		return false
	}
	return true
}
