package main

import (
  "fmt"
)

type Foo struct {
  a, b int
}

func (x Foo) printA() {
  fmt.Println("a = ", x.a)
}

func (x Foo) printB() {
  fmt.Println("b = ", x.b)
}

type Bar struct {
  foo Foo
  c int
}

func (x Bar) printC() {
  fmt.Println("c = ", x.c)
}

// 構造体の埋め込み ≒ inheritance(継承)
// 型を引き継げない (Foo型にBaz型を代入できない)
// Java List = ArrayList(アップキャスト)みたいな
// Goでやるならインタフェースを使う
type Baz struct {
  Foo
  c int
}

// オーバーライド
func (x Baz) printB() {
  fmt.Print("Baz : ")
  x.Foo.printB()
}

func (x Baz) printC() {
  fmt.Println("c = ", x.c)
}

func main() {
  x := Foo{1, 2}
  y := Bar{Foo{10, 20}, 30}
  z := Baz{Foo{100, 200}, 300}
  x.printA()
  x.printB()
  y.foo.printA()
  y.foo.printB()
  y.printC()
  z.printA()
  z.printB()
  z.printC()
}
