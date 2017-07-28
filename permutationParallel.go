package main

import (
	"fmt"
	"runtime"
	"time"
)

// n と等しい要素があるか
func member(n int, xs []int) bool {
	for _, x := range xs {
		if n == x {
			return true
		}
	}
	return false
}

// 順列の生成
func permutation(f func([]int), n, m, k int, xs []int) {
	if len(xs) == k {
		f(xs)
	} else {
		for i := n; i <= m; i++ {
			if !member(i, xs) {
				permutation(f, n, m, k, append(xs, i))
			}
		}
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	for i := 10; i <= 12; i++ {
		fmt.Println("-----", i, "-----")
		s := time.Now()
		// 逐次処理
		func() {
			c := 0
			a := make([]int, 0, i)
			permutation(func(_ []int) { c++ }, 1, i, i, a)
			fmt.Println(c)
		}()
		e := time.Now().Sub(s)
		fmt.Println(e)
		s = time.Now()
		ch := make(chan int, i)
		// 並列処理
		for j := 0; j < i; j++ {
			go func(i, j int) {
				c := 0
				a := make([]int, 0, i)
				a = append(a, j+1)
				permutation(func(_ []int) { c++ }, 1, i, i, a)
				ch <- c
			}(i, j)
		}
		sum := 0
		for j := 0; j < i; j++ {
			sum += <-ch
		}
		fmt.Println(sum)
		e = time.Now().Sub(s)
		fmt.Println(e)
	}
}
