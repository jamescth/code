package main

import "testing"

func TestIntToRoman(t *testing.T) {
	s := "hello" + "world" + string('i')
	t.Log(s)
	t.Log(intToRoman(3900))
}
