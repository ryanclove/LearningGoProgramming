# Section 28 - Testing and Benchmarking

Begins at this link: https://www.udemy.com/course/learn-how-to-code/learn/lecture/11922414#overview

## Links/Documentations

["testing" package](https://pkg.go.dev/testing)

[func (*T) Error(args ...any)](https://pkg.go.dev/testing#T.Error)
s
## Summary

**PATH TO DIRECTORY MUST NOT CONTAIN WHITE SPACE CHARS OR APOSTROPHES (')**

Tests must  
    - be in a file that ends with “_test.go”  
    - put the file in the same package as the one    being tested  
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

Examples show up in documentation.
- godoc -http :8080
- https://blog.golang.org/examples
- go test ./…
    - test everything where i'm at./ and everything down the directory tree
  


## Sections

194. [Introduction](https://github.com/ryanclove/LearningGoProgramming/tree/master/Testing_Lecture_and_Exercises/Section_28_Testing_and_Benchmarking/Sec_28.194)
195. [Table tests](https://github.com/ryanclove/LearningGoProgramming/tree/master/Testing_Lecture_and_Exercises/Section_28_Testing_and_Benchmarking/Sec_28.195)
196. [Example tests](https://github.com/ryanclove/LearningGoProgramming/tree/master/Testing_Lecture_and_Exercises/Section_28_Testing_and_Benchmarking/Sec_28.196)
197. [Golint](https://github.com/ryanclove/LearningGoProgramming/tree/master/Testing_Lecture_and_Exercises/Section_28_Testing_and_Benchmarking/Sec_28.197)
198. [Benchmark](https://github.com/ryanclove/LearningGoProgramming/tree/master/Testing_Lecture_and_Exercises/Section_28_Testing_and_Benchmarking/Sec_28.198)
199. [Coverage](https://github.com/ryanclove/LearningGoProgramming/tree/master/Testing_Lecture_and_Exercises/Section_28_Testing_and_Benchmarking/Sec_28.199)
200. [Benchmark examples](https://github.com/ryanclove/LearningGoProgramming/tree/master/Testing_Lecture_and_Exercises/Section_28_Testing_and_Benchmarking/Sec_28.200)
201. [Review](https://github.com/ryanclove/LearningGoProgramming/tree/master/Testing_Lecture_and_Exercises/Section_28_Testing_and_Benchmarking/Sec_28.201)
