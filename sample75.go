package main

import (
  "fmt"
)

type Cell struct {
  item int
  next *Cell
}

type List struct {
  top *Cell
}

func newCell(x int, cp *Cell) *Cell {
  newCp := new(Cell)
  newCp.item, newCp.next = x, cp
  return newCp
}

func newList() *List {
  lst := new(List)
  lst.top = new(Cell)
  return lst
}

func (cp *Cell) nthCell(n int) *Cell {
  i := -1
  for cp != nil {
    if i == n { return cp }
    i++
    cp = cp.next
  }
  return nil
}

func (lst *List) nth(n int) (int, bool) {
  cp := lst.top.nthCell(n)
  if cp == nil { return 0, false }
  return cp.item, true
}

func (lst *List) insertNth(n, x int) bool {
  cp := lst.top.nthCell(n-1)
  if cp == nil { return false }
  cp.next = newCell(x, cp.next)
  return true
}

func (lst *List) deleteNth(n int) bool {
  cp := lst.top.nthCell(n-1)
  if cp == nil || cp.next ==nil { return false }
  cp.next = cp.next.next
  return true
}

func (lst *List) isEmpty() bool {
  return lst.top.next == nil
}

func (lst *List) printList() {
  cp := lst.top.next
  for ; cp != nil; cp = cp.next {
    fmt.Print(cp.item, " ")
  }
  fmt.Println("")
}

func main() {
  a := newList()
  for i := 0; i < 4; i++ {
    fmt.Println(a.insertNth(i, i))
  }
  a.printList()
  for i := 0; i < 5; i++ {
    n, ok := a.nth(i)
    fmt.Println(n, ok)
  }
  for !a.isEmpty() {
    a.deleteNth(0)
    a.printList()
  }
}
