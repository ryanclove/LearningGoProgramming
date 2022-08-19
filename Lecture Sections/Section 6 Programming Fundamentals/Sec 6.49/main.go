package main

import "fmt"

// Iota
// (pronounced eye-oh-tuh)
// A pre-declared identifier

const (
	a = iota
	b = iota
	c = iota
)

const (
	d = iota
	e
	f
)

const (
	g = iota
	h
	i
	j = iota
	k
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

	iota012()

	iota01234()
}

func iota012() {
	fmt.Println("")
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Printf("%T\n", d)
	fmt.Printf("%T\n", e)
	fmt.Printf("%T\n", f)
}

func iota01234() {
	fmt.Println("")
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println("iota auto-increments")
}
