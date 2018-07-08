package longestpalindromic

/*
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example:

Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.

Example:

Input: "cbbd"
Output: "bb"
*/

func longestPalindrome(s string) string {
	l := len(s)
	if l == 0 || l == 1 {
		return s
	}

	var L, R int // left and right pointer
	idx := 0     // walk idx

	var max string
	for idx < l {

		if (idx+1) < l && s[idx] == s[idx+1] {
			tempidx := idx
			for (tempidx+1) < l && s[tempidx] == s[tempidx+1] {
				tempidx++
			}
			//fmt.Printf("s %s dup %s %d %d\n", s, s[idx:tempidx+1], idx, tempidx)
			if len(max) < len(s[idx:tempidx+1]) {
				max = s[idx : tempidx+1]
			}
			idx = tempidx
		}

		size := l - 1 - idx
		if size > idx {
			size = idx
		}
		R = idx + size + 1
		if R > l {
			R = l
		}
		L = idx - size
		if L < 0 {
			L = 0
		}
		//fmt.Printf("%s tmp %s max %s size %d %d %d %d\n", s, s[L:R], max, size, L, idx, R)
		if isPalindromic(s[L:R]) {
			tmp := s[L:R]
			//fmt.Printf("s %s tmp %s max %s %d %d %d\n", s, tmp, max, L, idx, R)
			if len(max) < len(tmp) {
				max = tmp
			}
		}
		idx++
	}

	return max
}

func isPalindromic(str string) bool {
	l := len(str)

	i, j := 0, l-1
	for i < j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}

	return true
}

/*

func longestPalindrome(s string) string {
	if len(s) < 2 { // 肯定是回文，直接返回
		return s
	}

	// 最长回文的首字符索引，和最长回文的长度
	begin, maxLen := 0, 1

	// 在 for 循环中
	// b 代表回文的**首**字符索引号，
	// e 代表回文的**尾**字符索引号，
	// i 代表回文`正中间段`首字符的索引号
	// 在每一次for循环中
	// 先从i开始，利用`正中间段`所有字符相同的特性，让b，e分别指向`正中间段`的首尾
	// 再从`正中间段`向两边扩张，让b，e分别指向此`正中间段`为中心的最长回文的首尾
	for i := 0; i < len(s); { // 以s[i]为`正中间段`首字符开始寻找最长回文。
		if len(s)-i <= maxLen/2 {
			// 因为i是回文`正中间段`首字符的索引号
			// 假设此时能找到的最长回文的长度为l, 则，l <= (len(s)-i)*2 - 1
			// 如果，len(s)-i <= maxLen/2 成立
			// 则，l <= maxLen - 1
			// 则，l < maxLen
			// 所以，无需再找下去了。
			fmt.Println("i", i)
			break
		}

		b, e := i, i
		for e < len(s)-1 && s[e+1] == s[e] {
			e++
			// 循环结束后，s[b:e+1]是一串相同的字符串
		}
		fmt.Println("1 s[b:e+1]", s[b:e+1], b, e)

		// 下一个回文的`正中间段`的首字符只会是s[e+1]
		// 为下一次循环做准备
		i = e + 1

		for e < len(s)-1 && b > 0 && s[e+1] == s[b-1] {
			e++
			b--
			// 循环结束后，s[b:e+1]是这次能找到的最长回文。
		}
		fmt.Println("2 s[b:e+1]", s[b:e+1], b, e)

		newLen := e + 1 - b
		// 创新记录的话，就更新记录
		if newLen > maxLen {
			begin = b
			maxLen = newLen
		}
	}

	return s[begin : begin+maxLen]
}
*/
