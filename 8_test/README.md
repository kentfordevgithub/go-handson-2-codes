Go言語は標準で組み込まれているtestingパッケージでユニットテスト、  
及びベンチマークテスト（ここでは割愛します）の機能をサポートしています。  

皆さんがGo言語で書いたソースコードに対してのテストアプローチは2つあります。  
上記のtestingパッケージを使ってテストを書く  
stretchr/testifyなどの外部パッケージを使ってテストを書く  

## 標準パッケージを使ったテスト実装  

testing.T構造体はユニットテストを実現する機能を持っています。  
以下に簡単なSum関数を実装してみました。  

```
example1/
   - example1.go
   - example1_test.go
```

```
package example1

func Sum(x int, y int) int {
  return x + y
}
```

この関数のテストは以下です。  

```
package example1_test
import (
  "testing" //testingパッケージのimport
  "../example1" //テスト対象のパッケージ(説明は後述)
)

//Testから始まる関数名である必要があります
func TestSum(t *testing.T) {
  //標準パッケージではassertがありません
  if example1.Sum(1, 2) != 3 {
    t.Error("エラーです！")
  }
}
```

テスト自体の構文はとてもシンプルです。  

ユニットテストのパッケージでテスト対象のパッケージをimportしています。  
わざわざ違うパッケージにする必要は実はなく、同じパッケージとしてユニットテストのコードを書くことができます。  
（そうする事で、ユニットテストではimportも必要ありませんし、パッケージ名を指定せずに関数呼び出しを記述できます。）  

反対に上記の例のように異なるパッケージにした場合、importが必要となり、テスト対象のパッケージで公開されている関数にしかアクセスできなくなります。  
これはテストコードと対象コードを疎結合にできるという利点があります。  
あくまでも対象のパッケージの利用側としてユニットテストを書くことができるようになります。  

さて、上記で書いたテストを実行してみましょう。  

```
$ go test .
PASS
ok   _/home/yukpiz/labo/repos/private/go-test/example1	0.001s
```

詳細なアウトプットが欲しい場合、-vオプションが利用できます。  

```
$ go test -v .
=== RUN   TestSum
--- PASS: TestSum (0.00s) PASS
ok   _/home/yukpiz/labo/repos/private/go-test/example1	0.001s
```

コードカバレッジの機能もついています、-coverオプションで利用してみましょう。  

```
$ go test -v -cover
=== RUN   TestSum
--- PASS: TestSum (0.00s)
PASS coverage: 100.0% of statements
ok   _/home/yukpiz/labo/repos/private/go-test/example1	0.001s
```

Go言語では標準でシンプルなユニットテストの機能が使える事がわかりました。  
次に、非公式ですが有志によって開発されている外部パッケージの利用を試してみましょう。  

## 外部パッケージを使ったテスト実装

testifyを使ったユニットテスト、assertの説明  
