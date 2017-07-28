package main

import "fmt"

type Stream chan int

// 数列の生成
func makeInt(n, m int) Stream {
	s := make(Stream)
	go func() {
		for i := n; i <= m; i++ {
			s <- i
		}
		close(s)
	}()
	return s
}

// フィルター
func streamFilter(f func(int) bool, in Stream) Stream {
	s := make(Stream)
	go func() {
		for {
			x, ok := <-in
			if !ok {
				break
			}
			if f(x) {
				s <- x
			}
		}
		close(s)
	}()
	return s
}

func filter(n int, in Stream) Stream {
	return streamFilter(func(x int) bool { return x%n != 0 }, in)
}

// sieve : エラトステネスのふるい
func sieve(n int) Stream {
	s := make(Stream)
	go func() {
		in := makeInt(2, n) // 2 ~ 500まで
		for {
			x, ok := <-in
			if !ok {
				break
			}
			s <- x
			if x*x <= n { // x <= root(n)
				in = filter(x, in)
			}
		}
		close(s)
	}()
	return s
}

func main() {
	for x := range sieve(500) {
		fmt.Print(x, " ")
	}
	fmt.Println("")
}
