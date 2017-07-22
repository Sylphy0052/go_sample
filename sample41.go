package main

import (
	"fmt"
	"os"
)

func f(n int) int {
	if n != 1 {
		return n * f(n-1)
	} else {
		return n
	}
}

func main() {
	fmt.Println("Input number")
	var n int
	fmt.Scan(&n)
	if n < 1 {
		fmt.Println("Input error")
		os.Exit(1)
	}
	fmt.Println(f(n))
}
