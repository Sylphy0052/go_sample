package main

import (
	"fmt"

	. "./bar"
	f "./foo"
)

func main() {
	f.Test()
	fmt.Println("f.A =", f.A)
	Test()
	fmt.Println("A =", A)
}
