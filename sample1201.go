package main

import(
  "fmt"
  "time"
)

func test(n int, name string) {
  for i := 1; i <= n; i++ {
    fmt.Println(i, name)
    time.Sleep(500 * time.Millisecond)
  }
}

func main() {
  test(5, "foo")
  test(5, "bar")
}
