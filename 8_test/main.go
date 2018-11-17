package main

import "fmt"

type Profile struct {
	ID   int32
	Name string
}

func main() {
	p1 := &Profile{}

	p1.ChangeName1()
	fmt.Println(p1)

	p1.ChangeName2()
	fmt.Println(p1)
}

func (p *Profile) ChangeName1() {
	p.Name = "yukpiz"
}

func (p Profile) ChangeName2() {
	p.Name = "kent"
}
