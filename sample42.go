// [ユークリッドの互除法]
// 負でない整数 a と b (a > b) で、a を b で割った余りを r とする。
// このとき、a と b の最大公約数は b と r の最大公約数に等しい。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readTwoInt() (int, int) {
	result := make([]int, 2)
	sc.Split(bufio.ScanWords)
	fmt.Println("Input two number e.g., 10 5")
	for i := 0; i < 2; i++ {
		sc.Scan()
		result[i], _ = strconv.Atoi(sc.Text())
		if result[i] < 1 {
			os.Exit(1)
		}
	}
	return result[0], result[1]
}

// func gcd(a, b int) int {
// 	// aの方が大きいように
// 	if a < b {
// 		a, b = b, a
// 	}
// 	r := a % b
// 	if r == 0 {
// 		return b
// 	} else {
// 		return gcd(r, b)
// 	}
// }

func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func test() {
	fmt.Println(gcd(42, 30))
	fmt.Println(gcd(15, 70))
	fmt.Println(lcm(5, 7))
	fmt.Println(lcm(14, 35))
}

func main() {
	// a, b := readTwoInt()
	// fmt.Println(gcd(a, b))
	test()
}
