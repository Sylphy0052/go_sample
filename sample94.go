package main

import "fmt"

// 数値を表すインタフェース
type Num interface {
  number()
}

// Intというインタフェース
type Int int

// numberメソッドを定義しているのでIntはNumとして扱える
func (n Int) number() {}

// Realというインタフェース
type Real float64

// numberメソッドを定義しているのでRealはNumとして扱える
func (n Real) number() {}

func sumOfNum(ary []Num) (Int, Real) {
  var sumi Int  = 0
  var sumr Real = 0.0
  for _, x := range ary {
    // xの型で場合分け
    switch v := x.(type) {
      case Int:  sumi += v
      case Real: sumr += v
    }
  }
  return sumi, sumr
}

func main() {
  var ary []Num = []Num{
    Int(1), Real(1.1), Int(2), Real(2.2), Int(3), Real(3.3),
  }
  a, b := sumOfNum(ary)
  fmt.Println(a, b)
}
