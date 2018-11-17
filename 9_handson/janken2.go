package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	PAA   = 1
	GUU   = 2
	CHOKI = 3
)

func main() {
	COUNT := 10000000
	rand.Seed(time.Now().UnixNano())

	result := [3]int{}
	for i := 0; i < COUNT; i++ {
		a := funcA()
		b := funcB()

		switch {
		case a == b:
			result[0]++
		case a == PAA && b == GUU, a == GUU && b == CHOKI, a == CHOKI && b == PAA:
			result[1]++
		case a == PAA && b == CHOKI, a == GUU && b == PAA, a == CHOKI && b == GUU:
			result[2]++
		}
	}

	fmt.Printf("あいこ: %d\n", result[0])
	fmt.Printf("Aが勝ち: %d\n", result[1])
	fmt.Printf("Bが勝ち: %d\n", result[2])
}

func funcA() int {
	s := []int{PAA, GUU, CHOKI, CHOKI}
	return s[rand.Intn(len(s))]
}

func funcB() int {
	s := []int{PAA, GUU, GUU, CHOKI}
	return s[rand.Intn(len(s))]
}
