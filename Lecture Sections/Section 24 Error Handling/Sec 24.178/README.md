# Understanding Errors in Go

We believe that coupling exceptions to a control structure, as in the try-catch-finally idiom, results in convoluted code. It also tends to encourage programmers to label too many ordinary errors, such as failing to open a file, as exceptional.

["errors" package](https://pkg.go.dev/errors)  
- func As(err error, target any) bool  
- func Is(err, target error) bool  
- func New(text string) error  
- func Unwrap(err error) error  

[Errors - Effective Go](https://go.dev/doc/effective_go#errors)
- By convention, errors have type error, a simple built-in interface.

[Error Handling and Go blog](https://go.dev/blog/error-handling-and-go)
- The error type is an interface type. An error variable represents any value that can describe itself as a string. 
- The error type, as with all built in types, is predeclared in the universe block.

[FAQ on Exceptions](https://go.dev/doc/faq#exceptions)
- Go does not favor try / catch / finally
