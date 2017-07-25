package main

import "fmt"

// リクエストの定義
type Req struct {
  Color string
  Reply chan<- int
}

// リクエストの生成
func newReq(color string, ch chan int) *Req {
  req := new(Req)
  req.Color = color
  req.Reply = ch
  return req
}

// color の送信
// チャネルchにn回colorを渡す
func sendColor(n int, color string, ch chan<- *Req) {
  in := make(chan int)
  v := newReq(color, in)
  for ; n > 0; n-- {
    ch <- v // 送信
    <- in // データが来るまで待つ
  }
  ch <- nil
}

// color の受信
func receiveColor(n int, ch <-chan *Req) {
  for n > 0 {
    req := <- ch
    if req == nil {
      n--
      } else {
        fmt.Println(req.Color)
        req.Reply <- 0
      }
    }
  }

  func main() {
    ch := make(chan *Req)
    go sendColor(8, "red", ch)
    go sendColor(7, "blue", ch)
    go sendColor(6, "green", ch)
    receiveColor(3, ch)
  }
