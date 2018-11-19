# 説明

```interface```として型が持つべき関数を定義することで、パッケージ間でやりとりする型を抽象化することができます。
構造体の型を隠蔽し、処理の抽象化をする際に使用します。

ここでは```interface```の定義方法と、基礎的な使い方を説明します。


# 実践

## 定義方法

```interface```の定義方法は下記です。構造体とています。
今回はサンプルなのでわかりやすくしていますが、下記の```Human```インターフェースは、前の章の```Person```構造体に対応する形になっています。

```
type Human interface {
	MyNameIs()
}
```

## interfaceを受け取る

引数に```interface```を受け取る関数を作ってみます。受け取った```Human```インターフェースの関数を実行します。
型が```interface```でも、引数での指定方法は他の型と変わりません。

```
func whoAreYou(h Human) {
	h.MyNameIs()
}
```


## 使い方

前の章で使用した```Person```構造体を例にします。レシーバーはMyNameIs()を引き続き持ちます。

```
type Person struct {
	Age  int
	Name string
}

func (p Person) MyNameIs() {
	fmt.Println("My name is " + p.Name)
}
```

```Person```は```Human```で定義づけされている関数を持つため、```whoAreYou(h Human)```関数に渡して、実行することができます。

```
p := Person{
    Name: "kent",
    Age:  29,
}

// 構造体PersonはインターフェースHumanのMyNameIs()を実装しているので、引数として渡すことができる
whoAreYou(p)
```

```My name is kent```が出力されます。


```Person```とは別に、```Human```の定義を満たす構造体を作ってみます。こちらはフィールドの定義が```Person```とは異なりますが、```MyNameIs()```を持っているのは同じです。

```
type Robot struct {
	Id string
}

func (r Robot) MyNameIs() {
	fmt.Println("My id is " + r.Id)
}
```

下記の様に使うことができます。

```
r := Robot{
    Id: "abcd",
}
whoAreYou(r)
```

```My id is abcd```が出力されます。


## 引数ありの関数を持つインターフェース

引数を持つインターフェースを定義して、使ってみます。

```
type Mouth interface {
    // 変数名は省略できるので、なくても大丈夫
	Bark(sound string)
}

func barkIt(m Mouth, sound string) {
	m.Bark(sound)
}
```

上記のインターフェースを満たす構造体を作成します。2つの構造体があり、それぞれ関数の処理内容が異なります。

```
type DogMouth struct {
}

type HumanMouth struct {
	Name string
}

func (m *DogMouth) Bark(sound string) {
	fmt.Println(sound + "!")
}

func (m *HumanMouth) Bark(sound string) {
	fmt.Println(m.Name + " said: " + sound)
}
```

下記のように使うことができます。

```
hm := &HumanMouth{
    Name: "yukpiz",
}
dm := &DogMouth{}
barkIt(hm, "hello.")
barkIt(dm, "Bowow")
```

実行結果は下記です。それぞれの処理内容を、構造体ごとに通過していることがわかります。

```
yukpiz said: hello.
Bowow!
```


## 返却値ありのケース

先程のインターフェースに返却値ありの関数を追加します。

```
type Mouth interface {
	Bark(sound string)
	StringBark(sound string) string
}
```

あわせて、```HumanMouth```に対応したレシーバーを実装します。

```
func (m *HumanMouth) StringBark(sound string) string {
	return m.Name + " said: " + sound
}
```

下記の用に使うことができます(結果は```Bark()```を使ったときと同じ)。

```
fmt.Println(hm.StringBark("hello."))
```
