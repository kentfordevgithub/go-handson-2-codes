package main

import (
	"fmt"
	"github.com/rickb777/date"
	"time"
)

func main() {
	date := date.New(2018, time.November, 27)
	fmt.Println(date)
}
