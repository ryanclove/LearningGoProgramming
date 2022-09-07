# Writing Documentation

## Summary
    
Documentation is a huge part of making software accessible and maintainable. Of course it
must be well-written and accurate, but it also must be easy to write and to maintain. Ideally, it
should be coupled to the code itself so the documentation evolves along with the code. The
easier it is for programmers to produce good documentation, the better for everyone.

- https://blog.golang.org/godoc-documenting-go-code
    - **godoc** parses Go source code - including comments - and produces documentation as HTML or plain text. The end result is documentation tightly coupled with the code it documents. For example, through godoc's web interface you can **navigate from a function's documentation to its implementation with one click**.
    - comments are just good comments, the sort you would want to read even if **godoc** didn't exist.
    - **to document**
        - a type, variable, constant, function, or package,
        - **write a comment** directly preceding its declaration, with no intervening blank line.
            - **begin with the name of the element**
            - for packages
                - first sentence appears in package list
                - if a large amount of documentation, place in its own file **doc.go**
                    - example: [package fmt](https://go.dev/src/fmt/doc.go)
    - the best thing about godoc's minimal approach is how easy it is to use. As a result, a lot of Go code, including all of the standard library, already follows the conventions.
- example
    - [errors package](https://go.dev/src/errors/errors.go#L9)



