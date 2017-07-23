package main

import "fmt"

type Point struct {
  x, y float64
}

func newPoint(x, y float64) *Point {
  p := new(Point)
  p.x, p.y = x, y
  return p
}

func main () {
  var a []Point = []Point{
    {x:0, y:0}, {10, 10}, {100, 100},
  }
  var b []*Point = make([]*Point, 8)
  fmt.Println(a) // [{0 0} {10 10} {100 100}]
  fmt.Println(b)
  for i := 0; i < 8; i++ {
    b[i] = newPoint(float64(i), float64(i))
  }
  fmt.Println(b)
  for i := 0; i < 8; i++ {
    fmt.Println(*b[i])
  }
}
