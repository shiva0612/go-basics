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
go build --race
go test ./... -race (to check if any race condition errors)

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

## at the top of the file mention the below line
```
// +build unit
// +build integration
go test -tags=unit -v
go test -tags=integration -v
```
