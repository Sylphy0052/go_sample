package main

import "fmt"

// エラーの定義
type MyError struct {
	msg string
}

// エラーの生成
func newMyError(s string) *MyError {
	err := new(MyError) //
	err.msg = s         //
	return err          // return &MyError{s} でもよい
}

// メソッドの定義
func (err *MyError) Error() string {
	return err.msg
}

// 階乗
func fact(n int) (int, error) {
	if n < 0 {
		return 0, newMyError("fact : domain error")
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
