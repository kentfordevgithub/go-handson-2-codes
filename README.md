# go-handson-2-codes
Go言語ハンズオン2回目のための、説明部分のサンプルコードを管理する。

# GOPATHの設定
各々でお好きな場所で。$HOME/go配下から変更しないなら不要。

## コマンドライン

```
export GOPATH=$PWD
export PATH=$PATH:$GOPATH/bin
```

## VSCode
コマンドラインでGOPATHを変更しても、VSCode上への反映はされないので、settings.jsonに描きを追記して、VSCode上でのGOPATHを設定する。

```
{
"go.gopath": "環境変数GOAPTHを同じパス",
}
```
