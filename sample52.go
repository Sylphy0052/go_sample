package main

import (
  "fmt"
)

func sumOf(f func(int) int, m, n int) int {
  a := 0
  for i := m; i <= n; i++ {
    a += f(i)
  }
  return a
}

func main() {
  // 匿名関数
  square := func(x int) int { return x * x }
  fmt.Println(sumOf(square, 0, 100))
  fmt.Println(sumOf(func(x int) int { return x * x * x }, 0, 100))
}
