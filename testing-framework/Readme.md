##  general commands for testing:
```
go test ./...
from dir where this cmd is run
it will recursively look for all files (*_test.go) and execute Test<funcName>

go test -v -run TestAdd #runs single test in a file
go test -v calc_test.go #runs all test in a file
        
v = verbose

go test ./... --cover #cover shows the test coverage percentage
go test ./... -coverprofile=coverage.output #generates coverage report to a file 
go tool cover  -html=coverage.output #to see the report in a html format
```

##  checking for race conditions
```
go build --race main.go
./main.go (will give warning if race condition)

for test files
go test ./... -race (to check if any race condition errors)

```

## benchmarking
```
#run flag to only run those tests funcs whose name start with Benchmark
go test ./... -run=Benchmark -bench=. -benchtime=5s -count=5 -benchmem
go test ./... -run=Benchmark -bench=. -benchtime=100x -count=5 -benchmem

100x = run benchmark 100 times rather than b.N
5s = run benchmark 5s 


go test -bench=. -count 2 -benchtime=2s -cpu=1 -benchmem -memprofile memprofile.out -cpuprofile profile.out
go tool pprof memprofile.out
top
list function_name
```


##  testing main - not IMP
```
func TestMain(m *testing.M) {
	fmt.Println("Hello World")
	ret := m.Run()
	fmt.Println("Tests have executed")
	os.Exit(ret)
}
go test file.go -v
```

## using tags to run specific tests all at once

```
add these lines above package declaration
// +build unit
// +build integration

use the bellow command to run specific tests mentioned in tags flag
go test ./... --tags=unit -v
go test ./... --tags=integration -v
```
