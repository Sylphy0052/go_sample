package main

import "fmt"

// Foo の定義
type Foo struct {
  a, b int
}

func (p *Foo) printA() {
  fmt.Println(p.a)
}

func (p *Foo) printB() {
  fmt.Println(p.b)
}

// Bar の定義
type Bar struct {
  a, c int
}

func (p *Bar) printA() {
  fmt.Println(p.a)
}

func (p *Bar) printC() {
  fmt.Println(p.c)
}

// Baz の定義
type Baz struct {
  Foo
  Bar
}

func main() {
  x := new(Baz)
  x.Bar.a = 10 // x.aはコンパイルエラー
  x.b = 20
  x.c = 30
  x.Foo.printA() // x.printA()はコンパイルエラー
  x.printB()
  x.printC()
}
