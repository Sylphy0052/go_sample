package main

import (
	"fmt"
	"math"
)

// 左端
func leftPoint(n int) float64 {
	w := 1.0 / float64(n)
	s := 0.0
	for i := 0; i < n; i++ {
		x := float64(i) * w
		s += 4.0 / (1.0 + x*x)
	}
	return s * w
}

// 中点則
func midPoint(n int) float64 {
	w := 1.0 / float64(n)
	s := 0.0
	for i := 1; i <= n; i++ {
		x := (float64(i) - 0.5) * w
		s += 4.0 / (1.0 + x*x)
	}
	return s * w
}

func main() {
	n := 10
	for i := 1; i <= 9; i++ {
		fmt.Println("-----", n, "-----")
		pi := midPoint(n)
		fmt.Println(pi, math.Pi-pi)
		pi = leftPoint(n)
		fmt.Println(pi, math.Pi-pi)
		n *= 10
	}
}
