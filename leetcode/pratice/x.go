package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(solve(1800))
}

func solve(year int32) string {
	var leap bool
	n := year

	if n%400 == 0 {
		fmt.Println("400")
		leap = true
	} else if n%100 == 0 {
		fmt.Println("100")
		leap = false
	} else if n%4 == 0 {
		fmt.Println("4")
		leap = true
	}
	fmt.Printf("%v", leap)
	if leap {
		return "12.09." + strconv.Itoa(int(year))
	}
	return "13.09." + strconv.Itoa(int(year))
}

//IP4
// IP6
