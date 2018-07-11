package main

import (
	"fmt"
	"testing"
	"time"
)

var (
	timedivider int64 = 1000000000

	dalogTimeFmt     string = "2006-01-02T15:04:05.000+0000"
	daheadTimeFmt    string = "2006-01-02T15:04:05"
	esxlogTimeFmt    string = "2006-01-02T15:04:05.000Z"
	esxSyslogTimeFmt string = "2006-01-02T15:04:05Z"
	testlogTimeFmt   string = "2006-01-02T15:04:05.000000+0000"
)

func daNanosecond(t int64) string {
	return time.Unix(t/timedivider, t%timedivider).In(time.UTC).Format(dalogTimeFmt)
}

func daSecond(t int64) string {
	return time.Unix(t, 0).In(time.UTC).Format(dalogTimeFmt)
}
func ParseTime(s string) (int64, error) {
	// timestamp := str[len(search) : len(search)+len(dalogTimeFmt)]
	parseTime, err := time.Parse(dalogTimeFmt, s)
	if err != nil {
		return 0, err
	}
	return parseTime.UnixNano(), nil
}

func TestParseTime(t *testing.T) {
	tests := []struct {
		n   int
		x   string
		ret string
	}{
		{1, "2018-07-02T15:04:05.000+0000", "2018-07-02T15:04:05.000+0000"},
	}

	fmt.Println("\nTestParseTime")
	for _, tc := range tests {
		unixTime, err := ParseTime(tc.x)
		if err != nil {
			t.Errorf("%d failed %s err %v\n", tc.n, tc.x, err)
		}
		if ret := daNanosecond(unixTime); ret != tc.ret {
			t.Errorf("%d failed %s expect %s got %s\n", tc.n, tc.x, tc.ret, ret)
		}
	}
}
