# go doc

## Summary
    
- Godoc extracts and generates documentation for Go programs. It has two modes
    - **without -http flag**
        - command-line mode; prints text documentation to standard out and exits
        - **-src** flag
            - godoc prints the exported interface of a package in Go source form, or the implementation of a specific exported language
    - **with -http flag**
        - runs as a web server and presents the documentation as a web page
        - **godoc -http=:8080**
            - http://localhost:8080/

## link

- [godoc](https://pkg.go.dev/golang.org/x/tools/cmd/godoc?utm_source=godoc)
    ○ command to read documentation at the command line
    ○ also can run a local server showing documentation
    ○ [command documentation](https://go.dev/doc/cmd)

