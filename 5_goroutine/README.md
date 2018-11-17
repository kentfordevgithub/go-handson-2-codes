# ゴルーチン

- - -

* ゴルーチンはGo言語で並行処理を表すものです。
* 本来の手続き実行ではなく、非同期での関数の実行に使う事ができます。
* ゴルーチンの為に特別な実装は不要です。
* 通常の関数をgoステートメントで呼び出すだけで簡単にゴルーチンを利用できます。

```go
package main

import "fmt"

func main() {
	go SayHello()
}

func SayHello() {
	fmt.Println("Hello golang!")
}
```

- - -

* 無名関数も扱う事ができます

```go
package main

import "fmt"

func main() {
	go func() {
		fmt.Println("Hello golang!")
	}()
}
```

- - -

* ゴルーチンの実行によって何が起こるか見てみましょう

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello golang1!")

	go func() {
		fmt.Println("Hello golang2!")
	}()

	fmt.Println("Hello golang3!")
	time.Sleep(3 * time.Second)
}
```

```
Hello golang1!
Hello golang3!
Hello golang2!
```

- - -

* ゴルーチンとして呼び出された関数
	* 呼び出し元とは別の実行単位で動作を開始します
	* 関数の実行が終わると、そのゴルーチンも同時に終了します
* ゴルーチンを呼び出した関数
	* ゴルーチンを呼び出した後、次の処理へ移ります
	* ゴルーチンの完了は待たずに関数が終わると全てのゴルーチンを中断して終了します

* 呼び出し元と、呼び出されたゴルーチンの間で連携する為には次のチャネルを利用する事ができます。

- - -

