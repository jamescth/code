package main

import (
	"fmt"
	"regexp"
	"testing"
)

func Example() {
	// Compile the expression once, usually at init time.
	// Use raw strings to avoid having to quote the backslashes.
	var validID = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)

	fmt.Println(validID.MatchString("adam[23]"))
	fmt.Println(validID.MatchString("eve[7]"))
	fmt.Println(validID.MatchString("Job[48]"))
	fmt.Println(validID.MatchString("snakey"))
	// Output:
	// true
	// true
	// false
	// false
}

func ExampleMatchString() {
	matched, err := regexp.MatchString("foo.*", "seafood")
	fmt.Println(matched, err)
	matched, err = regexp.MatchString("bar.*", "seafood")
	fmt.Println(matched, err)
	matched, err = regexp.MatchString("a(b", "seafood")
	fmt.Println(matched, err)
	// Output:
	// true <nil>
	// false <nil>
	// false error parsing regexp: missing closing ): `a(b`
}

func ExampleRegexp_FindString() {
	re := regexp.MustCompile("foo.?")
	fmt.Printf("%q\n", re.FindString("seafood fool"))
	fmt.Printf("%q\n", re.FindString("meat"))
	// Output:
	// "food"
	// ""
}

func ExampleRegexp_FindStringIndex() {
	re := regexp.MustCompile("ab?")
	fmt.Println(re.FindStringIndex("tablett"))
	fmt.Println(re.FindStringIndex("foo"))
	// Output:
	// [1 3]
	// []
}

func ExampleRegexp_FindStringSubmatch() {
	re := regexp.MustCompile("a(x*)b(y|z)c")
	fmt.Printf("%q\n", re.FindStringSubmatch("-axxxbyc-"))
	fmt.Printf("%q\n", re.FindStringSubmatch("-abzc-"))
	// Output:
	// ["axxxbyc" "xxx" "y"]
	// ["abzc" "" "z"]
}

func ExampleRegexp_FindAllString() {
	re := regexp.MustCompile("a.")
	fmt.Println(re.FindAllString("paranormal", -1))
	fmt.Println(re.FindAllString("paranormal", 2))
	fmt.Println(re.FindAllString("graal", -1))
	fmt.Println(re.FindAllString("none", -1))
	// Output:
	// [ar an al]
	// [ar an]
	// [aa]
	// []
}

func ExampleRegexp_FindAllStringSubmatch() {
	re := regexp.MustCompile("a(x*)b")
	fmt.Printf("%q\n", re.FindAllStringSubmatch("-ab-", -1))
	fmt.Printf("%q\n", re.FindAllStringSubmatch("-axxb-", -1))
	fmt.Printf("%q\n", re.FindAllStringSubmatch("-ab-axb-", -1))
	fmt.Printf("%q\n", re.FindAllStringSubmatch("-axxb-ab-", -1))
	// Output:
	// [["ab" ""]]
	// [["axxb" "xx"]]
	// [["ab" ""] ["axb" "x"]]
	// [["axxb" "xx"] ["ab" ""]]
}

func ExampleRegexp_FindAllStringSubmatchIndex() {
	re := regexp.MustCompile("a(x*)b")
	// Indices:
	//    01234567   012345678
	//    -ab-axb-   -axxb-ab-
	fmt.Println(re.FindAllStringSubmatchIndex("-ab-", -1))
	fmt.Println(re.FindAllStringSubmatchIndex("-axxb-", -1))
	fmt.Println(re.FindAllStringSubmatchIndex("-ab-axb-", -1))
	fmt.Println(re.FindAllStringSubmatchIndex("-axxb-ab-", -1))
	fmt.Println(re.FindAllStringSubmatchIndex("-foo-", -1))
	// Output:
	// [[1 3 2 2]]
	// [[1 5 2 4]]
	// [[1 3 2 2] [4 7 5 6]]
	// [[1 5 2 4] [6 8 7 7]]
	// []
}

func ExampleRegexp_MatchString() {
	re := regexp.MustCompile("(gopher){2}")
	fmt.Println(re.MatchString("gopher"))
	fmt.Println(re.MatchString("gophergopher"))
	fmt.Println(re.MatchString("gophergophergopher"))
	// Output:
	// false
	// true
	// true
}

func ExampleRegexp_ReplaceAllLiteralString() {
	re := regexp.MustCompile("a(x*)b")
	fmt.Println(re.ReplaceAllLiteralString("-ab-axxb-", "T"))
	fmt.Println(re.ReplaceAllLiteralString("-ab-axxb-", "$1"))
	fmt.Println(re.ReplaceAllLiteralString("-ab-axxb-", "${1}"))
	// Output:
	// -T-T-
	// -$1-$1-
	// -${1}-${1}-
}

func ExampleRegexp_ReplaceAllString() {
	re := regexp.MustCompile("a(x*)b")
	fmt.Println(re.ReplaceAllString("-ab-axxb-", "T"))
	fmt.Println(re.ReplaceAllString("-ab-axxb-", "$1"))
	fmt.Println(re.ReplaceAllString("-ab-axxb-", "$1W"))
	fmt.Println(re.ReplaceAllString("-ab-axxb-", "${1}W"))
	// Output:
	// -T-T-
	// --xx-
	// ---
	// -W-xxW-
}

func ExampleRegexp_SubexpNames() {
	re := regexp.MustCompile("(?P<first>[a-zA-Z]+) (?P<last>[a-zA-Z]+)")
	fmt.Println(re.MatchString("Alan Turing"))
	fmt.Printf("%q\n", re.SubexpNames())
	reversed := fmt.Sprintf("${%s} ${%s}", re.SubexpNames()[2], re.SubexpNames()[1])
	fmt.Println(reversed)
	fmt.Println(re.ReplaceAllString("Alan Turing", reversed))
	// Output:
	// true
	// ["" "first" "last"]
	// ${last} ${first}
	// Turing Alan
}

func ExampleRegexp_Split() {
	a := regexp.MustCompile("a")
	fmt.Println(a.Split("banana", -1))
	fmt.Println(a.Split("banana", 0))
	fmt.Println(a.Split("banana", 1))
	fmt.Println(a.Split("banana", 2))
	zp := regexp.MustCompile("z+")
	fmt.Println(zp.Split("pizza", -1))
	fmt.Println(zp.Split("pizza", 0))
	fmt.Println(zp.Split("pizza", 1))
	fmt.Println(zp.Split("pizza", 2))
	// Output:
	// [b n n ]
	// []
	// [banana]
	// [b nana]
	// [pi a]
	// []
	// [pizza]
	// [pi a]
}

func ExampleRegexp_Expand() {
	content := []byte(`
	# comment line
	option1: value1
	option2: value2

	# another comment line
	option3: value3
`)

	// Regex pattern captures "key: value" pair from the content.
	pattern := regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)

	// Template to convert "key: value" to "key=value" by
	// referencing the values captured by the regex pattern.
	template := []byte("$key=$value\n")

	result := []byte{}

	// For each match of the regex in the content.
	for _, submatches := range pattern.FindAllSubmatchIndex(content, -1) {
		// Apply the captured submatches to the template and append the output
		// to the result.
		result = pattern.Expand(result, template, content, submatches)
	}
	fmt.Println(string(result))
	// Output:
	// option1=value1
	// option2=value2
	// option3=value3
}

func ExampleRegexp_ExpandString() {
	content := `
	# comment line
	option1: value1
	option2: value2

	# another comment line
	option3: value3
`

	// Regex pattern captures "key: value" pair from the content.
	pattern := regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)

	// Template to convert "key: value" to "key=value" by
	// referencing the values captured by the regex pattern.
	template := "$key=$value\n"

	result := []byte{}

	// For each match of the regex in the content.
	for _, submatches := range pattern.FindAllStringSubmatchIndex(content, -1) {
		// Apply the captured submatches to the template and append the output
		// to the result.
		result = pattern.ExpandString(result, template, content, submatches)
	}
	fmt.Println(string(result))
	// Output:
	// option1=value1
	// option2=value2
	// option3=value3
}

func TestRegrep(t *testing.T) {
	tests := []struct {
		n  int
		fn func()
	}{
		{1, Example},
		{2, ExampleMatchString},
		{3, ExampleRegexp_FindString},
		//{4, ExampleRegexp_FindStringIndex},
		//{5, ExampleRegexp_FindStringSubmatch},
		{6, ExampleRegexp_FindAllString},
		//{7, ExampleRegexp_FindAllStringSubmatch},
		{8, ExampleRegexp_MatchString},
		//{9, ExampleRegexp_ReplaceAllLiteralString},
		//{10, ExampleRegexp_ReplaceAllString},
		//{11, ExampleRegexp_SubexpNames},
		{12, ExampleRegexp_Split},
		//{13, ExampleRegexp_Expand},
		//{14, ExampleRegexp_ExpandString},
	}
	t.Skip("Skip for now")
	for _, tc := range tests {
		fmt.Printf("******** test %d\n", tc.n)
		tc.fn()
	}
}
