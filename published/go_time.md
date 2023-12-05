Goのtime.Timeのformatが個人的にびっくりだった

# 時間をプログラムで扱いたい

リクエストやレスポンス、ロギングなど様々な場面で現在時刻を扱うと思います。

JavaScriptならDate型、Pythonならdatetimeモジュールを使うのではないでしょうか。

## 自由自在なフォーマット

JavaScriptでは`toString`, `toDateString`, `toLocaleString`, `toISOString` などのDate型で定義されたフォーマットで表示することができます。

一方で任意のフォーマットにする場合は年、月、日の値を取得することができます。

例えば`YYYY/MM/DD`にするならこんな感じに書きます。

```javascript
const now = new Date();
const year = now.getFullYear();
const month = ("0" + (now.getMonth() + 1)).slice(-2);
const day = ("0" + now.getDate()).slice(-2);

console.log(`${year}/${month}/${day}`);
```

処理としては

1. 現在の時刻をnowとして取得する
1. nowの年を4桁で取得してyearに格納
1. nowの月を取得して0パディングしてmonthに格納
1. nowの日を取得して0パディングしてdayに格納
1. year/month/dayの形式で表示

という流れになります。

## Goにおける日付操作

もちろんGoにも日付があります。

timeパッケージを使用します

リンク: https://pkg.go.dev/time

早速現在の時刻を`YYYY/MM/DD`で表示する処理を書いてみます。

まずはfmt.PrintfでJavaScriptのgetFullYearなどのメソッドを使った場合に似たコードを書いてみます。

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()
    fmt.Printf("%d/%02d/%02d\n", now.Year(), now.Month(), now.Day())
}
```

こんな感じで`%d`で数値を表示、`%02d`で2桁の0パディングした数値を表示できます。

さて、今回紹介したいtime.Timeのformatメソッドで書いてみるとどうでしょうか。

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()
    fmt.Println(now.Format("2006/01/02"))
}
```

なんと！

形式に`2006/01/02`というフォーマットでは見慣れない形式が出てきました。

これは何でしょうか？

調べるとすぐにヒットしますが、アメリカ式の日時の自然な順番を示します。

`1月2日午後3時4分5秒2006年”`

他のプログラミング言語ではなかなか見ない形で指定しています。

みなさんもGoで日時のフォーマットを見てびっくりしてみてください！
