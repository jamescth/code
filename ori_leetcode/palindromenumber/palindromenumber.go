package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	// store digits
	var digits []int
	for {
		d := x % 10

		// insert in reverse direction, doesn't matter
		// since Palindrome is the 1st/last pair anyway
		digits = append(digits, d)

		x = x / 10
		if x == 0 {
			break
		}
		// how to end
	}

	i, j := 0, len(digits)-1
	for i <= j {
		if digits[i] != digits[j] {
			return false
		}
		i++
		j--
	}
	return true
}
