package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

type SortI interface {
	Len() int
	Less(int, int) bool
	Swap(int, int)
}

func quickSortSub(data SortI, low, high int) {
	p := low + (high-low)/2
	i, j := low, high
	for {
		for data.Less(i, p) {
			i++
		}
		for data.Less(p, j) {
			j--
		}
		if i >= j {
			break
		}
		data.Swap(i, j)
		switch {
		case p == i:
			p = j
		case p == j:
			p = i
		}
		i++
		j--
	}
	if low < i-1 {
		quickSortSub(data, low, i-1)
	}
	if high > j+1 {
		quickSortSub(data, j+1, high)
	}
}

func quickSort(data SortI) {
	quickSortSub(data, 0, data.Len()-1)
}

func quickSortParaSub(data SortI, low, high int) {
	if high-low < 1024 {
		quickSortSub(data, low, high)
		return
	}
	p := low + (high-low)/2
	i, j := low, high
	for {
		for data.Less(i, p) {
			i++
		}
		for data.Less(p, j) {
			j--
		}
		if i >= j {
			break
		}
		data.Swap(i, j)
		switch {
		case p == i:
			p = j
		case p == j:
			p = i
		}
		i++
		j--
	}
	ch := make(chan int, 2)
	go func() {
		quickSortParaSub(data, low, i-1)
		ch <- 0
	}()
	go func() {
		quickSortParaSub(data, j+1, high)
		ch <- 0
	}()
	<-ch
	<-ch
}

func quickSortPara(data SortI) {
	quickSortParaSub(data, 0, data.Len()-1)
}

type IntArray []int

func (ary IntArray) Len() int {
	return len(ary)
}

func (ary IntArray) Less(i, j int) bool {
	return ary[i] < ary[j]
}

func (ary IntArray) Swap(i, j int) {
	ary[i], ary[j] = ary[j], ary[i]
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	for i := 1; i <= 8; i *= 2 {
		a := make(IntArray, i*1000000)
		b := make(IntArray, i*1000000)
		for j := 0; j < i*1000000; j++ {
			x := rand.Int()
			a[j] = x
			b[j] = x
		}
		fmt.Println("-----", i, "-----")
		s := time.Now()
		quickSort(a)
		e := time.Now().Sub(s)
		fmt.Println(e)
		s = time.Now()
		quickSortPara(b)
		e = time.Now().Sub(s)
		fmt.Println(e)
	}
}
