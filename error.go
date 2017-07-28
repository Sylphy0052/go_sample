package main

import (
	"errors"
	"fmt"
)

// 階乗
func fact(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("fact : domain error")
	}
	a := 1
	for ; n > 1; n-- {
		a *= n
	}
	return a, nil
}

func main() {
	for x := 10; x >= -1; x-- {
		v, err := fact(x)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(v)
		}
	}
}
