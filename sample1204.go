package main

import (
  "fmt"
  "time"
)

func makeRoutine(code string, in <-chan int) chan int {
  out := make(chan int)
  go func(){
    for {
      <- in
      fmt.Print(code)
      time.Sleep(50 * time.Millisecond)
      out <- 0
    }
  }()
  return out
}

  func main(){
    ch1 := make(chan int)
    ch2 := makeRoutine("h", ch1)
    ch3 := makeRoutine("e", ch2)
    ch4 := makeRoutine("y", ch3)
    ch5 := makeRoutine("!", ch4)
    ch6 := makeRoutine(" ", ch5)
    for i := 0; i < 10; i++ {
      ch1 <- 0
      <- ch6
    }
    fmt.Println()
  }
