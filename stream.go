package main

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
func makeNum(n int) Stream {
	s := make(Stream)
	go func() {
		a, b := 1, 1
		for {
			s <- n
		}
	}()
	return s
}

// フィボナッチ数列
