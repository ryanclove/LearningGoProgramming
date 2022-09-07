# go doc

## Summary

go doc prints the documentation for a package, const, func, type, var, or method
- go doc accepts zero, one, or two **arguments**.
    - **zero**
        - prints package documentation for the package in the current directory
            - go doc
    - **one**
        - argument Go-syntax-like representation of item to be documented
            - fyi: <sym> also known as “identifier”
                - go doc <pkg>
                - go doc <sym>[.<method>]
                - go doc [<pkg>.]<sym>[.<method>]
                - go doc [<pkg>.][<sym>.]<method>
            - The first item in this list that succeeds is the one whose documentation is printed. If there is a symbol but no package, the package in the current directory is chosen. However, if the argument begins with a capital letter it is always assumed to be a symbol in the current directory.
    - **two**
        - first argument must be a full package path
            - go doc <pkg> <sym>[.<method>]

## link

- [go doc](https://pkg.go.dev/cmd/go)]
    ○ command to read documentation at the command line
    ○ notice the space in go doc
    
## terminal commands

- go

- go doc -cmd cmd/doc

- go help doc
- - allows us to see documentation at the terminal
