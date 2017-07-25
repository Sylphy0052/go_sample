package main

import (
  "fmt"
  "strconv"
  "time"
)

func test1(n int, ch, quit chan<- int) {
  for ; n > 0; n-- {
    ch <- n
    time.Sleep(500 * time.Millisecond)
  }
  quit <- 0
}

func test2(n int, ch chan<- float64, quit chan<- int) {
  for ; n > 0; n-- {
    ch <- float64(n) / 10.0
    time.Sleep(250 * time.Millisecond)
  }
  quit <- 0
}

func test3(n int, ch chan<- string, quit chan<- int) {
  for ; n > 0; n-- {
    ch <- strconv.Itoa(n * 10)
    time.Sleep(750 * time.Millisecond)
  }
  quit <- 0
}

func main() {
  ch1 := make(chan int)
  ch2 := make(chan float64)
  ch3 := make(chan string)
  quit := make(chan int)
  go test1(6, ch1, quit)
  go test2(8, ch2, quit)
  go test3(4, ch3, quit)
  for n := 3; n > 0; {
    select {
      case c := <- ch1: fmt.Println(c)
      case c := <- ch2: fmt.Println(c)
      case c := <- ch3: fmt.Println(c)
      case <- quit: n--
    default:
      fmt.Println("None")
      time.Sleep(250 * time.Millisecond)
    }
  }
}
