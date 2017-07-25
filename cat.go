package main

import (
  "fmt"
  "os"
  "io"
  "bufio"
)

func cat(filename string) {
  file, err := os.Open(filename)
  if err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
  rd := bufio.NewReader(file)
  for {
    s, err := rd.ReadString('\n')
    if err == io.EOF { break }
    fmt.Print(s)
  }
  file.Close()
}

func main() {
  for _, name := range os.Args[1 : ] {
    cat(name)
  }
}
