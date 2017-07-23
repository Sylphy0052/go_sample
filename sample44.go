// その中でもクイックソート (quick sort) は高速なソートアルゴリズムとして有名です。
// クイックソートはある値を基準にして、要素をそれより大きいものと小さいものの
// 2つに分割していくことでソートを行います。
// 2つに分けた各々のグループを同様に分割して2つのグループに分けます。
// 最後はグループの要素が一つになってソートが完了します。

package main

import "fmt"

func quicksort(a, b int, ary []int) {
	mid := ary[a+(b-a)/2]
	i, j := a, b
	for {
		for mid > ary[i] {
			i++
		}
		for mid < ary[j] {
			j--
		}
		if i >= j {
			break
		}
		ary[i], ary[j] = ary[j], ary[i]
		i++
		j--
		fmt.Println(ary)
	}
	if a < i-1 {
		quicksort(a, i-1, ary)
	}
	if j+1 < b {
		quicksort(j+1, b, ary)
	}
}

func main() {
	a := []int{5, 6, 4, 7, 3, 8, 2, 9, 1, 0}
	fmt.Println(a)
	quicksort(0, 9, a)
}
