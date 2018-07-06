package main

import (
	"math"
	"strings"
)

func myAtoi(str string) int {
	var ret int

	str = strings.Trim(str, " ")
	if len(str) == 0 {
		return 0
	}

	sign, i := 1, 0
	if str[0] == '+' {
		sign, i = 1, 1
	}
	if str[0] == '-' {
		sign, i = -1, 1
	}

	for ; i < len(str); i++ {
		// fmt.Printf("%s %c %d %d\n", str, str[i], i, ret)
		if str[i] < '0' || str[i] > '9' || ret > math.MaxInt32 {
			break
		}
		ret = ret*10 + int(str[i]-'0')
	}

	// fmt.Printf("%d\n", ret)
	r, over := isOverflow(sign, ret)
	if over {
		return r
	}
	return sign * ret
}

func isOverflow(sign int, n int) (int, bool) {
	if sign > 0 && n > math.MaxInt32 {
		return math.MaxInt32, true
	}
	if sign < 0 && sign*n < math.MinInt32 {
		return math.MinInt32, true
	}
	return n, false
}

/*
func myAtoi(str string) int {
	var ret int

	fn := func(b byte) bool {
		if b < '0' || b > '9' {
			return false
		}
		return true
	}

	if str == "" {
		return 0
	}
	l := len(str)

	i := 0

	for i < l && str[i] == ' ' {
		i++
	}
	if i >= l {
		return 0
	}

	var tmp string
	if str[i] == '-' {
		tmp += string(str[i])
		if i == l-1 || !fn(str[i+1]) {
			return 0
		}
		i++
	}
	if str[i] == '+' {
		if i == l-1 || !fn(str[i+1]) {
			return 0
		}
        i++
	}

	if !fn(str[i]) {
		return 0
	}

	for ; i < l; i++ {
		if !fn(str[i]) {
			break
		}

		tmp += string(str[i])
	}

	ret, _ = strconv.Atoi(tmp)

	if ret > math.MaxInt32 {
		ret = math.MaxInt32
	}

	if ret < math.MinInt32 {
		ret = math.MinInt32
	}
	return ret
}
*/
