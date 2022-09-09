# Review

Here is a review of the different commands useful with benchmarks, examples, and tests.
- godoc -http=:8080
- go test
- go test -bench .
	- don’t forget the “.” in the line above
- go test -cover
- go test -coverprofile c.out
- go tool cover -html=c.out