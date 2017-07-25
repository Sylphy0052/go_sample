// タイムアウト処理
package main

import (
	"fmt"
	"time"
)

func fibo(n int) int {
	if n < 2 {
		return 1
	} else {
		return fibo(n-2) + fibo(n-1)
	}
}

func main() {
	ch := make(chan int, 5)
	for _, n := range []int{41, 40, 38, 36, 34} {
		go func(x int) {
			ch <- fibo(x)
		}(n)
	}
	for i := 5; i > 0; {
		select {
		case n := <-ch:
			fmt.Println(n)
			i--
		case <-time.After(time.Second / 10.0):
			fmt.Println("Timeout")
			i = 0
		}
	}
}
