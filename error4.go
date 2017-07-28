package main

import "fmt"

// エラーの定義
type MyError struct {
	msg string
}

// エラーの生成
func newMyError(s string) *MyError {
	err := new(MyError)
	err.msg = s
	return err
}

// メソッドの定義
func (err *MyError) Error() string {
	return err.msg
}

func baz1() {
	panic(newMyError("oops!"))
}

func baz2() {
	panic("oops!")
}

func bar(f func()) {
	defer fmt.Println("bar end")
	fmt.Println("bar start!")
	f()
}

func foo(f func()) {
	defer func() {
		fmt.Println("foo end")
		err := recover()
		if err != nil {
			v := err.(*MyError)
			fmt.Println(v)
		}
	}()
	fmt.Println("foo start!")
	bar(f)
}

func main() {
	fmt.Println("main start!")
	foo(baz1)
	foo(baz2)
	fmt.Println("main end")
}
