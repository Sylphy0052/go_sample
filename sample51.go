package main

import (
  "fmt"
)

func sumOfInteger(m, n int) int {
  a := 0
  for i := m; i <= n;i++ {
    a += i
  }
  return a
}


func square(m int) int {
  return m * m
}

func sumOfSquare(m, n int) int {
  a := 0
  for i := m; i < n; i++ {
    a += square(i)
  }
  return a
}

func cube(m int) int {
  return m * m * m
}

func sumOfCube(m, n int) int {
  a := 0
  for i := m; i < n; i++ {
    a += cube(i)
  }
  return a
}

func sumOf(f func(int) int, m, n int) int {
  a := 0
  for i := m; i < n; i++ {
    a += f(i)
  }
  return a
}

func main() {
  fmt.Println(sumOfInteger(1, 100))
  fmt.Println(sumOfSquare(1, 100))
  fmt.Println(sumOfCube(1, 100))
  fmt.Println(sumOf(square, 1, 100))
  fmt.Println(sumOf(cube, 1, 100))
}
