package main

import (
	"log"
	"time"
)

func main() {
	log.Print("started.")

	//チャネル
	finished := make(chan bool)

	funcs := []func(){
		func() {
			log.Print("sleep1 started.")
			time.Sleep(1 * time.Second)
			log.Print("sleep1 finished.")
			finished <- true
		},
		func() {
			log.Print("sleep2 started.")
			time.Sleep(2 * time.Second)
			log.Print("sleep2 finished.")
			finished <- true
		},
		func() {
			log.Print("sleep3 started.")
			time.Sleep(3 * time.Second)
			log.Print("sleep3 finished.")
			finished <- true
		},
	}

	//並列化
	for _, sleep := range funcs {
		go sleep()
	}

	//終わるまで待つ
	for i := 0; i < len(funcs); i++ {
		<-finished
	}

	log.Print("all finished.")
}
