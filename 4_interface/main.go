package main

import "fmt"

// 前の章で出てきたPersonと同じ
type Person struct {
	Age  int
	Name string
}

func (p Person) MyNameIs() {
	fmt.Println("My name is " + p.Name)
}

// この章で説明するインターフェース
type Human interface {
	MyNameIs()
}

// インターフェースを実行する関数
func whoAreYou(h Human) {
	h.MyNameIs()
}

// TODO 引数ありの実装

// TODO 返却値ありの実装

// TODO 引数も、返却値もありの実装

func main() {
	p := Person{
		Name: "kent",
		Age:  29,
	}

	// 構造体PersonはインターフェースHumanのMyNameIs()を実装しているので、引数として渡すことができる
	whoAreYou(p)
}
