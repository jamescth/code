package main

import (
	"bytes"
	"fmt"
	"net"
	"regexp"
	"strings"
	"testing"
)

func TestIPV4re(t *testing.T) {
	fmt.Println("\nTestIPV4re")

	str1 := `Proxy Port Last Check Proxy Speed Proxy Country Anonymity 118.99.81.204
    118.99.81.204 8080 34 sec Indonesia - Tangerang Transparent 2.184.31.2 8080 58 sec 
    Iran Transparent 93.126.11.189 8080 1 min Iran - Esfahan Transparent 202.118.236.130
    7777 1 min China - Harbin Transparent 62.201.207.9 8080 1 min Iraq Transparent
	255.255.255.255 0.0.0.0 256.1.1.1 266.266.266.266
	`

	re := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)

	fmt.Printf("Pattern: %v\n", re.String()) // print pattern
	fmt.Println(re.MatchString(str1))        // true

	fmt.Println("FindAllString")
	submatchall := re.FindAllString(str1, -1)
	for _, element := range submatchall {
		fmt.Println(element)
	}
}

var (
	ip1 = net.ParseIP("216.14.49.184")
	ip2 = net.ParseIP("216.14.49.191")
)

func checkNetIP(ip string) bool {
	trial := net.ParseIP(ip)
	if trial.To4() == nil {
		// fmt.Printf("%v is not an IPv4 address\n", trial)
		return false
	}
	if bytes.Compare(trial, ip1) >= 0 && bytes.Compare(trial, ip2) <= 0 {
		// fmt.Printf("%v is between %v and %v\n", trial, ip1, ip2)
		return true
	}
	// fmt.Printf("%v is NOT between %v and %v\n", trial, ip1, ip2)
	return false
}

func TestCheckIPRange(t *testing.T) {
	tests := []struct {
		n   int
		x   string
		ret bool
	}{
		{1, "1.2.3.4", false},
		{2, "216.14.49.185", true},
		{3, "1::16", false},
	}

	fmt.Println("\nTestCheckIPRange")

	for _, tc := range tests {
		if ret := checkNetIP(tc.x); ret != tc.ret {
			t.Errorf("%d failed %s expect %v got %v\n", tc.n, tc.x, tc.ret, ret)
		}
	}
}

func checkStringIP(s string) bool {
	// testIP := "216.14.49.185"
	IP1 := "216.14.49.184"
	IP2 := "216.14.49.191"

	// fmt.Println("\nTestStringIP")

	if s <= IP2 && s >= IP1 {
		return true
	} else {
		return false
	}
}
func TestCheckStringIPRange(t *testing.T) {
	tests := []struct {
		n   int
		x   string
		ret bool
	}{
		{1, "1.2.3.4", false},
		{2, "216.14.49.185", true},
		{2, "216.14.49.1", false},
		{2, "216.1.49.185", false},
		{3, "1::16", false},
	}

	fmt.Println("\nTestCheckStringIPRange")

	for _, tc := range tests {
		if ret := checkStringIP(tc.x); ret != tc.ret {
			t.Errorf("%d failed %s expect %v got %v\n", tc.n, tc.x, tc.ret, ret)
		}
	}
}

func BenchmarkCheckStringIP(b *testing.B) {
	for n := 0; n < b.N; n++ {
		checkStringIP("216.14.49.185")
	}
}

func BenchmarkCheckNetIP(b *testing.B) {
	for n := 0; n < b.N; n++ {
		checkNetIP("216.14.49.185")
	}
}

func TestParseIPV6(t *testing.T) {
	tests := []struct {
		n   int
		x   string
		ret string
	}{
		{1, "192.168.0.0", "192.168.0.0"},
		{2, "2001:db8::1:0:0:1", "2001:db8:0:0:1:0:0:1"},
		{3, "::1", "0:0:0:0:0:0:0:1"},
		{4, "fe80::", "fe80:0:0:0:0:0:0:0"},
		{5, "fe80::1", "fe80:0:0:0:0:0:0:1"},
		{6, "fe80:0:0:0:204:61ff:fe9d:f156", "fe80:0:0:0:204:61ff:fe9d:f156"},
		{7, "fe80:0:0:204:61ff:fe9d:f156", "0:fe80:0:0:204:61ff:fe9d:f156"},
	}

	fmt.Println("\nTestParseIPV6")

	fn := func(c rune) bool {
		if c == '.' || c == ':' {
			return true
		}
		return false
	}
	_ = fn

	for _, tc := range tests {
		if ret := parseIPV6(tc.x); ret != tc.ret {
			t.Errorf("%d failed %s expect %s got %s\n", tc.n, tc.x, tc.ret, ret)
		}
	}
}

func parseIPV6(addr string) string {
	str := strings.Split(addr, ":")

	// need to calculate how many empty []
	l := len(str)
	needed := 0

	if l < 8 {
		needed = 8 - l + 1
	}
	// fmt.Printf("%v len %d\n", str, len(str))

	var ret string
	var empty int
	for i, r := range str {
		//fmt.Printf("i %d r %v\n", i, r)
		if i == 0 && len(r) == 0 {
			// 1st []int is empty.
			// we are in ::1 case
			needed = 7
			for needed > 0 {
				ret += "0:"
				needed--
			}
			empty++
		} else if i == (l-1) && len(str[i]) == 0 {
			// this is fe00::
			// len is only 3
			ret += "0:"
			empty++
		} else if len(r) == 0 {
			// for non 0 empty []int
			for needed > 0 {
				ret += "0:"
				needed--
			}
			empty++
		} else {
			ret += r + ":"
		}
	}
	ret = ret[:len(ret)-1]
	if empty == 0 && l == 7 {
		ret = "0:" + ret
	}
	return ret
}
