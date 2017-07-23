package main

import(
  "fmt"
)

// ジェネレータを作る関数
func makeGen() func() int {
  prevNum := -1
  return func() int {
    prevNum += 2
    return prevNum
  }
}

func main() {
  g1 := makeGen()
  for i := 0; i < 8; i++ {
    fmt.Println(g1())
  }
  g2 := makeGen()
  for i := 0; i < 8; i++ {
    fmt.Println(g1())
  }
  for i := 0; i < 8; i++ {
    fmt.Println(g2())
  }
}
