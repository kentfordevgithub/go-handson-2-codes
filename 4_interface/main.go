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

type Robot struct {
	Id string
}

func (r Robot) MyNameIs() {
	fmt.Println("My id is " + r.Id)
}

// この章で説明するインターフェース
type Human interface {
	MyNameIs()
}

// インターフェースを実行する関数
func whoAreYou(h Human) {
	h.MyNameIs()
}

type Mouth interface {
	// 変数名は省略できるので、なくても大丈夫
	Bark(sound string)
	StringBark(sound string) string
}

type DogMouth struct {
}

type HumanMouth struct {
	Name string
}

func (m *DogMouth) Bark(sound string) {
	fmt.Println(sound + "!")
}

func (m *DogMouth) StringBark(sound string) string {
	return sound + "!"
}

func (m *HumanMouth) Bark(sound string) {
	fmt.Println(m.Name + " said: " + sound)
}

func (m *HumanMouth) StringBark(sound string) string {
	return m.Name + " said: " + sound
}

func barkIt(m Mouth, sound string) {
	m.Bark(sound)
}

func main() {
	p := &Person{
		Name: "kent",
		Age:  29,
	}

	// 構造体PersonはインターフェースHumanのMyNameIs()を実装しているので、引数として渡すことができる
	whoAreYou(p)

	r := &Robot{
		Id: "abcd",
	}
	whoAreYou(r)

	hm := &HumanMouth{
		Name: "yukpiz",
	}
	dm := &DogMouth{}
	barkIt(hm, "hello.")
	barkIt(dm, "Bowow")
	fmt.Println(hm.StringBark("hello."))
	fmt.Println(dm.StringBark("hello."))
}
