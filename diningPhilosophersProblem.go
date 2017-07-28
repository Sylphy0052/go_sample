// Dining Philosophers Problem

package main

import (
	"fmt"
	"time"
)

// 定数定義
const (
	GET   = 0 // フォークの取得
	RET   = 1 // フォークの返却
	LEFT  = 2
	RIGHT = 3
)

// フォークのリクエスト
type Req struct {
	req, fork, side int
	reply           chan bool
}

// リクエストの生成
func newReq(req, fork, side int, reply chan bool) *Req {
	p := new(Req)
	p.req = req
	p.fork = fork
	p.side = side
	p.reply = reply
	return p
}

// フォークの管理
func forks(n int, ch chan *Req) { // n : フォークの本数
	forkTbl := make([]bool, n) // フォークの有無
	for i := 0; i < n; i++ {
		forkTbl[i] = true // 初期値をtrueに変更
	}
	for {
		r := <-ch
		switch r.req {
		case GET:
			if forkTbl[r.fork] {
				if n == 1 && r.side == RIGHT {
					r.reply <- false
				} else {
					forkTbl[r.fork] = false
					n--
					r.reply <- true
				}
			} else {
				r.reply <- false
			}
		case RET:
			forkTbl[r.fork] = true
			n++
			r.reply <- true
		}
	}
}

// フォークの取得
// フォークをサーバに要求
func getFork(fork, side int, out chan *Req, in chan bool) bool {
	r := newReq(GET, fork, side, in)
	for {
		out <- r
		if <-in {
			time.Sleep(100 * time.Millisecond)
			return true
		} else if side == LEFT {
			return false
		} else {
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// デッドロックの防止
// 左のフォークを要求する
// func getFork1(fork int, out chan *Req, in chan bool) bool {
// 	out <- newReq(GET, fork, in)
// 	return <-in
// }

// フォークの返却
func retFork(fork, side int, out chan *Req, in chan bool) bool {
	time.Sleep(100 * time.Millisecond)
	out <- newReq(RET, side, fork, in)
	return <-in
}

// 哲学者の動作
// m : 哲学者の番号 forkR : 右のフォーク forkL : 左のフォーク
func person(m, forkR, forkL int, out chan *Req, quit chan bool) {
	in := make(chan bool) // リクエストの返答を受ける
	for n := 2; n > 0; {
		fmt.Printf("Philosopher%d is thinking\n", m)
		time.Sleep(1000 * time.Millisecond)
		getFork(forkR, RIGHT, out, in)
		// getFork(forkL, out, in)
		if getFork(forkL, LEFT, out, in) {
			fmt.Printf("Philosopher%d is eating\n", m)
			time.Sleep(500 * time.Microsecond)
			retFork(forkR, RIGHT, out, in)
			retFork(forkL, LEFT, out, in)
			n--
		} else {
			retFork(forkR, RIGHT, out, in)
		}
	}
	fmt.Printf("Philosopher%d is sleeping\n", m)
	quit <- true
}

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
