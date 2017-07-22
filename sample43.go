// x ** y = (x ** (y / 2)) ** 2       （n は偶数）
// x ** y = ((x ** (y / 2)) ** 2) * x （n は奇数）

package main

import (
	"fmt"
)

// func pow(a ,b int) int {
// 	if b == 0 {
// 		return 1
// 	} else {
// 		return a * pow(a, b-1)
// 	}
// }

func pow(a float64, b int) float64 {
	if b == 0 {
		return 1
	} else if z := pow(a, b/2); b%2 == 0 { // 偶数
		return z * z
	} else { // 奇数
		return z * z * a
	}
}

func test() {
	fmt.Println(pow(2, 16))
	fmt.Println(pow(2, 32))
	fmt.Println(pow(2, 64))
}

func main() {
	test()
}
