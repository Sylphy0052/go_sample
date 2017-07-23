package main

import(
  "fmt"
)

// 引数のスライス全てに渡した関数を適用する関数を返す
func mapcar(f func(x int) int) func([]int) []int {
  return func(ary []int) []int {
    buff := make([]int, len(ary))
    for i, v := range ary {
      buff[i] = f(v)
    }
    return buff
  }
}

func main() {
  square := func(x int) int { return x * x }
  cube := func(x int) int { return x * x * x }

  a := []int{1,2,3,4,5,6,7,8}
  // []intを引数にして全ての[]intにsquare関数を適用する関数
  squareAry := mapcar(square)
  // []intを引数にして全ての[]intにcube関数を適用する関数
  cubeAry := mapcar(cube)
  fmt.Println(squareAry(a))
  fmt.Println(cubeAry(a))
  fmt.Println(mapcar(square)(a))
  fmt.Println(mapcar(cube)(a))
}
