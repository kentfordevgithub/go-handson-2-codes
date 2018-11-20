# チャネル

* チャネルはゴルーチン間のメッセージのやり取りをするものです
* メッセージは型を指定できます

* チャネルは以下の特性を持っています
* キャパシティ
* 受信
* 送信
* クローズ

```go
package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 5) //容量5のチャネルを初期化

	go func() {
		ch <- 99 //送信文
		close(ch) //チャネルを閉じる
	}()

	v, closed := <-ch // 受信の2番目の返り値でcloseかどうかの判定を受け取れます

	fmt.Println(v) // 99
	fmt.Println(closed) // true
}
```
- - -

* チャネルにはキャパシティ(容量)があります
* チャネルのキャパシティは送信されたデータをバッファできるサイズです
* 受信側が待機している場合は、バッファは使用されません
* バッファの容量が足りなくなると、チャネルへの送信はブロックされます

```go
ch := make(chan int) //容量0のチャネル

go func() {
	// 受信がないので、送信側はバッファに入れようとしますが、
	// 容量が0なのでブロックされます
	ch <- 99
	fmt.Println("送信できない！")
}()

time.Sleep(60 * time.Second)
```

```go
ch := make(chan int) //容量0のチャネル

go func() {
	// 受信があるので、バッファは使用されずに送信できます
	ch <- 99
	fmt.Println("送信できる！")
}()

fmt.Println(<-ch)
time.Sleep(60 * time.Second)
```

- - -

* チャネルは便利ですが、使い方によっては問題を起こしてしまう事もあります
* 並列処理の完了を検知したいとき、チャネルは有効です　
* 以下の例では受信側が処理の完了を待ち続けてしまいます

```go
done := make(chan bool)

go func() {
	// ... 何かの処理

	// ここで例外などでゴルーチンを抜けてしまった場合、
	// 受信側がdoneへの送信を待ち続けてしまう
	done <- true
}()

fmt.Println(<-ch)
fmt.Println("処理完了！")
```

```go
ch := make(chan bool)

go func() {
	defer close(ch)
	// ... 何かの処理
	// 何が起きても最後にcloseされる
}()

fmt.Println(<-ch)
fmt.Println("処理完了！")
```













