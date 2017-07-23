package main

import(
  "fmt"
)

type Foo struct {
  a, b int
}

func (p Foo) swap() {
  p.a, p.b = p.b, p.a
  fmt.Println("Swap Foo!", p)
}

type Bar struct {
  a, b int
}

// ポインター
func (p *Bar) swap() {
  p.a, p.b = p.b, p.a
  fmt.Println("Swap Bar!", p)
}

func main() {
  a := Foo{1, 2}
  b := &Foo{3, 4}
  c := Bar{5, 6}
  d := &Bar{5, 6}
  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)
  a.swap()
  b.swap()
  c.swap()
  d.swap()
  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)
}
