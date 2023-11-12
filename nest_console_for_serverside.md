ネストが深いオブジェクトのconsole出力

# JavaScriptのデバッグ

JavaScriptであるオブジェクトの中身を確認したいとき、多くの人は`console.log`を使うと思います。

フロントエンドで開発するときにはconsole.logを使ってブラウザのコンソールを確認することで

どれだけネストしたオブジェクトもクリックして開くことで中身を確認することができて便利ですよね、

ただしバックエンドAPIを使っている場合はそうもいきませんでした。

## ネストされたオブジェクトのconsole.logで Objectが表示される問題

バックエンドでネストしたオブジェクトにconsole.logを使った場合は以下のようになりました。

```javascript
const nested = {
    nest1: {
      nest2: {
        nest3: {
          nest4: { nest5: "nest" },
        },
      },
    },
  };

console.log(nested);

// { nest1: { nest2: { nest3: [Object] } } }
```

このようにネストは3階層目までしか表示されず、リクエストや外部APIのレスポンスで複雑な構造をしている場合は中身が不明になってしまいます


## 解決方法

今回提示する解決策は3通り

- JSON.stringify を使う方法
- 置換文字列を使う方法
- console.dir を使う方法

### JSON.stringify を使う方法

JavaScriptのobjectを文字列に置換するのに使われる関数で、オブジェクトのdeep copyをしたいときに調べるとヒットしますね（経験上あまり使わないほうがよい）

```javascript

console.log(JSON.stringify(nested));


// {"nest1":{"nest2":{"nest3":{"nest4":{"nest5":"nest"}}}}}
```

確かにnest5まで表示されていて一見は良さそうですが、キーにダブルクォートがついていたり
、文字列のため改行されていないため見づらく見えます。

### 置換文字列を使う方法

C言語などで使われるフォーマット指定子っぽいものがJavaScriptにもあります。

参考 https://developer.mozilla.org/ja/docs/Web/API/console#%E4%BE%8B


この中で `%o`を使うことで最下部までコンソールに表示されるようになりました

```javascript

console.log("%o", nested)

```

```
{
  nest1: {
    nest2: { nest3: { nest4: { nest5: 'nest' } } }
  }
}
```

いい感じにインデントもしてくれていて見やすくなっていますね

### JSON.stringify を使う方法

consoleにはlog以外にもメソッドがあります。その中でdirを使ってみます。

console.dirは引数を2つとり、第一引数は任意のitem、第二引数はoptionを指定します。

optionにはobjectとして `{ depth: null }`を入れましょう。

```javascript

console.log(JSON.stringify(nested));

```

```
{
  nest1: {
    nest2: { nest3: { nest4: { nest5: 'nest' } } }
  }
}
```
こちらも見やすく表示されました。
