go test -count=1 -run=TestDpFib -v
#=== RUN   TestDpFib
#--- PASS: TestDpFib (0.00s)
#	fib_dp_test.go:31: 102334155
#=== RUN   TestDpFibSlow
#--- PASS: TestDpFibSlow (1.58s)
#	fib_dp_test.go:35: 102334155

$  go test -test.bench BenchmarkFib -run=^a
goos: linux
goarch: amd64
pkg: github.com/jamescth/code/leetcode/fib_dp
BenchmarkFibFast-2   	300000000	         5.55 ns/op
BenchmarkFibSlow-2   	       1	1682826167 ns/op
PASS
ok  	github.com/jamescth/code/leetcode/fib_dp	3.889s

go test -count=1 -run=TestFac

