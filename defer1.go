package main

import "fmt"

func baz() {
	panic("oops!")
}

func bar() {
	defer fmt.Println("bar end")
	fmt.Println("bar start!")
	baz()
}

func foo() {
	defer fmt.Println("foo end")
	fmt.Println("foo start!")
	bar()
}

func main() {
	foo()
}
