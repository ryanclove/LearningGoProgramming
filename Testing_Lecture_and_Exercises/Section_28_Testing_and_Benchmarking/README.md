# Section 28 - Testing and Benchmarking

Begins at this link: https://www.udemy.com/course/learn-how-to-code/learn/lecture/11922414#overview

## Links/Documentations

["testing" package](https://pkg.go.dev/testing)

[func (*T) Error(args ...any)](https://pkg.go.dev/testing#T.Error)

## Summary

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


## Sections

194. [Introduction](https://github.com/ryanclove/LearningGoProgramming/blob/master/Lecture%20Sections/Section%2028%20Testing%20and%20Benchmarking/Sec%2028.194/main.go)
195. [Table tests](https://github.com/ryanclove/LearningGoProgramming/blob/master/Lecture%20Sections/Section%2028%20Testing%20and%20Benchmarking/Sec%2028.195/main.go)
196. [Example tests](https://github.com/ryanclove/LearningGoProgramming/blob/master/Lecture%20Sections/Section%2028%20Testing%20and%20Benchmarking/Sec%2028.196/main.go)
197. [Golint](https://github.com/ryanclove/LearningGoProgramming/blob/master/Lecture%20Sections/Section%2028%20Testing%20and%20Benchmarking/Sec%2028.197/main.go)
198. [Benchmark](https://github.com/ryanclove/LearningGoProgramming/blob/master/Lecture%20Sections/Section%2028%20Testing%20and%20Benchmarking/Sec%2028.198/main.go)
199. [Coverage](https://github.com/ryanclove/LearningGoProgramming/blob/master/Lecture%20Sections/Section%2028%20Testing%20and%20Benchmarking/Sec%2028.199/main.go)
200. [Benchmark examples](https://github.com/ryanclove/LearningGoProgramming/blob/master/Lecture%20Sections/Section%2028%20Testing%20and%20Benchmarking/Sec%2028.200/main.go)
201. [Review](https://github.com/ryanclove/LearningGoProgramming/blob/master/Lecture%20Sections/Section%2028%20Testing%20and%20Benchmarking/Sec%2028.201/main.go)
