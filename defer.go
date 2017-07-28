package main

import "fmt"

func bar() {
	defer fmt.Println("bar end") // 関数終了時実行される
	fmt.Println("bar start!")
}

func foo() {
	defer fmt.Println("foo end")
	fmt.Println("foo start!")
	bar()
}

func main() {
	foo()
}
