package main

import (
  "fmt"
  "math"
)

// Point 型
// ここで定義したメソッドを全て実装されている構造体ならこの型として扱える
// diatance0 をオーバーライドしてればPoint型として扱えるようになる
type Point interface {
  //レシーバーの型を書いてはいけない
  distance0() float64 // 原点からの距離を求める
}

// 二次元
type Point2d struct {
  x, y float64
}

// Point2d の生成
func newPoint2d(x, y float64) *Point2d {
  p := new(Point2d)
  p.x, p.y = x, y
  return p
}

// メソッド
func (p *Point2d) distance0() float64 {
  return math.Sqrt(p.x * p.x + p.y * p.y)
}

// 三次元
type Point3d struct {
  x, y, z float64
}

// Point3d の生成
func newPoint3d(x, y, z float64) *Point3d {
  p := new(Point3d)
  p.x, p.y, p.z = x, y, z
  return p
}

// メソッド
func (p *Point3d) distance0() float64 {
  return math.Sqrt(p.x * p.x + p.y * p.y + p.z * p.z)
}

// 合計値を求める
func sumOfDistance0(ary []Point) float64 {
  sum := 0.0
  for _, p := range ary {
    sum += p.distance0()
  }
  return sum
}

func main() {
  a := []Point{
    newPoint2d(0, 0), newPoint2d(10, 10),
    newPoint3d(0, 0, 0), newPoint3d(10, 10, 10),
  }
  fmt.Println(a[0].distance0())
  fmt.Println(a[1].distance0())
  fmt.Println(a[2].distance0())
  fmt.Println(a[3].distance0())
  fmt.Println(sumOfDistance0(a))
}
