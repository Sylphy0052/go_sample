package main

import (
	"fmt"
	"runtime"
	"time"
)

// フィボナッチ関数
func fibo(n int) int {
	if n < 2 {
		return 1
	} else {
		return fibo(n-1) + fibo(n-2)
	}
}

func main() {
	ch := make(chan int, 10)
	for n := 1; n <= 10; n++ {
		fmt.Println("-----", n, "-----")
		runtime.GOMAXPROCS(n)
		s := time.Now()
		for i := 0; i < 8; i++ {
			go func() {
				ch <- fibo(40)
			}()
		}
		for i := 8; i > 0; i-- {
			fmt.Print(<-ch, " ")
		}
		e := time.Now().Sub(s)
		fmt.Println(e)
	}
}
