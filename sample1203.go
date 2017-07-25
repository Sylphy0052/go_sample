package main

import (
  "fmt"
  "time"
)

func test(n int, name string, c chan<- string) {
  for i := 1; i <= n; i++ {
    fmt.Println(i, name)
    time.Sleep(500 * time.Millisecond)
  }
  c <- name // nameをcに入れる
}

func main() {
  c := make(chan string) // チャネルの生成
  go test(6, "foo", c)
  go test(4, "bar", c)
  go test(8, "baz", c)
  for i := 0; i < 3; i++ {
    fmt.Println(<- c) // cに値が入ると出力
  }
}
