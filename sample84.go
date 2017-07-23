package main

import "fmt"

// データ型の定義
type Queue struct {
  front, rear, cnt int
  buff []int
}

// キューの生成
func makeQueue(size int) *Queue {
  q := new(Queue)
  q.buff = make([]int, size)
  return q
}

// キューは空か
func (q *Queue) isEmpty() bool {
  return q.cnt == 0
}

// キューは満杯か
func (q *Queue) isFull() bool {
  return q.cnt == len(q.buff)
}

// データの挿入
func (q *Queue) enqueue(x int) bool {
  if q.isFull() { return false }
  q.buff[q.rear] = x
  q.cnt++
  q.rear++
  if q.rear >= len(q.buff) { q.rear = 0 }
  return true
}

// データを取り出す
func (q *Queue) dequeue() (int, bool) {
  if q.isEmpty() { return 0, false }
  x := q.buff[q.front]
  q.cnt--
  q.front++
  if q.front >= len(q.buff) { q.front = 0 }
  return x, true
}

// 先頭データの参照
func (q *Queue) top() (int, bool) {
  if q.isEmpty() { return 0, false }
  return q.buff[q.front], true
}

// キューを空にする
func (q *Queue) clear() {
  q.front = 0
  q.rear = 0
  q.cnt = 0
}

// キューの要素数を返す
func (q *Queue) length() int {
  return q.cnt
}

// 簡単なテスト
func main() {
  q := makeQueue(10)
  fmt.Println(q.isEmpty())
  fmt.Println(q.length())
  for i := 0; i < 10; i++ {
    q.enqueue(i)
  }
  fmt.Println(q.isFull())
  fmt.Println(q.length())
  for !q.isEmpty() {
    fmt.Println(q.dequeue())
  }
  fmt.Println(q.isEmpty())
  fmt.Println(q.length())
}
