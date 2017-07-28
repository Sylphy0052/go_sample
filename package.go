package main

import (
	"fmt"

	"./bar"
	"./foo"
)

func main() {
	foo.Test()
	fmt.Println("foo.A =", foo.A)
	bar.Test()
	fmt.Println("bar.A =", bar.A)
}
