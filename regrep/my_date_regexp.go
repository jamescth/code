package main

import (
	"regexp"
	"testing"
)

func myDateParser(s string) []string {
	// re := regexp.MustCompile(`^(19|20)\d\d[ \.\-\/](1[012]|0[0-9])[ \.\-\/]([0-2][0-9]|3[01])::([01][0-9]|2[0-4]):[0-5][0-9]:[0-5][0-9]$`)
	re := regexp.MustCompile(`(19|20)\d\d[ -\.\/](1[012]|0[0-9])[ -\.\/]([0-2][0-9]|3[01])::([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9]`)
	return re.FindAllString(s, -1)
}

func TestTime(t *testing.T) {
	tests := []struct {
		n   int
		s   string
		ret []string
	}{
		{1, "2017-12-21::01:02:33", []string{"2017-12-21::01:02:33"}},
		{2, "2017-12-21::01:02:33 asdfasdf", []string{"2017-12-21::01:02:33"}},
		{3, "2017-12-21::01:02:33  2017-12-21::01:02:33", []string{"2017-12-21::01:02:33", "2017-12-21::01:02:33"}},
		{4, "2017 12 21::01:02:33", []string{"2017 12 21::01:02:33"}},
		{5, "2017/12/21::24:02:33", []string{}},
		{5, "2017/12/21::23:59:33", []string{"2017/12/21::23:59:33"}},
	}

	for _, tc := range tests {
		ret := myDateParser(tc.s)
		l := len(ret)
		if l != len(tc.ret) {
			t.Errorf("%d failed %s epxect %s got %s\n", tc.n, tc.s, tc.ret, ret)
		}
		for i := 0; i < l; i++ {
			if ret[i] != tc.ret[i] {
				t.Errorf("%d failed %s epxect %s got %s\n", tc.n, tc.s, tc.ret, ret)
			}
		}
	}
}
