# Section 28 - Testing and Benchmarking

Begins at this link: https://www.udemy.com/course/learn-how-to-code/learn/lecture/11922414#overview

## Links/Documentations

["testing" package](https://pkg.go.dev/testing)

[func (*T) Error(args ...any)](https://pkg.go.dev/testing#T.Error)

# Remember to BET
- Benchmark
- Example
- Test

```
BenchmarkCat(b *testing.B)
ExampleCat()
TestCat(t *testing.T)
```

# Commands

```
godoc -http=:8080

go test
go test -bench .
go test -cover
go test -coverprofile c.out
go tool cover -html=c.out
```

# Summary

**PATH TO DIRECTORY MUST NOT CONTAIN WHITE SPACE CHARS OR APOSTROPHES (')**

Tests must  
- be in a file that ends with “_test.go”  
- put the file in the same package as the one being tested  
- be in a func with a signature “func TestXxx(*testing.T)” 

Run a test  
- go test  

Deal with test failure  
- use t.Error to signal failure  

Nice idiom
- expected
- got

cmd: 	go test -v  
- for verbose printout  
- will test all files with _test


# Sections

194. [Introduction](https://github.com/ryanclove/LearningGoProgramming/tree/master/Testing_Lecture_and_Exercises/Section_28_Testing_and_Benchmarking/Sec_28.194)
195. [Table tests](https://github.com/ryanclove/LearningGoProgramming/tree/master/Testing_Lecture_and_Exercises/Section_28_Testing_and_Benchmarking/Sec_28.195)
196. [Example tests](https://github.com/ryanclove/LearningGoProgramming/tree/master/Testing_Lecture_and_Exercises/Section_28_Testing_and_Benchmarking/Sec_28.196)

Examples show up in documentation.
- godoc -http :8080
- https://blog.golang.org/examples
- go test ./…
    - test everything where i'm at./ and everything down the directory tree

197. [Golint](https://github.com/ryanclove/LearningGoProgramming/tree/master/Testing_Lecture_and_Exercises/Section_28_Testing_and_Benchmarking/Sec_28.197/README.md)

prints out style mistakes and is concerned with coding style

go get -u golang.org/x/lint/golint

- gofmt
	- formats go code
- go vet
	- reports suspicious constructs
- golint
	- reports poor coding style

198. [Benchmark](https://github.com/ryanclove/LearningGoProgramming/tree/master/Testing_Lecture_and_Exercises/Section_28_Testing_and_Benchmarking/Sec_28.198)

Part of the testing package allows us to measure the speed of our code. This could also be
called “measuring the performance” of your code, or “benchmarking” your code - finding out how
fast the code runs.  
- go test -bench .

199. [Coverage](https://github.com/ryanclove/LearningGoProgramming/tree/master/Testing_Lecture_and_Exercises/Section_28_Testing_and_Benchmarking/Sec_28.199)

Coverage in programming is how much of our code is covered by tests. We can use the “-cover”
flag to run coverage analysis on our code. We can use the flag and required file name
“-coverprofile <some file name>” to write our coverage analysis to a file.
code:
- go test -cover
	- go test -coverprofile c.out
		- show in browser:
			- go tool cover -html c.out
		- learn more
			- go tool cover -h

200. [Benchmark examples](https://github.com/ryanclove/LearningGoProgramming/tree/master/Testing_Lecture_and_Exercises/Section_28_Testing_and_Benchmarking/Sec_28.200)

201. [Review](https://github.com/ryanclove/LearningGoProgramming/tree/master/Testing_Lecture_and_Exercises/Section_28_Testing_and_Benchmarking/Sec_28.201/README.md)


Here is a review of the different commands useful with benchmarks, examples, and tests.
- godoc -http=:8080
- go test
- go test -bench .
	- don’t forget the “.” in the line above
- go test -cover
- go test -coverprofile c.out
- go tool cover -html=c.out