// 1 + 2 + 3 + ... + n = n(n+1)/ 2
// 1 + 22 + 32 + ... + n2 = n(n+1)(2n+1)/6
// 1 + 23 + 33 + ... + n3 = n2(n+1)2/4

package main

import "fmt"

func sumofInt(n, m int) int {
	f := func(x int) int { return x * (x + 1) / 2 }
	return f(m) - f(n-1)
}

func sumofSquare(n, m int) int {
	f := func(x int) int { return x * (x + 1) * (2*x + 1) / 6 }
	return f(m) - f(n-1)
}

func sumofCube(n, m int) int {
	f := func(x int) int { return x * x * (x + 1) * (x + 1) / 4 }
	return f(m) - f(n-1)
}

func main() {
	n, m := 1, 10
	fmt.Println(sumofInt(n, m))
	fmt.Println(sumofSquare(n, m))
	fmt.Println(sumofCube(n, m))
}
