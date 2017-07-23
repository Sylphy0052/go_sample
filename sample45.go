package main

import (
  "fmt"
  "os"
)

var count int

func permutation(n int, perms []int, used []bool) {
  if n == len(perms) {
    fmt.Println(perms)
    count++
  } else {
    for i := 1; i <= len(perms); i++ {
      if used[i] == true { continue }
      used[i] = true
      perms[n] = i
      permutation(n + 1, perms, used)
      used[i] = false
    }
  }
}

func reader() (n int) {
  fmt.Println("Input number")
  fmt.Scan(&n)
  if n < 2 {
    os.Exit(1)
  }
  return n
}

func main() {
  n := reader()
  permutation(0, make([]int, n), make([]bool, n + 1))
  fmt.Println(count)
}
