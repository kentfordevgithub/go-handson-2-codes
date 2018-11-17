package main

import (
	"fmt"
	"time"

	"github.com/rickb777/date"
)

func main() {
	date := date.New(2018, time.November, 27)
	fmt.Println(date)
}
