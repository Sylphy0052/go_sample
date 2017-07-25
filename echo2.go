package main

import (
  "os"
  "io"
  "bufio"
)

func main() {
  r := bufio.NewReader(os.Stdin)
  w := bufio.NewWriter(os.Stdout)
  for {
    s, err := r.ReadString('\n')
    if err == io.EOF { break }
    w.WriteString(s)
    w.Flush()
  }
}
