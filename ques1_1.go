package main

import "fmt"

const (
	INT    = 0
	SQUARE = 1
	CUBE   = 2
)

func type_func(t int) func(int) int {
	switch t {
	case INT:
		return func(i int) int { return i }
	case SQUARE:
		return func(i int) int { return i * i }
	case CUBE:
		return func(i int) int { return i * i * i }
	}
	return nil
}

func sum(n, m, t int) int {
	f := type_func(t)
	ret := 0
	for i := n; i <= m; i++ {
		ret += f(i)
	}
	return ret
}

func sumofInt(n, m int) int {
	return sum(n, m, INT)
}

func sumofSquare(n, m int) int {
	return sum(n, m, SQUARE)
}

func sumofCube(n, m int) int {
	return sum(n, m, CUBE)
}

func main() {
	n, m := 1, 10
	fmt.Println(sumofInt(n, m))
	fmt.Println(sumofSquare(n, m))
	fmt.Println(sumofCube(n, m))
}
