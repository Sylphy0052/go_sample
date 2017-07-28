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

// マッピング
func streamMap(f func(int) int, in Stream) Stream {
	s := make(Stream)
	go func() {
		for {
			x, ok := <-in
			if !ok {
				break
			}
			s <- f(x)
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

func main() {
	// 1 ~ 10を二乗にして表示
	square := func(x int) int { return x * x }
	s3 := streamMap(square, makeInt(1, 10))
	for x := range s3 {
		fmt.Print(x, " ")
	}
	fmt.Println("")

	// 1 ~ 20 で奇数だけ表示
	isOdd := func(x int) bool { return x%2 != 0 }
	s4 := streamFilter(isOdd, makeInt(1, 20))
	for x := range s4 {
		fmt.Print(x, " ")
	}
	fmt.Println("")
}
