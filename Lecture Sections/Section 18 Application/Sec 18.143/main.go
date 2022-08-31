package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// bcrypt

/*
Bcrypt is one of the tools you can use to protect user data. Using bcrypt,
we will gain further understanding as to how to read and implement code from the
standard library.
*/

/*
Not part of standard library, to add bcrypt package in the terminal found GOPATH = go-workspace
Then did `go install golang.org/x/crypto/...@latest`
Is now sitting in $C:\Users\ryanc\Documents\go-workspace\pkg\mod\golang.org\x and the crypto ones

!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
still have a package error, and each one i solve leads to another. So, i give up. Here is Todd's code below,
trust the process:
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
*/

func main() {
	s := `password123`
	// func GenerateFromPassword(password []byte, cost int) ([]byte, error)
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)

	loginPword1 := `password1234`

	// func CompareHashAndPassword(hashedPassword, password []byte) error
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPword1))
	if err != nil {
		fmt.Println("YOU CAN'T LOGIN")
		return
	}

	fmt.Println("You're logged in")
}
