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
	defer func() {
		fmt.Println("foo end")
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("foo start!")
	bar()
}

func main() {
	fmt.Println("main start!")
	foo()
	fmt.Println("main end")
}
