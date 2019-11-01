package main

import (
	"fmt"
	"testing"
	"time"
)

func TestHelloWorld(t *testing.T) {
	fn := func(c chan<- string) {
		defer close(c)
		s := "hello"
		for i := 0; i < len(s); i++ {
			c <- string(s[i])
		}
	}

	c := make(chan string)
	go fn(c)
	for i := range c {
		fmt.Println(i)
	}
}

func TestTwoChannels(t *testing.T) {
	timer := time.NewTimer(3 * time.Second)

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		case <-timer.C:
			return
		}
	}

	// dict 

	dict := map[string]string

	commits := map[string]int {
		"rsc": 3711,
	}

	for key, value := range dict {
		:Q

	}

	var m map[string]int

	m = make(map[string]int)
	delete(m, "route")
	if i, ok := m["route"]

	// array
	var a [4]int
	b := [...]string{"penn","Teller"}

	// slice
	letters := []string{"a","b","c"}

	func make([]T, len, cap) []T

	var s []byte
	s = make([]byte, 5, 5)

	a := []string{"a", "b"}
	b := []string{"c", "d"}
	a = append(a, b...)




}



