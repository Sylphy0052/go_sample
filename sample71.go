package main

import(
  "fmt"
  "math"
)

// 英大文字から始まる名前(関数名、型名、フィールド名など)は,
// 他のパッケージからアクセスすることができる．
// 英小文字で始まる名前は，他のパッケージからアクセスすることができない．
// 同じパッケージ内であれば英小文字から始まる名前でもアクセスすることができる．
type Point struct {
  // x float64
  // y float64
  x, y float64
}

// 構造体を引数で渡すとコピーされる
func distance(p, q Point) float64 {
  dx := p.x - q.x
  dy := p.y - q.y
  return math.Sqrt(dx * dx + dy * dy)
}

func check(p Point) float64 {
  p.x += 10
  return p.x
}

func newPoint(x, y float64) *Point {
  p := new(Point)
  p.x, p.y = x, y
  return p
}

func main() {
  // 初期値は0
  // var p Point                          // p := Point{}
  p := Point{}
  var q Point = Point{10, 10}          // q := Point{10, 10}
  var r Point = Point{x:100, y:100}    // r := Point{x:100, y:100}
  fmt.Println(p)
  fmt.Println(q)
  fmt.Println(r)
  fmt.Println(p.x)
  fmt.Println(p.y)
  fmt.Println(q.x)
  fmt.Println(q.y)
  fmt.Println(r.x)
  fmt.Println(r.y)
  fmt.Println(distance(p, q))
  fmt.Println(distance(p, r))
  fmt.Println(distance(q, r))
  fmt.Println(p.x) // 0
  fmt.Println(check(p)) // 10
  fmt.Println(p.x) // 0
}
