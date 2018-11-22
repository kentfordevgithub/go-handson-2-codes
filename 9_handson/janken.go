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
	COUNT := 1000000
	rand.Seed(time.Now().UnixNano())

	result := [3]int{}
	for i := 0; i < COUNT; i++ {
		a := yukpiz()
		b := kent()

		switch {
		case a == b:
			result[0]++
		case a == PAA && b == GUU, a == GUU && b == CHOKI, a == CHOKI && b == PAA:
			result[1]++
		case a == PAA && b == CHOKI, a == GUU && b == PAA, a == CHOKI && b == GUU:
			result[2]++
		default:
			panic("(´・ω・｀)")
		}
	}

	fmt.Printf("引分: %d\n", result[0])
	fmt.Printf("勝利: %d\n", result[1])
	fmt.Printf("敗北: %d\n", result[2])
}

func yukpiz() int {
	s := []int{PAA, GUU, CHOKI, CHOKI}
	return s[rand.Intn(len(s))]
}

func kent() int {
	s := []int{PAA, GUU, GUU, CHOKI}
	return s[rand.Intn(len(s))]
}
