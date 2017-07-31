package main

import "fmt"

func sum(n, m int, f func(int) int) int {
	ret := 0
	for i := n; i <= m; i++ {
		ret += f(i)
	}
	return ret
}

func sumofInt(n, m int) int {
	return sum(n, m, func(x int) int { return x })
}

func sumofSquare(n, m int) int {
	return sum(n, m, func(x int) int { return x * x })
}

func sumofCube(n, m int) int {
	return sum(n, m, func(x int) int { return x * x * x })
}

func main() {
	n, m := 1, 10
	fmt.Println(sumofInt(n, m))
	fmt.Println(sumofSquare(n, m))
	fmt.Println(sumofCube(n, m))
}
