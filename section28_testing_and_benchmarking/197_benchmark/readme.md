## Remember BET
- Benchmark
- Example
- Test

```
BenchmarkCat(b *testing.B)
ExampleCat()
TestCat(t *testing.T)
```

## Commands
```
godoc -http=:8080
go test
go test -bench .
go test -cover
go test -coverprofile=c.out
go tool cover -html=c.out
```


### Test coverage
```
We can use the “-cover” flag to run coverage analysis on our code.

go test -cover

We can use the flag and required file name “-coverprofile <some file name>” to write our coverage analysis to a file.
code:

go test -coverprofile=c.out

show in browser (visual for each line):
go tool cover -html=c.out
```
