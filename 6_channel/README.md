# チャネル

* チャネルはゴルーチン間のメッセージのやり取りをするものです
* メッセージは型を指定できます

```go
package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		fmt.Println("goroutine!")
		c <- 99
	}()

	fmt.Println(<-c)
}
```

- - -

* 前章のゴルーチンだけの実装だと、main関数はゴルーチンの実装を待たずに終了してしまう為、Sleepをしていました。
* チャネルを使うと、ゴルーチンからのメッセージの受信待ちをする事ができます。
