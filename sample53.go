package main

import (
  "fmt"
)

// func square(m int) int {
//   return m * m
// }
//
// func cube(m int) int {
//   return m * m * m
// }

func mapcar(f func(int) int, a []int) []int {
  buff := make([]int, 0)
  for _, v := range a {
    buff = append(buff, f(v))
  }
  return buff
}

func removeIf(f func(int) bool, a []int) []int {
  buff := make([]int, 0)
  for i, v := range a {
    if f(v) { buff = append(buff, a[i]) }
  }
  return buff
}

func foldl(f func(x, y int) int, a int, ary []int) int {
  for _, v := range ary {
    a = f(a, v)
  }
  return a
}

func foldr(f func(x, y int) int, a int, ary []int) int {
  for i := len(ary) - 1; i >= 0; i-- {
    a = f(ary[i], a)
  }
  return a
}

func main() {
  square := func(x int) int { return x * x }
  cube := func(x int) int { return x * x * x }
  a := []int{1,2,3,4,5,6,7,8}
  b := mapcar(square, a)
  c := mapcar(cube, a)
  fmt.Println(b)
  fmt.Println(c)

  isEven := func(x int) bool { if x % 2 == 0 { return true } else { return false }}
  isOdd := func(x int) bool { if x % 2 != 0 { return true } else { return false }}
  d := removeIf(isEven, a)
  e := removeIf(isOdd, a)
  fmt.Println(d)
  fmt.Println(e)

  add := func(x, y int) int { return x + y }
  fmt.Println(foldl(add, 0, a))
  fmt.Println(foldr(add, 0, a))
}
