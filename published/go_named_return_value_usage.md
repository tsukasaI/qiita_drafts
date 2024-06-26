# GoのNamed Return Valueとそのうまい使い方

Goには関数/メソッドを定義する時に返り値を名前をつけて関数の内部で初期化することなく変数を作ることができる仕組みがある。

これをNamed Return Valueと言う。

本記事ではNamed Return Valueのメリットデメリットと使い所を解説していく。

## Named Return Value

こちらにかかれている。

https://go.dev/doc/effective_go#named-results

参考のコードは以下のようになる。

```go
func sampleNamedReturnValue() (value int) {
    value = 2
    return
}
```

この返り値valueはint型で扱われ、ゼロ値（intでは0）で初期化される。

また関数内ではreturnだけ書くことでvalueにセットされている変数を返すことができる。

### メリット

個人的には2つあると考えている。

- 命名から返り値が何かが予測しやすい
- 少し短く書ける

まず命名については言うまでもなく、返す値がどのような役目をするのかが型情報に加えて補強されるため慣れている人にとっては可読性が向上する。

短く書けることは変数をあらかじめ準備されていることから変数宣言が不要になります。

### デメリット

こちらは一つで

- 変数のスコープが広い

必要以上にスコープを広げることは良しとされていないが、Named Return Valueでは変数が関数内全体で有効になる。

変数のスコープを狭めたい場合には使えないことになります。

## 使い所

個人的には返り値の情報を付け足してあげるのが良いと考える。

Goはご存知の通り複数の返り値を持たせることができるため、それぞれが何を示すのかを使うとよい。

例えばゲームを作るとして得点をまとめて取得するような関数があるとする。

あるユーザーの今回の得点、ユーザーの最高得点、全ユーザーの最高得点を返したいときには以下のようにすると読みやすくなる。

```go
func GetScores(userID string) (currentScore, userHighScore, highScore int) {
    // 具体的な処理
    return
}
```
