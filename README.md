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
* gitで操作ミスしたときに見る落ち着いて対応する方法
    * 仙道「まだ慌てる時間じゃない」
    * 状況を整理する
    * 大前提わからないときには分かる人に聞く
    * ブランチをコピー
    * mainにコミットした -> すぐ消してはならない
    * reset, revert, rebaseどれ使うか
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
* &&とか||の詳しい挙動についてTypeScriptを使って実験してみた
    * &&は左がtrueで右を評価
    * ||は無条件で右を評価
    * Reactでこういうのみない？ boolVar && ~
* ローカルでServerless Framework と AWS Lambdaを使ってGoを動かしたい
* アプリを使わずにSlackでいい感じにリマインドをセットするコマンドとサンプル
* copilotは万能ではない話
    * 全ておまかせではこっちの意図は汲み取ってくれない。（あくまで学習して予想したサジェストをしてくれるだけ）
* （あと5個）
