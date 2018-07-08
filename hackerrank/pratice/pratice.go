package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	nd := strings.Split(readLine(reader), " ")
	ntmp, _ := strconv.ParseInt(nd[0], 10, 64)

	fmt.Printf("%v\n", ntmp)
	second := strings.Split(readLine(reader), " ")
	fmt.Printf("%v %v\n", second[0], second[1])
}

func readLine(r *bufio.Reader) string {
	str, _, err := r.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}
