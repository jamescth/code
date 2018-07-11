package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(toGoatLatin("I speak Goat Latin"))
}

func toGoatLatin(s string) string {
	strs := strings.Fields(s)

	var ret string
	for i, str := range strs {
		//fmt.Printf("i %d str %s\n", i, str)

		if isVowel(string(str[0])) {
			ret += str + "ma"
		} else {
			ret += str[1:] + string(str[0]) + "ma"
		}

		for j := 0; j < i+1; j++ {
			ret += "a"
		}
		ret += " "
	}

	return ret[:len(ret)-1]
}

func isVowel(b string) bool {
	switch b {
	case string('a'), string('A'), string('e'), string('E'), string('i'), string('I'):
		return true
	case string('o'), string('O'), string('u'), string('U'):
		return true
	}
	return false
}
