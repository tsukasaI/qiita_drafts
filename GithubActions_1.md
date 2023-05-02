# Github Actionsで過去のコミットを参照しようとしたらunknown revision or path not in the working tree と言われたときの解決法

# TL;DR

actions/checkout@v3 はデフォルトでheadのみ設定になるため
過去のコミットを全て取得する場合はwithで `fetch-depth: 0` とする必要がある。

```yaml
~
jobs:
  codedeploy:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout repository
        uses: actions/checkout@v3
        with:
          fetch-depth: 0 # ここを追加してあげる
```
リポジトリにきちんと説明されている
参考
https://github.com/actions/checkout#fetch-all-history-for-all-tags-and-branches

# 背景

Github ActionsでJSへのトランスパイルとオートデプロイを別ワークフローで管理するように構築していたため、

最新のコミットが最後にプッシュされたコミットと比較したときに、特定の拡張子のファイル（例えば.vue, .scssなど）が変更されているかを判定する必要があった。

# できなかった実装

```yaml
~
jobs:
  codedeploy:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout repository
        uses: actions/checkout@v3

      - name: get build required files
        id: get-build-required-files
        run: |
          BaseHash=`git merge-base HEAD HEAD^`
```

[actions/checkout](https://github.com/actions/checkout) でブランチチェックアウトをして
`git merge-base`コマンドで過去のコミットハッシュを取得したかった。

が `unknown revision or path not in the working tree` というエラーが発生して上手くいかず。。。

# 原因と解決方法

actions/checkout@v3の設定には `fetch-depth` があり、デフォルトでは 1 に設定されている。
上記によりデフォルトではheadのみfetchするようになるため、すべてのコミットを取得するために`fetch-depth = 0` とすると良い。

# 上手くいった実装

```yaml
~
jobs:
  codedeploy:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout repository
        uses: actions/checkout@v3
        with:
          fetch-depth: 0

      - name: see past commit
        run: |
          BaseHash=`git merge-base HEAD HEAD^`
          ~
```

# 最後に
ローカルで同一の `git merge-base HEAD HEAD^` を実行してもエラーが出なかったため原因突き止めるまでに時間がかかった。
使用するライブラリのReadmeを早めに読んでおくべきでした。
