# 概要

標準パッケージとはなんぞやを説明する。

ここで取り上げるパッケージ。
- fmt
- io
- log


# 標準パッケージとは

Go自身に含まれるライブラリ群のため、```go get```などせずに外部ライブラリに依存せず使える機能のこと。
入出力周りやJSONパース、サーバ起動などの機能が提供されており、これだけでもある程度の処理を記述することができる。


# fmt

## 説明

標準入力や標準出力などの機能を提供します。

## 実践

標準入力された文字列を標準出力する。

```
func fmtCode() {
	fmt.Println("入力してください")

	var s string
	fmt.Scan(&s)

	fmt.Printf("入力された値は %s です", s)
	fmt.Println()
}
```


# io

## 説明

入出力の扱う、Reader/Writerを提供し、

## 実践

先程の標準入力、標準出力のコードを、ioのReaderとWriterを使った書き方に変えてみましょう。

```
func ioCode() {
	fmt.Println("入力してください")
	var reader io.Reader = os.Stdin
	var writer io.Writer = os.Stdout
	for i := 0; i < 5; i++ {
		var s string
		fmt.Fscan(reader, &s)
		fmt.Fprintf(writer, "入力された文字=%s\n", s)
	}
}
```

Reader/Writerはインターフェース(後述)として定義されており、広く使われるため、覚えておくと汎用性のあるものです。


# log

## 説明

下記のような機能な機能を備えており、アプリケーションのロギング機能を提供します。
- 日付やファイル名などのフォーマットを指定しつつ、任意の文字列を都度表示する。
- エラーログを出力して異常終了させる。

## 実践

標準出力の対象をLoggerに変更し、標準入力の内容をログ出力しましょう。
エラーログを出力しましょう。

```
func logCode() {
	writer := os.Stdout
	var logger *(log.Logger)
	// log.LstdFlagsは標準日付、log.LshortFileはファイル名、をそれぞれログに必ずつける設定
	logger = log.New(writer, "[HANDSON_LOG]", log.LstdFlags|log.Lshortfile)

	var reader io.Reader = os.Stdin
	for i := 0; i < 5; i++ {
		fmt.Println("入力してください")
		var s string
		fmt.Fscan(reader, &s)
		logger.Printf("%s\n", s)
	}

	// 最後はエラーログを出力して、終了ステータスもエラーとする
	logger.Fatal("!!!error exit!!!")
}
```

終了時は```exit status 1```が出力される。
