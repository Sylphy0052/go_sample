package tools

// xs と ys は等しいか
func Equal(xs, ys []int) bool {
	if len(xs) != len(ys) {
		return false
	}
	for i := 0; i < len(xs); i++ {
		if xs[i] != ys[i] {
			return false
		}
	}
	return true
}

// xs を複製する
func Dup(xs []int) []int {
	ys := make([]int, len(xs))
	copy(ys, xs)
	return ys
}

// n は xs に含まれるか
func Member(n int, xs []int) bool {
	for _, x := range xs {
		if n == x {
			return true
		}
	}
	return false
}

// n と等しい要素の位置を求める
func Position(n int, xs []int) int {
	for i, x := range xs {
		if n == x {
			return i
		}
	}
	return -1
}

// n と等しい要素を数える
func Count(n int, xs []int) int {
	c := 0
	for _, x := range xs {
		if n == x {
			c++
		}
	}
	return c
}

// マッピング
func Map(f func(int) int, xs []int) []int {
	ys := make([]int, len(xs))
	for i, x := range xs {
		ys[i] = f(x)
	}
	return ys
}

// フィルター
func Filter(f func(int) bool, xs []int) []int {
	ys := make([]int, 0, len(xs))
	for _, x := range xs {
		if f(x) {
			ys = append(ys, x)
		}
	}
	return ys
}

// 畳み込み
func Fold(f func(int, int) int, a int, xs []int) int {
	for _, x := range xs {
		a = f(a, x)
	}
	return a
}

// n と等しい要素を取り除く
func Remove(n int, xs []int) []int {
	return Filter(func(x int) bool { return x != n }, xs)
}

// 順列の生成
func permSub(f func([]int), n int, xs, ys []int) {
	if n == 0 {
		f(ys)
	} else {
		for _, x := range xs {
			permSub(f, n-1, Remove(x, xs), append(ys, x))
		}
	}
}

func Permutation(f func([]int), n int, xs []int) {
	permSub(f, n, xs, make([]int, 0, n))
}

func PermGen(n int, xs []int) <-chan []int {
	ch := make(chan []int)
	go func() {
		Permutation(func(ys []int) { ch <- Dup(ys) }, n, xs)
		close(ch)
	}()
	return ch
}

// 組み合わせの生成
func combSub(f func([]int), n int, xs, ys []int) {
	switch {
	case n == 0:
		f(ys)
	case len(xs) == n:
		f(append(ys, xs...))
	default:
		combSub(f, n-1, xs[1:], append(ys, xs[0]))
		combSub(f, n, xs[1:], ys)
	}
}

func Combination(f func([]int), n int, xs []int) {
	combSub(f, n, xs, make([]int, 0, n))
}

func CombGen(n int, xs []int) <-chan []int {
	ch := make(chan []int)
	go func() {
		Combination(func(ys []int) { ch <- Dup(ys) }, n, xs)
		close(ch)
	}()
	return ch
}
