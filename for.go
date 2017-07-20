package main

import (
  "fmt"
)

func main() {
  // 条件式のみ
  i := 1
  for i <= 3 {
    fmt.Println(i)
    i = i + 1
  }
  // 初期化;条件式;後処理
  for j := 7; j <= 9; j++ {
    fmt.Println(j)
  }
  // breakでforを抜ける
  for {
    fmt.Println("loop")
    break
  }
  // continueで次の実行に進む
  for n := 0; n <= 5; n++ {
    if n % 2 == 0 {
      continue
    }
    fmt.Println(n)
  }
}
