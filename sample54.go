package main

import(
  "fmt"
)

// クロージャー　→関数を返す関数
func foo(x int) func(int) int {
  return func(y int) int { return x * y }
}

func main() {
  foo10 := foo(10)
  foo20 := foo(20)
  fmt.Println(foo10(foo20(1)))
}
