package main

import (
    "fmt"
    "math"
)
// global定数
const s string = "const"

func main() {
  fmt.Println(s)

  // 定数式は任意の精度
  const n = 500000000
  const d = 3e20 / n
  fmt.Println(d)
  // 明示的キャスト
  fmt.Println(int64(d))
  // nはfloat64として認識される
  fmt.Println(math.Sin(n))
}
