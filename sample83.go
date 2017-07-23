package main

import "fmt"

// 要素の型
type int interface{}

// セル
type Cell struct {
  item int
  next *Cell
}

// リスト
type List struct {
  top *Cell
}

// セルの生成
func newCell(x int, cp *Cell) *Cell {
  newcp := new(Cell)
  newcp.item, newcp.next = x, cp
  return newcp
}

// リストの生成
func newList() *List {
  lst := new(List)
  lst.top = new(Cell)
  return lst
}

// n 番目のセルを求める
func (cp *Cell) nthCell(n int) *Cell {
  i := -1
  for cp != nil {
    if i == n { return cp }
    i++
    cp = cp.next
  }
  return nil
}

// 参照
func (lst *List) nth(n int) (int, bool) {
  cp := lst.top.nthCell(n)
  if cp == nil { return nil, false }
  return cp.item, true
}

// 挿入
func (lst *List) insertNth(n, x int) bool {
  cp :=lst.top.nthCell(n - 1)
  if cp == nil { return false }
  cp.next = newCell(x, cp.next)
  return true
}

// 削除
func (lst *List) deleteNth(n int) bool {
  cp := lst.top.nthCell(n - 1)
  if cp == nil || cp.next == nil { return false }
  cp.next = cp.next.next
  return true
}

// リストは空か？
func (lst *List) isEmpty() bool {
  return lst.top.next == nil
}

// 表示
func (lst *List) printList() {
  cp := lst.top.next
  for ; cp != nil; cp = cp.next {
    fmt.Print(cp.item, " ")
  }
  fmt.Println("")
}

// 制限つき連結リスト
type FixedList struct {
  List
  size, limit int
}

func newFixedList(limit int) *FixedList {
  p := new(FixedList)
  p.top = new(Cell)
  p.limit = limit
  return p
}

func (p *FixedList) insertNth(n, x int) bool {
  if p.size >= p.limit { return false }
  result := p.List.insertNth(n, x)
  if result { p.size++ }
  return result
}

func (p *FixedList) deleteNth(n int) bool {
  result := p.List.deleteNth(n)
  if result { p.size-- }
  return result
}

// スタック
type Stack struct {
  content *List
}

// スタックの生成
func newStack() *Stack {
  st := new(Stack)
  st.content = newList()
  return st
}

// 挿入
func (st *Stack) push(x int) {
  st.content.insertNth(0, x)
}

// 削除
func (st *Stack) pop() (int, bool) {
  x, ok := st.content.nth(0)
  if ok { st.content.deleteNth(0) }
  return x, ok
}

// 参照
func (st *Stack) top() (int, bool) {
  return st.content.nth(0)
}

// スタックは空か？
func (st *Stack) isEmpty() bool {
  return st.content.isEmpty()
}

// キュー
type Queue struct {
  size int
  front, rear *Cell
}

// キューの生成
func newQueue() *Queue {
  return new(Queue)
}

// 挿入
func (q *Queue) enqueue(x int) {
  cp := newCell(x, nil)
  if q.size == 0 {
    q.front = cp
    q.rear = cp
  } else {
    q.rear.next = cp
    q.rear = cp
  }
  q.size++
}

// 削除
func (q *Queue) dequeue() (int, bool) {
  if q.size == 0 {
    return 0, false
  } else {
    x := q.front.item
    q.front = q.front.next
    q.size--
    if q.size == 0 { q.rear = nil }
    return x, true
  }
}

// キューは空か？
func (q *Queue) isEmpty() bool {
  return q.size == 0
}

// 簡単なテスト
func main() {
  // リスト
  a := newList()
  for i := 0; i < 8; i++ {
    fmt.Println(a.insertNth(i, i))
  }
  a.printList()
  for i := 0; i < 8; i++ {
    n, ok := a.nth(i)
    fmt.Println(n, ok)
  }
  for !a.isEmpty() {
    a.deleteNth(0)
    a.printList()
  }
  // 制限つきリスト
  b := newFixedList(6)
  for i := 0; i < 8; i++ {
    fmt.Println(b.insertNth(i, i))
  }
  b.printList()
  for i := 0; i < 8; i++ {
    fmt.Println(b.nth(i))
  }
  for !b.isEmpty() {
    b.deleteNth(0)
    b.printList()
  }
  // スタック
  st := newStack()
  for i := 0; i < 8; i++ {
    st.push(i)
    fmt.Println(st.top())
  }
  for !st.isEmpty() {
    fmt.Println(st.pop())
  }
  // キュー
  que := newQueue()
  for i := 0; i < 8; i++ {
    que.enqueue(i)
    fmt.Println(i)
  }
  for !que.isEmpty() {
    fmt.Println(que.dequeue())
  }
}
