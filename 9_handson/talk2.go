package main

import (
	"fmt"
)

type Worker struct {
	Name          string
	BroadcastChan chan int
	DoneChan      chan int
	ReceiveChan   chan int
	Receive       func(*Worker)
}

func main() {
	bch := make(chan int)  //ブロードキャスト用のチャネル
	done := make(chan int) //終了判定用のチャネル

	var workers []*Worker
	workers = append(workers, &Worker{
		Name:          "yukpiz",
		BroadcastChan: bch,
		ReceiveChan:   make(chan int),
		Receive:       yukpiz,
	})
	workers = append(workers, &Worker{
		Name:          "kent",
		BroadcastChan: bch,
		ReceiveChan:   make(chan int),
		Receive:       kent,
	})
	workers = append(workers, &Worker{
		Name:          "ariaki",
		BroadcastChan: bch,
		ReceiveChan:   make(chan int),
		Receive:       ariaki,
	})

	go BroadcastObserve(workers, bch, done)
	for _, w := range workers {
		go w.Receive(w)
	}

	// 全てのチャネルが閉じられている事を確認
	<-done
	<-bch
	for _, w := range workers {
		<-w.ReceiveChan
	}
	fmt.Println("exit!")
}

func BroadcastObserve(workers []*Worker, bch chan int, done chan int) {
	defer close(done)
	for v := range bch {
		for _, w := range workers {
			w.ReceiveChan <- v
		}
	}

	fmt.Println("closed broad cast")
	for _, w := range workers {
		close(w.ReceiveChan)
	}
}

func yukpiz(me *Worker) {
	fmt.Printf("%s: 皆さんの名前を教えてください\n", me.Name)
	me.BroadcastChan <- 1
	for v := range me.ReceiveChan {
		switch v {
		case 3:
			fmt.Printf("%s: ありがとうございます\n", me.Name)
			fmt.Printf("%s: 好きなモノを教えてください\n", me.Name)
			v++
			me.BroadcastChan <- v
		case 5:
			fmt.Printf("%s: 日本酒が好きなんですね\n", me.Name)
			v++
			me.BroadcastChan <- v
		case 7:
			fmt.Printf("%s: 生ハムを食べにいきましょう\n", me.Name)
			v++
			me.BroadcastChan <- v

		}
	}
}

func kent(me *Worker) {
	for v := range me.ReceiveChan {
		switch v {
		case 1:
			fmt.Printf("%s: 私はけんとです\n", me.Name)
			v++
			me.BroadcastChan <- v
		case 6:
			fmt.Printf("%s: 私は生ハムが好きです\n", me.Name)
			v++
			me.BroadcastChan <- v
		case 8:
			fmt.Printf("%s: 行きましょう\n", me.Name)
			close(me.BroadcastChan)
		}
	}
}

func ariaki(me *Worker) {
	for v := range me.ReceiveChan {
		switch v {
		case 2:
			fmt.Printf("%s: 私はありあきです\n", me.Name)
			v++
			me.BroadcastChan <- v
		case 4:
			fmt.Printf("%s: 私は日本酒です\n", me.Name)
			v++
			me.BroadcastChan <- v
		}
	}
}
