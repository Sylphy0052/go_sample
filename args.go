package main

import (
  "fmt"
  "os"
)

func main() {
  // fmt.Println(os.Args)
  for i, v := range os.Args {
    fmt.Println(v, os.Args[i])
  }
}
