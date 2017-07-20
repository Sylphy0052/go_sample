package main

import "fmt"

func insertSort(array []int) {
  for i := 1; i < len(array); i++ {
    tmp := array[i]
    j := i - 1
    for ;j >= 0 && tmp < array[j]; j-- {
      array[j + 1] = array[j]
    }
    array[j + 1] = tmp
  }
}

func main() {
  a := []int{5,6,4,7,3,8,2,9,1,0}
  b := []int{9,8,7,6,5,4,3,2,1,0}
  c := []int{0,1,2,3,4,5,6,7,8,9}
  insertSort(a)
  insertSort(b)
  insertSort(c)
  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}
