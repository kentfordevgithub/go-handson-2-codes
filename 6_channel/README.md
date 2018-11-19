# チャネル

* チャネルはゴルーチン間のメッセージのやり取りをするものです
* メッセージは型を指定できます
* チャネルは**送信**と**受信**の操作を持っています

```go
package main

import (
	"fmt"
)

func main() {
	ch := make(chan int) //チャネルの初期化

	go func() {
		ch <- 99 //送信文
	}()

	v := <-ch //代入による受信式
	fmt.Println(v)
}
```
- - -

* チャネルはキャパシティ(容量)があります
* チャネルは実体に容量を持っていて、容量が足りない場合、送信側は容量が空くまでまたなければいけません

```go
package main

func main() {
	ch1 := make(chan int) //容量0のint型のチャネル
	ch2 := make(chan int, 5) //容量5のint型のチャネル
}
```
















* 前章のゴルーチンだけの実装だと、main関数はゴルーチンの実装を待たずに終了してしまう為、Sleepをしていました
* チャネルを使うと、ゴルーチンからのメッセージの受信待ちをする事ができます

```go
package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	go func() {
		ch <- "yukpiz"
	}()

	go func() {
		ch <- "kent"
	}()

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
```

- - -
