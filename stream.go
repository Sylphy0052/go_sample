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

// 数列の生成(2)

// 数列を無限に生成
func makeNum(n int) Stream {
	s := make(Stream)
	go func() {
		for {
			s <- n
		}
	}()
	return s
}

// フィボナッチ数列
func makeFibo() Stream {
	s := make(Stream)
	go func() {
		a, b := 1, 1
		for {
			s <- a
			a, b = b, a+b
			if a < 0 { // オーバーフロー
				break
			}
		}
		close(s)
	}()
	return s
}

func main() {
	s0 := makeInt(1, 20)
	for x := range s0 {
		fmt.Print(x, " ")
	}
	fmt.Println("")
	s1 := makeNum(1)
	for i := 0; i < 10; i++ {
		fmt.Print(<-s1, " ")
	}
	fmt.Println("")
	s2 := makeFibo()
	for x := range s2 {
		fmt.Print(x, " ")
	}
	fmt.Println("")
}
