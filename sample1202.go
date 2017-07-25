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
  // goroutine -> スレッド
  go test(5, "foo")
  test(5, "bar")
  // go test() こうするとmainが終了してしまうため，表示されなくなってしまう
}
