package main

import (
	"fmt"
	"log"
	"os"
)

/*
Printing and logging

You have a few options to choose from when it comes to printing out, or logging, an error
message:
	● fmt.Println()
	● log.Println()
	● log.Fatalln()
		○ os.Exit()
	● log.Panicln()
		○ deferred functions run
		○ can use “recover”
	● panic()
		func panic(v any) - https://pkg.go.dev/builtin#panic
*/

func main() {
	// =======================================================================================
	//                      Create a log file which uses log package
	// =======================================================================================

	// Note: Println formats using the default formats for its operands and writes to standard output

	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f) // a file implements a writer interface

	// =======================================================================================
	//    Using log package -  Log an error to log.txt file when opening a non-existent file
	// =======================================================================================

	f2, err := os.Open("no-file.txt")
	if err != nil {
		// gives a timestamp of error
		log.Println("Log print of error:", err)
	}
	defer f2.Close()

	fmt.Println("Logs are in log.txt in this directory")

	// ======================================================================================================================================
	//                              Comment out Fatal Error to see Panic Error, as Fatal is coming first and exits the program
	// ======================================================================================================================================

	// =======================================================================================
	//                        Fatal Error (foo func is defered here)
	// =======================================================================================

	/*
		the Fatal functions call os.Exit(1) after writing the log message
		Fatalln is equivalent to Println() followed by a call to os.Exit(1)
	*/

	defer foo()
	log.Fatalln(err)

	// =======================================================================================
	//                                     Panic Error
	// =======================================================================================

	/*
		The panic built-in function stops normal execution of the current goroutine.
		When a function F calls panic, normal execution of F stops immediately. Any functions
		whose execution was deferred by F are run in the usual way, and then F returns to its caller.
	*/

	_, err = os.Open("no-file.txt")
	if err != nil {
		panic(err)
	}

}

func foo() {
	fmt.Println("When os.Exit() is called, deferred functions don't run")
}
