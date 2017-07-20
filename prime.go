package main

import (
  "fmt"
)

func main() {
  var array [100]bool
  for i := 2; i < len(array); i++ {
    if array[i] { continue }
    fmt.Print(i, " ")
    for j := i * 2; j < len(array); j += i {
      array[j] = true
    }
  }
  fmt.Println()
}
