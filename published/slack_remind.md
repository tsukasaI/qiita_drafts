# アプリを使わずにSlackでいい感じにリマインドをセットするコマンドとサンプル

多くの職場の連絡でSlackがよく使われると思います。

個人的にはSlack用の個人アプリを好きに追加できたり、Jira/GitHubからの通知を受け取ったり、canvas機能があったりしてお気に入りアプリの一つです。

そんなSlackで便利なリマインドを紹介して皆さんにより好きになってもらいたいと思いこの記事を作成しました。


## リマインダーとは

公式の説明は[こちら](https://slack.com/intl/ja-jp/help/articles/208423427-%E3%83%AA%E3%83%9E%E3%82%A4%E3%83%B3%E3%83%80%E3%83%BC%E3%82%92%E8%A8%AD%E5%AE%9A%E3%81%99%E3%82%8B)

Slackのリマインダーは特定のタスクやイベントを忘れないようにするための超便利な機能です。

リマインダーは `/remind` コマンドを使用して設定できます。これを使いこなすことでサクッと設定できます。

以下にいくつかの例を見てみましょう。Slackアプリで表示されるサンプルを示します。

```
/remind me on June 1st to wish Linda happy birthday

-> 6/1に wish Linda happy birthdayメッセージが現れる

/remind #team-alpha to "Update the project status today" Monday at 9am

-> 月曜日AM9時に #team-alphaチャンネルで Update the project status todayメッセージが現れる
```

基本構文は `/remind {対象} {内容} {いつ}` となっています

これらのコマンドは、リマインダーを設定するための基本的な例です。必要に応じてカスタマイズして使用できます。

## 内容にメンションを含める方法

リマインド内容にメンションを含めるにはダブルクオーテーションで囲みましょう。

メンションもハイパーリンクも自在に入れることができます。

## 定期的なリマインダーにしてみる

いつの部分にevery ~ を入れることで定期的なリマインダーにできます。

例えば

```
/remind me to "毎月のリマインダー" on every 10th
```

とすることで毎月10日にリマインダーを出してくれます。

## 曜日指定

曜日固定はevery {曜日}でできます

```
/remind me to "毎週13時のリマインダー" at 13:00 every Thursday.
```


## 隔週のリマインダーだってお手の物

2週間に一度の予定もあるでしょう。その場合は`every 2 weeks`とすると可能です。

```
/remind me to "毎週13時のリマインダー" at 13:00 every 2 weeks Monday.
```

また次に出してほしい日付を指定する場合は最後に`next occurrence is {日付}`を入れると定期リマインダーのスタート日を決められる。

公開日である12/12からリマインドする場合は`next occurrence is Dec 12th`とする。


便利なSlack Lifeをお楽しみください。
