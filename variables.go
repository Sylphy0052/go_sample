package main

import "fmt"

func main() {
  var a string = "string"
  fmt.Println(a)

  // 複数の変数を宣言できる
  var b, c int = 1, 2
  fmt.Println(b, c)

  // 型推測
  var d = true
  fmt.Println(d)

  // 初期化しないとゼロ値が入る
  var e int
  fmt.Println(e)

  // 宣言かつ初期化
  // var f string = "short"
  f := "short"
  fmt.Println(f)
}
