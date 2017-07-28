package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 20; i++ {
		fmt.Print(rand.Intn(100), " ") // 同じ結果
	}
	fmt.Println("")
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		fmt.Print(rand.Intn(100), " ") // 違う結果
	}
	fmt.Println("")
}
