package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello golang1!")

	go func() {
		fmt.Println("Hello golang2!")
	}()

	fmt.Println("Hello golang3!")
	time.Sleep(3 * time.Second)
}
