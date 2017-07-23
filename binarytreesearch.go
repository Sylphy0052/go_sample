package main

import(
  "fmt"
)

// 節を表すインタフェース
type Item interface {
  Eq(Item) bool // x.Eq(y) -> xとyが同じならtrueを返す
  Less(Item) bool // x.Less(y) -> x < yならtrueを返す.ポリモーフィズムと型アサーションが働くので実行時間は遅い…
}

// 節
type Node struct {
  item Item
  left, right *Node // 節を表すNodeが2つ必要
}

// 節の作成
func newNode(x Item) *Node {
  p := new(Node)
  p.item = x
  return p
}

// データの探索
func searchNode(node *Node, x Item) bool {
  for node != nil {
    switch {
    case x.Eq(node.item): return true // 一致したらtrueを返す
    case x.Less(node.item): node = node.left // 小さかったら左の節を調べる
    default: node = node.right // 大きかったら右の節を調べる
    }
  }
  return false // 最後まで一致しなかったらfalseを返す
}

// データの挿入
func insertNode(node *Node, x Item) *Node {
  switch {
  case node == nil: return newNode(x) // 子が存在しないならそこに新たなNodeを作成する
  case x.Eq(node.item): return node // 同じ値があればそれを返す
  case x.Less(node.item): node.left = insertNode(node.left, x) // 小さかったら左の節でinsertを試す
  default: node.right = insertNode(node.right, x) // 大きかったら右の節でinsertを試す
  }
  return node
}

// データの削除 -> 葉を削除する場合と木の途中を削除する場合がある

// 最小値を探す関数
func searchMin(node *Node) Item {
  if node.left == nil { // もし左の子が存在しないならそれが最小値
    return node.item
  }
  return searchMin(node.left) // そうでないなら左の子で試す
}

// 最小値を削除する
func deleteMin(node *Node) *Node {
  if node.left == nil {
    return node.right // なんで？
  }
  node.left = deleteMin(node.left)
  return node
}

// データの削除
func deleteNode(node *Node, x Item) *Node { // xは削除するデータ
  if node != nil { // nilなら木は空なので何もしない
    if x.Eq(node.item) { // 等しいならその節を削除
      if node.left == nil { // 左の節が存在しないなら右の節を返す
        return node.right
      } else if node.right == nil { // 右の節が存在しないなら左の節を返す
        return node.left
      } else { // 子が2つあれば右の部分木の最小値を書き換え，最小値を削除する
        node.item = searchMin(node.right)
        node.right = deleteMin(node.right)
      }
    } else if x.Less(node.item) {
      node.left = deleteNode(node.left, x)
    } else {
      node.right = deleteNode(node.right, x)
    }
  }
  return node
}

// 巡回
func foreachNode(f func(Item), node *Node) {
  if node != nil {
    foreachNode(f, node.left)
    f(node.item)
    foreachNode(f, node.right)
  }
}

// 木の構造体
type Tree struct {
  root *Node
}

// 二分木の生成
func newTree() *Tree {
  return new(Tree)
}

// データの探索
func (t *Tree) searchTree(x Item) bool {
  return searchNode(t.root, x)
}

// データの挿入
func (t *Tree) insertTree(x Item) {
  t.root = insertNode(t.root, x)
}

// データの削除
func (t *Tree) deleteTree(x Item) {
  t.root = deleteNode(t.root, x)
}

// 二分木の巡回
func (t *Tree) foreachTree(f func(Item)) {
  foreachNode(f, t.root)
}

// 表示
func (t *Tree) printTree() {
  t.foreachTree(func(x Item) { fmt.Print(x, " ") })
  fmt.Println()
}

// インタフェースの実装
type Int int

func (n Int) Eq(m Item) bool {
  return n == m.(Int)
}

func (n Int) Less(m Item) bool {
  return n < m.(Int)
}

func main() {
  a := newTree()
  for _, v := range []int{5,6,4,7,3,8,2,9,1,0} {
    a.insertTree(Int(v))
  }
  a.printTree()
  for i := 0; i < 10; i++ {
    fmt.Println(a.searchTree(Int(i)))
  }
  for i := 0; i < 10; i++ {
    a.deleteTree(Int(i))
    a.printTree()
  }
}
