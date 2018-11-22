package main

import (
	"fmt"
)

func main() {
	ych := make(chan int)
	kch := make(chan int)
	done := make(chan int)

	go yukpiz(ych, kch)
	go kent(ych, kch, done)

	<-done
	fmt.Println("exit!")
}

func yukpiz(ych chan int, kch chan int) {
	name := "yukpiz"

	fmt.Printf("%s: こんにちは、名前を教えてください\n", name)
	kch <- 1

	for v := range ych {
		switch v {
		case 1:
			fmt.Printf("%s: ありがとうございます\n", name)
			fmt.Printf("%s: 好きな食べ物はなんですか？\n", name)
			kch <- 2
		case 2:
			fmt.Printf("%s: いいですね、生ハムを食べに行きましょう\n", name)
			kch <- 3
		}
	}
}

func kent(ych chan int, kch chan int, done chan int) {
	name := "kent"
	for v := range kch {
		switch v {
		case 1:
			fmt.Printf("%s: はじめまして、私はkentです\n", name)
			ych <- 1
		case 2:
			fmt.Printf("%s: 生ハムが好きです\n", name)
			ych <- 2
		case 3:
			fmt.Printf("%s: 是非、行きましょう\n", name)
			close(done)
		}
	}
}
