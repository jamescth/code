package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(detectCapitalUse("usa"))
}

func detectCapitalUse(word string) bool {
	l := len(word)
	if l == 0 {
		return true
	}

	// cut the head
	h := word[0]
	tails := word[1:]

	hU := unicode.IsUpper(rune(h))

	var allTailLow, allTailUp bool

	allTailUp = allUp(tails)
	allTailLow = allLow(tails)

	fmt.Printf("h %s tails %s hU %v allTailLow %v\n", string(h), tails, hU, allTailLow)
	if hU && allTailLow {
		return true
	}
	if !hU && allTailLow {
		return true
	}
	if hU && allTailUp {
		return true
	}
	return false
}

func allLow(s string) bool {
	for _, c := range s {
		if !unicode.IsLower(c) {
			return false
		}
	}
	return true
}

func allUp(s string) bool {
	for _, c := range s {
		if !unicode.IsUpper(c) {
			return false
		}
	}
	return true
}
