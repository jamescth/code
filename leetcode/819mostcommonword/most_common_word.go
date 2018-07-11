package main

import (
	"fmt"
	"strings"
	"unicode"
)

func mostCommonWord(paragraph string, banned []string) string {
	ban := make(map[string]struct{})
	wMap := make(map[string]int)

	fn := func(r rune) bool {
		return !unicode.IsLetter(r)
	}

	for _, b := range banned {
		b = strings.TrimLeftFunc(b, fn)
		b = strings.TrimRightFunc(b, fn)
		ban[b] = struct{}{}
	}

	strs := strings.Fields(paragraph)
	for _, s := range strs {
		if len(s) == 0 {
			continue
		}
		s = strings.TrimLeftFunc(s, fn)
		s = strings.TrimRightFunc(s, fn)
		s = strings.ToLower(s)

		if _, ok := ban[s]; ok {
			continue
		}

		if v, ok := wMap[s]; !ok {
			wMap[s] = 1
		} else {
			wMap[s] = v + 1
		}
	}
	fmt.Printf("wMap %v\n", wMap)

	var ret string
	var max int = 0
	for k, v := range wMap {
		fmt.Printf("v %d k %s\n", v, k)
		if v > max {
			max = v
			ret = k
		}
	}

	return ret
}
