package main

import (
	"fmt"
	"runtime"
	"time"
)

const (
	N = 1000000000
	W = 1.0 / float64(N)
)

func leftPoint(n, m int) float64 {
	s := 0.0
	for i := n; i < m; i++ {
		x := float64(i) * W
		s += 4.0 / (1.0 + x*x)
	}
	return s * W
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	for i := 1; i <= 8; i *= 2 {
		fmt.Println("-----", i, "-----")
		ch := make(chan float64, i)
		s := time.Now()
		k := N / i
		for j := 0; j < i; j++ {
			go func(n, m int) {
				ch <- leftPoint(n, m)
			}(j*k, (j+1)*k)
		}
		sum := 0.0
		for j := i; j > 0; j-- {
			sum += <-ch
		}
		e := time.Now().Sub(s)
		fmt.Println(sum)
		fmt.Println(e)
	}
}
