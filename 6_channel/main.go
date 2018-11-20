package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	go func() {
		ch <- "yukpiz"
	}()

	go func() {
		ch <- "kent"
	}()

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
