package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

/*
Checking Errors

Write the code with errors before writing the code without errors. Always check for errors.
Always, always, always.*
(*almost always)
*/

func main() {
	n, err := fmt.Println("Hello")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(n) // number of bytes written
	// 5 for hello, 1 for newline character

	// uses scan
	exampleFromFmtPackage()

	// creating and writing to a file using os and io and string packages
	createFile()

	// open a file
	openFile()
}

func exampleFromFmtPackage() {
	var answer1, answer2, answer3 string

	fmt.Print("Name: ")
	_, err := fmt.Scan(&answer1)
	if err != nil {
		panic(err)
	}

	fmt.Print("Fav Food: ")
	_, err = fmt.Scan(&answer2)
	if err != nil {
		panic(err)
	}

	fmt.Print("Fav Sport: ")
	_, err = fmt.Scan(&answer3)
	if err != nil {
		panic(err)
	}

	fmt.Println(answer1, answer2, answer3)

}

func createFile() {
	f, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	r := strings.NewReader("James Bond")

	// copy the string to the file
	io.Copy(f, r)

	fmt.Println("Successfully written string to file")
}

func openFile() {
	f, err := os.Open("names.txt")
	// if file not in directory, returns error
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	// print string of byte slices in file to stdout
	fmt.Println(string(bs))
}
