package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	t.Log(lengthOfLongestSubstring("abcabcbb"))
	t.Log(lengthOfLongestSubstring("c"))
	t.Log(lengthOfLongestSubstring("cc"))
	t.Log(lengthOfLongestSubstring("ab"))
	t.Log(lengthOfLongestSubstring("dvdf"))

}
