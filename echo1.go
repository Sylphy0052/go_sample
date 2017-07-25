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
    c, err := r.ReadByte()
    if err == io.EOF { break }
    w.WriteByte(c)
    if c == '\n' { w.Flush() }
  }
  w.Flush() // これが書き込み
}
