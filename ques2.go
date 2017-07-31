// 素因数分解

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

func printFactor(num int, a []int) {
	fmt.Print(num, " = ")
	for i, v := range a {
		if i != 0 {
			fmt.Print("* ")
		}
		fmt.Print(v, " ")
	}
	fmt.Println("")
}

func primeFactor(n int) {
	var a []int
	num := n
	f := func(x, y int) bool {
		if x%y == 0 {
			return true
		} else {
			return false
		}
	}
	for i := range sieve(n) { // 素数
		if n < i {
			printFactor(num, a)
			return
		}
		for {
			if f(n, i) {
				a = append(a, i)
				n = n / i
			} else {
				break
			}
		}
	}
}

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	fmt.Print("Input number : ")
	for stdin.Scan() {
		n := stdin.Text()
		if n == "q" {
			os.Exit(1)
		}
		i, _ := strconv.Atoi(n)
		primeFactor(i)
		fmt.Print("Input number : ")
	}
}
