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

func foo(n int) (m int) {
	defer func() {
		fmt.Println("foo end")
		err := recover()
		if err != nil {
			fmt.Println(err)
			m = -1
		}
	}()
	fmt.Println("foo start!")
	bar()
	m = n * n
	return
}

func main() {
	fmt.Println("main start!")
	fmt.Println(foo(10))
	fmt.Println("main end")
}
