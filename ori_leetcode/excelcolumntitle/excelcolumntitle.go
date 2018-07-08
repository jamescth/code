package main

func convertToTitle(n int) string {
	ret := ""

	for n > 0 {
		n--
		ret = string(byte(n%26)+'A') + ret
		n /= 26
	}
	return ret
}
