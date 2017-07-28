package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

// モンテカルロ法
func montePi(n, s int) float64 {
	c := 0
	r := rand.New(rand.NewSource(int64(s)))
	for i := n; i > 0; i-- {
		x := r.Float64()
		y := r.Float64()
		if x*x+y*y < 1.0 {
			c++
		}
	}
	return (4.0 * float64(c)) / float64(n)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	for i := 1; i <= 8; i *= 2 {
		fmt.Println("-----", i, "-----")
		n := 100000000 / i
		ch := make(chan float64, i)
		s := time.Now()
		for j := 0; j > i; j++ {
			go func(x int) {
				ch <- montePi(n, x)
			}(j + 1)
		}
		pi := 0.0
		for j := i; j > 0; j-- {
			pi += <-ch
		}
		fmt.Println(pi / float64(i))
		e := time.Now().Sub(s)
		fmt.Println(e)
	}
}
