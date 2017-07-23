package main

import (
  "fmt"
  "strconv"
)

type Stringer interface {
  String() string
}

type myInt int

func (n myInt) String() string {
  return "myInt;" + strconv.Itoa(int(n))
}

func main() {
  n := myInt(10)
  m := myInt(20)
  fmt.Println(n + m)
  fmt.Println(n * m)
}
