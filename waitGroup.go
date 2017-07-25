package main

import(
  "fmt"
  "time"
  "sync"
)

func test(n int, name string, wg *sync.WaitGroup) {
  for i := 1; i <= n; i++ {
    fmt.Println(i, name)
    time.Sleep(500 * time.Millisecond)
  }
  wg.Done() // wgを1減らす
}

func main() {
  var wg sync.WaitGroup
  wg.Add(3) // 3つgoroutineを待ち合わせる
  go test(6, "foo", &wg)
  go test(4, "bar", &wg)
  go test(8, "baz", &wg)
  wg.Wait() // wgが0になるまで待つ
}
