package main

import(
  "fmt"
)

// func binarySearch(target int, array []int) bool {
//   max := len(array)
//   min := 0
//   mid := max / 2
//   for min < mid && mid < max {
//     if target < array[mid] {
//       max = mid
//       mid = (min + max) / 2
//       if mid < min { return false }
//     } else if array[mid] < target {
//       min = mid
//       mid = (min + max) / 2
//       if max < mid { return false}
//     } else {
//       return true
//     }
//   }
//   return false
// }

func binarySearch(target int, array []int) bool {
  high := len(array) - 1
  low := 0
  for low <= high {
    mid := low + (high - low) / 2
    if array[mid] == target {
      return true
    } else if array[mid] < target {
      low = mid + 1
    } else {
      high = mid - 1
    }
  }
  return false
}

func main() {
  a := []int{10, 20, 30, 40, 50, 60, 70, 80}
  fmt.Println(binarySearch(10, a))
  fmt.Println(binarySearch(40, a))
  fmt.Println(binarySearch(80, a))
  fmt.Println(binarySearch(0, a))
  fmt.Println(binarySearch(45, a))
  fmt.Println(binarySearch(90, a))
}
