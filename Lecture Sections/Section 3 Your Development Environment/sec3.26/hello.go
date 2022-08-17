package hello

func Hello() string {
	return "Hello, world."
}

/*
Supposed to use:
	go list -m all
to see imported package/dependency versions

then do
	go get golang.org/x/text
should PASS

check go list -m all again
and the golang.org/x/text package should be upgraded
to the latest tag version and go.mod
should reflect the updated version too
*/
