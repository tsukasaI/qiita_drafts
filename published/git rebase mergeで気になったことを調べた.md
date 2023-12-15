git rebase mergeで気になったことを調べた

# Gitが全くわからない

Git難しくないですか？

あるブランチの変更を別のブランチに統合するための手法がいくつか準備されていてどれ使っていいかわからん。

それぞれの特性が何なのか？今どっちを選択すべきか？

そんな疑問が一生ついてまわっていた。

ということで今回はmergeとrebaseについて調べてまた一つ成長していこうと思う。

## 登場人物

- mainブランチ
- featureブランチ

## mergeの基本

マージ(merge)とは`複数のものを統合して一つのものにする`という意味で、gitでも2つのブランチの変更内容をまとめるという操作をする。

featureブランチにmainブランチの変更コミットC3をmergeする場合は以下のようになる。

merge前
```
C1---C2---C3      : main
      \
      C4---C5     : feature
```

```sh-session
$ git merge main
```

merge後
```
C1---C2---C3           : main
      \
      C4---C5---C6     : feature
```

mainのC3コミットを取り込む際には新しいコミットC6が生成される。


ちなみにfeatureブランチにC4, C5コミットがない場合はfast-forward mergeとなり、C3コミットのまま取り込まれる


merge前
```
C1---C2---C3      : main

C1---C2           : feature
```

```sh-session
$ git merge main
```

merge後
```
C1---C2---C3           : main

C1---C2---C3           : feature
```

## rebaseの基本

rebaseとはベースを変更するみたいな意味を持つ（と思っています）。

rebase前
```
C1---C2---C3      : main
      \
      C4---C5     : feature
```

```sh-session
$ git rebase main
```

rebase後
```
C1---C2---C3                    : main

C1---C2---C3---C4---C5---C6     : feature
```

このようにブランチの派生元の変更をベースに取り込んで現在のブランチの変更を追加したような形になる。

## じゃあ全部rebaseで良くない？

個人的にはrebaseは反対している。

理由は2つで

- 過去のコミットの履歴を変更するからリモートリポジトリにpushするときにforce pushになる
- コンフリクトを解消するのとても大変

### 過去のコミットの履歴を変更するからリモートリポジトリにpushするときにforce pushになる

upstreamブランチに対してpushする際に`failed to push some refs to`というメッセージを見たことある人は多いだろう。

これはブランチをpushする際にはfast-fowardであることが期待されていて、履歴が変更されると通常のpushはできなくなる。

その場合はforce pushする必要があるが、これはとても危険。

履歴が強制的に書き換わるためチーム開発なら全員のローカルリポジトリに影響を及ぼす。

可能な限りやめていただきたいところ。

### コンフリクトを解消するのとても大変

コンフリクトはエンジニア的に超嫌なワードだと思う。

以下の状況を考える。

rebase前
```
C1---C2---C3      : main
      \
      C4          : feature
```

C3とC4のコンフリクトがあった場合にrebaseの結果どうなるかを解説する。


```sh-session
$ git rebase main
```

アウトプット
```
Auto-merging {ファイル名}
CONFLICT (content): Merge conflict in {ファイル名}
error: could not apply {コミットハッシュ}... update {ファイル名}
hint: Resolve all conflicts manually, mark them as resolved with
hint: "git add/rm <conflicted_files>", then run "git rebase --continue".
hint: You can instead skip this commit: run "git rebase --skip".
hint: To abort and get back to the state before "git rebase", run "git rebase --abort".
Could not apply {コミットハッシュ}... update {ファイル名}
```

はい、読んでるだけで嫌になりますね。

基本的にはファイルでコンフリクトが起きているので解消してから以下コマンドで進められます。

```sh-sesshon
$ git add {ファイル名}
$ git rebase --continue
```

アウトプット
```
Successfully rebased and updated refs/heads/feature.
```

はい、めでたくrebaseはできました。この状況でlogを確認するとブランチの履歴は以下のようになる。

rebase後
```
C1---C2---C3              : main

C1---C2---C3---C4         : feature
```

コンフリクトを解消することで同様にrebaseできました。

## 結論どっち使うか

正直理解しているならどっち使ってもOK。

gitの作者がコマンドとして作っているので、目的にあうコマンドを選びましょう。

自信のない方は一旦全部コミットしておけば後からの変更はなんとかやりやすい。

こっそり練習してgitのコマンドの意味と動きを確認しましょう！
