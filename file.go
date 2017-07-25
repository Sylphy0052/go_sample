package main

import (
  "os"
  "fmt"
)

func main() {
  input, err := os.Open("testin.txt")
  if err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
  output, err := os.Create("testout.txt")
  buff := make([]byte, 256)
  for {
    c, _ := input.Read(buff)
    if c == 0 { break }
    output.Write(buff[ : c])
  }
  input.Close()
  output.Close()
}
