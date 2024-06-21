# インターフェースについてめっちゃ感覚的にまとめる

プログラミングを学習する、実装をしているとよく聞くインターフェースという言葉があり、筆者は最初あまりピンと来なかったので、ここでインターフェースについてまとめてみます。

インターフェースとは境界、表面、といった意味があり、プログラミングの世界では、クラスやメソッドの仕様を定義するためのものとして使われます。

本記事では体系的に解説するのではなく感覚的にまとめることを目指します。

## 身近なものに例えてみる

いきなりプログラムの世界のinterfaceを説明するのは難しいので、身近なものに例えてみます。

### 自動販売機

例えば自動販売機は人間がお金を入れて、ボタンを押すと商品が出てくる動きをします。

購入の対象が飲み物でもお菓子でも、自動販売機は同じようにお金を受け取り、商品を出すという動きをします。

これをインターフェースという考え方で表現すると、「お金を受け取りボタンを押す」と「商品を出す」という動きを期待していると考えることができます。

利用する我々からすると、お金を入れてボタンを押すという動きをするだけで、商品が出てくるという動きを期待しているので、その動きをインターフェースとして考えることができます。

ここでのポイントは商品が出てくるまでの動きを知らなくても、お金を入れてボタンを押すという動きをするだけで商品が出てくればそれでいいということです。

### 洗濯機

少しくどいがもう一つ例を上げて抽象度を上げてみます。

洗濯機に対して考えてみます。

我々は洗濯機に対しては洗濯物を入れて、洗剤を入れて、ボタンを押すという動きをするときれいになった洗濯物が出てくることを期待します。

ドラム式でも縦型でも、静音機能やイオンを使っているかどうか、洗濯機の中でどのような動きをしているかに関係なく、洗濯物を入れて洗剤を入れてボタンを押すという動きをするだけで、洗濯物がきれいになるという動きを期待しているので、その動きをインターフェースとして考えることができます。

## プログラムの世界でのインターフェース

さてここまで身近なものに例えてみたので今後はプログラムの世界でのインターフェースを見てみます。

現実世界でのインターフェースは機械を使いましたがプログラムの世界ではクラスやメソッドの仕様を定義するためのものとして使われます。

例えばGoのinterfaceはメソッドの集まりを定義するためのものです。

自動販売機の例で言うと、お金を入れることで商品が出てくるという動きを期待しているので、その動きを定義すると以下のように考えられます。

```go
type VendinMachine interface {
    GetItem(int) string
}

// 飲み物の自動販売機
type DrinkVendingMachine struct {
    Items []string
}

func (d DrinkVendingMachine) GetItem() string {
    return d.Items[0]
}

// お菓子の自動販売機
type SnackVendingMachine struct {
    Items []string
}

func (s SnackVendingMachine) GetItem() string {
    return s.Items[0]
}

```
これらの例では、自動販売機のinterfaceを定義して、飲み物/お菓子それぞれの自動販売機がそのinterfaceを満たすように実装しています。

interfaceを満たすメソッドを実装することで自動販売機のinterfaceを満たすことができ、以下のように利用することができます。

```go

func main() {
    drinkVendingMachine := getDrinkVendingMachine()
    snackVendingMachine := getSnackVendingMachine()

    fmt.Println(drinkVendingMachine.GetItem()) // コーラ
    fmt.Println(snackVendingMachine.GetItem()) // ポテトチップス
}

func getDrinkVendingMachine() VendinMachine {
    return DrinkVendingMachine{Items: []string{"コーラ", "オレンジジュース"}}
}

func getSnackVendingMachine() VendinMachine {
    return SnackVendingMachine{Items: []string{"ポテトチップス", "チョコレート"}}
}

```

ここでのポイントはdrinkVendingMachineやsnackVendingMachineはそれぞれの自動販売機のインターフェースを満たすように実装されているので、それぞれのインスタンスをVendinMachine型として扱うことができています。

仮に各structの中身が入れ替わったとしてもGetItemメソッドがあればそれを利用することができるので、interfaceを使うことで柔軟にプログラムを書くことができます。

これのメリットはDBサーバーを変更したり、APIサーバーを変更したりする際に、特定のメソッドを有するinterfaceを使うことでその変更に対して柔軟に対応することができるということです。

interface(抽象)に依存することで柔軟に変更に対応できるというのがinterfaceのメリットです。

## まとめ

具体的な実装に依存することなく変更に強いプログラムを書くためにinterfaceへの理解を深めましょう！
