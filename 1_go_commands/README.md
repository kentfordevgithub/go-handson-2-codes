スライドの内容のものになる物を書いていきます。

# get

## 説明

ライブラリを```$GOPATH```配下にダウンロードし、プロジェクトから使得るようにしてくれるコマンド。

## 実践

```echo```を```go get```して、HTTPサーバを立ててみよう。
下記のコマンドを実行。

```
$ go get github.com/labstack/echo
```

```$GOPATH/src/```配下に、```github.com/labstack/echo```があれば```echo```を```go get```できています。
確認してみましょう。


# run

## 説明

プログラムを即時実行してくれるコマンド。コンパイル結果のバイナリはファイルに出力されない。

## 実践

```main.go```のコードを使って、```go get```した```echo```を使ったHTTPサーバを立ててみよう。

ソース。

```
package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Start(":80")
}
```

下記のコマンドを実行。

今回のサンプルでは、どの書き方も動作は一緒。

```
$ go run main.go
$ go run .
$ go run ./
```

# fmt

## 説明

プロジェクト内のソースファイルを整形する。パスの指定方法で、整形対象をファイル単体・ディレクトリ全体を変えることができる。

## 実践

```
# 対象のファイルを整形する
$ go fmt main.go

# カレントディレクトリ直下のファイルを整形する
$ go fmt .
$ go fmt ./

# カレントディレクトリ配下のファイル、ディレクトリを再帰的にたどり全て整形する
$ go fmt ./...
```

下記みたいに汚いソースでも、

```
package main
import (
"net/http"
"github.com/labstack/echo"
)
func main(){
e := echo.New()
e.GET("/",func(c echo.Context) error {
return c.String(http.StatusOK, "Hello, World!")
})
e.Start(":80")
}
```

下記のように整形してくれる。

```
package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Start(":80")
}
```

# test

## 説明

```_test.go```というサフィックスを持つ、ファイルに対して、テストを実行してくれる。
8_testで細かく説明するので、ここでは存在だけ。


# build

## 説明

プログラムをコンパイルし、実行ファイルを出力するコマンド。
開発している環境とは別の環境の実行ファイルでも、オプションの設定を変えればそれ専用にコンパイルできる。(Windows環境で、Linux用バイナリを作るなど。)

## 実践

```
# 対象ファイルをビルド
$ go build main.go

# 出力ファイルの名前を指定して対象ファイルをビルド
$ go build -o output main.go

# 対象ファイル名は省略可能
$ go build .
$ go build ./
```

# 補足

```go fmt```と同じように、```go run```も```go build```も対象ファイルの部分にに```./```や```./...```を使うことができるが、対象ファイルを明確に指定することが実際のプロジェクトではでは多いかと思います。
