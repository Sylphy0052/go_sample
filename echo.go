package main

import "os"

func main() {
  buff := make([]byte, 256)
  for {
    c, _ := os.Stdin.Read(buff)
    if c == 0 { break }
    os.Stdout.Write(buff[ : c])
  }
}
