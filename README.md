# qiita_drafts

## future

- algorithmの勉強でやったソートをGoでまとめる
- orverload

```
  private getKeys<T extends { [key: string]: unknown }>(obj: T): (keyof T)[] {
    return Object.keys(obj);
  }
```


* GoのCIでテストのカバレッジをコメントするようにしたった
* Gomockを使ってテストのmock入れたら意外と大変だった
* もう怖くない楽しい正規表現
    * 最初から色々やろうとするから難しくなる
    * まずは文字の意味を知る
    * 少しだけテクいことを
        * ., ?, ^, $, *
* serial使う？UUID使う？ULID使う？
    * セキュリティ
    * 文字数
    * 順序性
* WinとMacで部分的にスクショを撮るショートカット
    * Win
    * Mac
* Goのtime.Timeの書式が特殊で初見殺し過ぎる
    * ワイ記法
* ローカルでServerless Framework と AWS Lambdaを使ってGoを動かしたい
* copilotは万能ではない話
    * 全ておまかせではこっちの意図は汲み取ってくれない。（あくまで学習して予想したサジェストをしてくれるだけ）
* ファイルの最終は改行を入れる
* （あと4個）
