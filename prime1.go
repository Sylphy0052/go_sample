package main

import "fmt"

func isPrime(n int, table []int) bool {
  for _, v := range table {
    if n % v == 0 {
      return false
    }
  }
  return true
}

func main() {
  table := []int{2}
  for i := 3; i <= 100; i++ {
    if isPrime(i, table) {
      table = append(table, i)
    }
  }
  fmt.Println(table)
}
