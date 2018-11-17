package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println("------")
	fmtCode()
	fmt.Println("------")
	ioCode()
	fmt.Println("------")
	logCode()
}

// fmtの説明に使用するコード
func fmtCode() {
	fmt.Println("入力してください")

	var s string
	fmt.Scan(&s)

	fmt.Printf("入力された値は %s です", s)
	fmt.Println()
}

// ioの説明に使用するコード
func ioCode() {
	var reader io.Reader = os.Stdin
	var writer io.Writer = os.Stdout
	for i := 0; i < 5; i++ {
		fmt.Println("入力してください")
		var s string
		fmt.Fscan(reader, &s)
		fmt.Fprintf(writer, "入力された文字=%s\n", s)
	}
}

// logの説明に使用するコード
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
