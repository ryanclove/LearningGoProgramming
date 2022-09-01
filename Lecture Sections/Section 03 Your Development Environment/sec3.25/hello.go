package hello

// Hello returns a greeting.
func Hello() string {
	return "Hello, world."
}

/* In the time of the video, he used the following lines:

package hello

import "rsc.io/quote"

func Hello() string {
    return quote.Hello()
}

******************************************
I was getting an error on the import statement, saying it could not be
found in the GOROOT

Realized quote is now a built-in package/library for the latest
version of Go

Can be found in this directory of my desktop:
C:\Program Files\Go\src\cmd\go\testdata\mod

So above is my necessary adjustments

*/
