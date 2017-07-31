package main

import "fmt"

func sumofInt(n, m int) int {
	result := 0
	for i := n; i <= m; i++ {
		result += i
	}
	return result
}

func sumofSquare(n, m int) int {
	result := 0
	for i := n; i <= m; i++ {
		result += i * i
	}
	return result
}

func sumofCube(n, m int) int {
	result := 0
	for i := n; i <= m; i++ {
		result += i * i * i
	}
	return result
}

func main() {
	n, m := 1, 10
	fmt.Println(sumofInt(n, m))
	fmt.Println(sumofSquare(n, m))
	fmt.Println(sumofCube(n, m))
}
