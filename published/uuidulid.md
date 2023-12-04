# SERIALつかう？UUIDつかう？新たな選択肢ULID

SERIALとUUIDの両方がDBのプライマリーキーとして広く使用されている。

これらそれぞれには利点と欠点があります。

## SERIAL

メリット

- 生成が簡単。データベースが自動的に次のIDを生成する。
- **データのソートが容易**。IDの順番はデータの作成順と一致します。
- スペース効率が良い。整数はUUIDよりも消費ストレージが格段に少ない。

一方でデメリット

- 分散システムではIDの一意性を担保できない。
- ユーザーが他のレコードのIDを推測しやすい。（セキュリティ上の問題）

## UUID

メリット

- 一意性が高い。UUIDは全世界で一意らしい。
- 分散システムでも安全に使用可能。各システムは独立してUUIDを生成できる。
- レコードのIDを推測するのが困難。

デメリット

- 通常128ビットのスペースを消費し、整数型のIDよりも大きい。
- 人が読みにくいし覚えにくい。
- ランダムに生成されるため、レコードの作成順を保証できない。

## 比較

以上から

| | SERIAL(整数) | UUID |
| -- | -- | -- |
| 一意にしやすいか | No | Yes |
| ソート順がコントロールできるか | Yes | No |
| 他のIDが推測されにくいか | No | Yes |


どちらを選ぶかは要件とトレードオフによる

ここで今回紹介したいのはUUIDとSERIALの良いところをいい感じに併せ持つULIDをを紹介したい。

## ULIDとは

Universally Unique Lexicographically Sortable Identifierの略で特徴は以下

- 128ビットの識別子で、タイムスタンプ部分とランダム部分から構成される。
- タイムスタンプ部分は48ビットで、ミリ秒単位での精度を持つ。
- ランダム部分は80ビットで、一意性を保証する。

すごそうな説明となっていますね。一般的なメリットとデメリットを見てみましょう。

メリット

- タイムスタンプのおかげで生成された順序に従ってソートできる！
- UUIDと同様に、全世界で一意！
- 分散システムでも安全に使用できる！（各システムは独立してULIDを生成可能）
- レコードのIDを推測するのが困難！

さてデメリットはというと

- UUIDと同様に128ビットのスペースを消費し、整数型のIDよりも大きい。
- 人が読みにくいし覚えにくい。

UUIDのデメリットは完全には取り除けはしませんが、いくらか改善できていますね。


## 1ミリ秒での順序を検証

今回はTypeScriptで簡単な例を作成してみました

```typescript
import { ulid } from "ulid";

const printer = () => {
  let i = 1;
  const id = setInterval(() => {
    if (i > 10) clearInterval(id);
    console.log(new Date().getTime(), ":", ulid());
    i++;
  }, 1);
};

printer();
```

解説するとprinter関数内では現在時刻のUNITTIMEとulidの値を出力します。

setIntervalを使って1ミリ秒の間隔を空けてprinter関数をコールしています。

結果はこちら（UNITTIMEは本質ではないので1ミリ秒がわかる範囲でマスクしています。）

```
xxxxxxxxxxx58 : 01HGQ8RQPAFDMGHG5TF2HE6QBJ
xxxxxxxxxxx59 : 01HGQ8RQPBF8RJ4D7N2RS70QGN
xxxxxxxxxxx60 : 01HGQ8RQPC63EJ4JF38TB7FSQG
xxxxxxxxxxx61 : 01HGQ8RQPDGRMYACBK1BBZPN7D
xxxxxxxxxxx62 : 01HGQ8RQPEHBJYXMG00F7ZJW16
xxxxxxxxxxx63 : 01HGQ8RQPFXFQTMZKCE3W3YPRV
xxxxxxxxxxx64 : 01HGQ8RQPGR6J25HWGRMX1H7BK
xxxxxxxxxxx65 : 01HGQ8RQPHKGKZKE1C304NZ6M7
xxxxxxxxxxx66 : 01HGQ8RQPJG0KPFKT1EWB29XHC
xxxxxxxxxxx67 : 01HGQ8RQPKVGV9FFY7X370BA5S
xxxxxxxxxxx68 : 01HGQ8RQPMEQCWE6GGN2P1HW9C
```

さて`ulid()`の値に着目してみましょう。

先頭9桁はすべて`01HGQ8RQP`となっていて、次の文字がABCDE...と続いていて確かに時間順にならんでいます。

順序が守られないUUIDとSERIALの特徴を併せ持つフォーマットですね。

みなさんも使ってみてはいかがでしょうか！
