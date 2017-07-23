package main

import (
  "fmt"
  "math/rand"
  "time"
)

// []int をクイックソート
func quickSortInt(buff []int, low, high int) {
  pivot := buff[low + (high - low) / 2]
  i, j := low, high
  for {
    for pivot > buff[i] { i++ }
    for pivot < buff[j] { j-- }
    if i >= j { break }
    buff[i], buff[j] = buff[j], buff[i]
    i++
    j--
  }
  if low < i - 1  { quickSortInt(buff, low,   i - 1) }
  if high > j + 1 { quickSortInt(buff, j + 1, high) }
}

// ソートインターフェース
type SortI interface{
  Len() int
  Less(int, int) bool
  Swap(int, int)
}

// SortI によるクイックソート
func quickSortSub(data SortI, low, high int) {
  p := low + (high - low) / 2
  i, j := low, high
  for {
    for data.Less(i, p) { i++ }
    for data.Less(p, j) { j-- }
    if i >= j { break }
    data.Swap(i, j)
    switch {
      case p == i: p = j
      case p == j: p = i
    }
    i++
    j--
  }
  if low < i - 1  { quickSortSub(data, low,   i - 1) }
  if high > j + 1 { quickSortSub(data, j + 1, high) }
}

func quickSort(data SortI) {
  quickSortSub(data, 0, data.Len() - 1)
}

// インターフェースの実装
type IntArray []int

func (ary IntArray) Len() int {
  return len(ary)
}

func (ary IntArray) Less(i, j int) bool {
  return ary[i] < ary[j]
}

func (ary IntArray) Swap(i, j int) {
  ary[i], ary[j] = ary[j], ary[i]
}

// データを比較するインターフェース
type CmpI interface {
  Less(CmpI) bool // x < yならtrueを返す関数
}

// []CmpI をクイックソート
func quickSortCmpI(buff []CmpI, low, high int) {
  p := buff[low + (high - low) / 2]
  i, j := low, high
  for {
    for buff[i].Less(p) { i++ }
    for p.Less(buff[j]) { j-- }
    if i >= j { break }
    buff[i], buff[j] = buff[j], buff[i]
    i++
    j--
  }
  if low < i - 1  { quickSortCmpI(buff, low,   i - 1) }
  if high > j + 1 { quickSortCmpI(buff, j + 1, high) }
}

// CmpI インターフェースを実装する
type Cint int

func (n Cint) Less(m CmpI) bool {
  return n < m.(Cint)
}

// 簡単なテスト
func main() {
  a := make([]int,    1000000)
  b := make(IntArray, 1000000)
  c := make([]CmpI,   1000000)
  for i := 0; i < 1000000; i++ {
    x := rand.Int()
    a[i] = x
    b[i] = x
    c[i] = Cint(x)
  }
  s := time.Now()
  quickSortInt(a, 0, len(a) - 1)
  e := time.Now().Sub(s)
  fmt.Println(e)
  s = time.Now()
  quickSort(b)
  e = time.Now().Sub(s)
  fmt.Println(e)
  s = time.Now()
  quickSortCmpI(c, 0, len(a) - 1)
  e = time.Now().Sub(s)
  fmt.Println(e)
}
