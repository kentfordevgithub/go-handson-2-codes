# 説明

定義した構造体に対して関数を実装することができる。
構造体に関数をつけることで、構造体内のデータを他のプログラムから、ロジックとセットにして取り出したり、変更したりすることができる。

構造体の値自身に定義する```値レシーバー```と、構造体のポインタに対して定義する```ポインタレシーバー```の2つが使い分けとして存在する。

書き方としては、関数の```func```と```関数の名前```の間に、レシーバーを定義する、```構造体の型```と```仮引数```をつけます。関数の中身にレシーバーの中で実行したい処理を記述します。仮引数はその処理の中で使用することができます。


# 実践

**かなり流れが大事なので、細かい説明も交えながら書いていきます**

まずレシーバーを定義してみましょう。サンプルとして構造体を定義。

```
type Person struct {
	Age  int
	Name string
}
```

年齢と名前があるので、それを使った処理をレシーバーで実装してみます。

```
// 値レシーバ
func (p Person) MyNameIs() {
	fmt.Println("My name is " + p.Name)
}

// ポインタレシーバー
func (p *Person) afterYear() {
	p.Age++
}
```

実際にレシーバーを使ってみます。

```
p := &Person{
    Age:  29,
    Name: "kent",
}

p.MyNameIs() // kentが出力される

fmt.Println(p.Age) // 29が出力される
p.afterYear()
fmt.Println(p.Age) // 30が出力される
```

ポインタレシーバーを使う理由は、構造体の中の値を変更するためです。理屈としては、ポインタレシーバーで定義しないと、afterYear関数の仮引数には、Personの構造体がコピーされたものが入り、呼び出し元の構造体の中身は変更できなくなるためです。
下記のように、値レシーバーで定義してしまうと、出力結果は29のままとなってしまいます。

```
// afterYear()と同じ処理だけど、こっちは値レシーバー
func (p Person) afterYearDummy() {
	p.Age++
}
```

また、同じ名前、同じ引数で、値レシーバーかポインタレシーバーの2つを同時に定義することはできません。
どちらかの方法でレシーバーを定義すれば、呼び出し元がポインタかどうかにかかわらず、どちらでも使うことができます。

上記の例では、pはPersonのポインタ型ですが、値レシーバーのMyNameIs()を呼び出せています。