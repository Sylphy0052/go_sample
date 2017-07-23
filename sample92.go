package main

import (
  "fmt"
  "math"
)

// 図形のインタフェース
type Figure interface {
  kindOf() string // 図形の種類を文字列で返す関数
  area()   float64 // 面積を求める関数
  print()         // 種類と面積を表示する関数
}

// 三角形
type Triangle struct {
  altitude, base float64
}

func newTriangle(a, b float64) *Triangle {
  p := new(Triangle)
  p.altitude, p.base = a, b
  return p
}

func (_ *Triangle) kindOf() string { return "Triangle" }

func (p *Triangle) area() float64 {
  return p.altitude * p.base / 2
}

func (p *Triangle) print() {
  fmt.Println("Triangle : area =", p.area())
}

// 四角形
type Rectangle struct {
  width, height float64
}

func newRectangle(w, h float64) *Rectangle {
  p := new(Rectangle)
  p.width, p.height = w, h
  return p
}

func (_ *Rectangle) kindOf() string { return "Rectangle" }

func (p *Rectangle) area() float64 {
  return p.width * p.height
}

func (p *Rectangle) print() {
  fmt.Println("Rectangle: area =", p.area())
}

// 円
type Circle struct {
  radius float64
}

func newCircle(r float64) *Circle {
  p := new(Circle)
  p.radius = r
  return p
}

func (_ *Circle) kindOf() string { return "Circle" }

func (p *Circle) area() float64 {
  return p.radius * p.radius * math.Pi
}

func (p *Circle) print() {
  fmt.Println("Circle: area =", p.area())
}

// 面積の合計値を求める
func sumOfArea(a []Figure) float64 {
  sum := 0.0
  for _, fig := range a {
    sum += fig.area()
  }
  return sum
}

func main() {
  var a Figure = newTriangle(10, 10)
  fmt.Println(a.kindOf())
  fmt.Println(a.area())
  a.print()
  a = newRectangle(10, 10)
  fmt.Println(a.kindOf())
  fmt.Println(a.area())
  a.print()
  a = newCircle(10)
  fmt.Println(a.kindOf())
  fmt.Println(a.area())
  a.print()

  var b []Figure = []Figure{
    newTriangle(100, 100), newRectangle(100, 100), newCircle(100),
  }
  fmt.Println(sumOfArea(b))
}
