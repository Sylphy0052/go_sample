package main

import (
	"fmt"
	"math/rand"
)

func main() {
	r1 := rand.New(rand.NewSource(1))
	r2 := rand.New(rand.NewSource(2))
	r3 := rand.New(rand.NewSource(3))
	for i := 0; i < 20; i++ {
		fmt.Print(r1.Intn(100), " ")
	}
	fmt.Println("")
	for i := 0; i < 20; i++ {
		fmt.Print(r2.Intn(100), " ")
	}
	fmt.Println("")
	for i := 0; i < 20; i++ {
		fmt.Print(r3.Intn(100), " ")
	}
	fmt.Println("")
}
