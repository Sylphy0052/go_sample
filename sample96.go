package main

import "fmt"

// Foo
type Foo struct {
  a int
}

type FooI interface {
  getA() int
}

func (p *Foo) getA() int { return p.a }

// Bar
type Bar struct {
  b int
}

type BarI interface {
  getB() int
}

func (p *Bar) getB() int { return p.b }

// Baz
type Baz struct {
  Foo
  Bar
}

// InterfaceにInterfaceを埋め込む
type BazI interface {
  FooI
  BarI
}

func main() {
  a := []FooI{
    &Foo{1}, &Foo{2}, &Baz{},
  }
  b := []BarI{
    &Bar{10}, &Bar{20}, &Baz{},
  }
  c := []BazI{
    &Baz{}, &Baz{Foo{1}, Bar{2}}, &Baz{Foo{3}, Bar{4}},
  }
  for i := 0; i < 3; i++ {
    fmt.Println(a[i].getA())
    fmt.Println(b[i].getB())
    fmt.Println(c[i].getA())
    fmt.Println(c[i].getB())
  }
}
