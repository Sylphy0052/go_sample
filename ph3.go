//
// ph3.go : 哲学者の食事
//
//          Copyright (C) 2014 Makoto Hiroi
//
package main

import (
	"fmt"
	"time"
)

// 定数
const (
	GET = 0
	RET = 1
)

// フォークのリクエスト
type Req struct {
	req, fork int
	reply     chan bool
}

// リクエストの生成
func newReq(req, fork int, reply chan bool) *Req {
	p := new(Req)
	p.req = req
	p.fork = fork
	p.reply = reply
	return p
}

// フォークの管理
func forks(n int, ch chan *Req) {
	forkTbl := make([]bool, n)
	for i := 0; i < n; i++ {
		forkTbl[i] = true
	}
	for {
		r := <-ch
		switch r.req {
		case GET:
			if forkTbl[r.fork] {
				forkTbl[r.fork] = false
				r.reply <- true
			} else {
				r.reply <- false
			}
		case RET:
			forkTbl[r.fork] = true
			r.reply <- true
		}
	}
}

// フォークの取得
func getFork(fork int, out chan *Req, in chan bool) int {
	r := newReq(GET, fork, in)
	for {
		out <- r
		if <-in {
			time.Sleep(100 * time.Millisecond)
			return fork
		} else {
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// フォークの返却
func retFork(fork int, out chan *Req, in chan bool) bool {
	time.Sleep(100 * time.Millisecond)
	out <- newReq(RET, fork, in)
	return <-in
}

// 哲学者の動作
func person(m, forkR, forkL int, out chan *Req, quit chan bool) {
	in := make(chan bool)
	for n := 2; n > 0; n-- {
		fmt.Printf("Philosopher%d is thinking\n", m)
		time.Sleep(1000 * time.Millisecond)
		if m%2 == 0 {
			getFork(forkR, out, in)
			getFork(forkL, out, in)
		} else {
			getFork(forkL, out, in)
			getFork(forkR, out, in)
		}
		fmt.Printf("Philosopher%d is eating\n", m)
		time.Sleep(500 * time.Millisecond)
		retFork(forkR, out, in)
		retFork(forkL, out, in)
	}
	fmt.Printf("Philosopher%d is sleeping\n", m)
	quit <- true
}

// 実行
func main() {
	ch := make(chan *Req)
	quit := make(chan bool)
	go forks(5, ch)
	go person(1, 0, 1, ch, quit)
	go person(2, 1, 2, ch, quit)
	go person(3, 2, 3, ch, quit)
	go person(4, 3, 4, ch, quit)
	go person(5, 4, 0, ch, quit)
	for n := 5; n > 0; n-- {
		<-quit
	}
}
