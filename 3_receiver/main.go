package main

import "fmt"

type Person struct {
	Age  int
	Name string
}

func (p Person) MyNameIs() {
	fmt.Println("My name is " + p.Name)
}

func (p *Person) afterYear() {
	p.Age++
}

func (p Person) afterYearDummy() {
	p.Age++
}

func main() {
	p := Person{
		Age:  29,
		Name: "kent",
	}

	p.MyNameIs()

	fmt.Println(p.Age)
	p.afterYear()
	fmt.Println(p.Age)
	p.afterYearDummy()
	fmt.Println(p.Age)
}

func test1() {
	fmt.Println("test1")
}

func test2() {
	fmt.Println("test2")
}
