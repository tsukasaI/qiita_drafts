# Gitで過去のコミットを取り消す方法3選

## TL;DR

以下の3つが使える

```
gir revert

git reset --hard

git rebase -i
```

revertは取り消したいコミットを打ち消すコミットを作るのに対して
reset/rebaseはコミット自体を削除する。

**reset/rebaseを実行する前には必ずテスト用のブランチで試してから本番ブランチで行いましょう。**

## ざっと復習とお約束

Gitとは Version Controll Systemの一つでソースコードのバージョニングに使用されるサービスである。

本記事では個人的な趣味から「歴史を作る」という表現をします。

また対象のコミットは今回は簡略化のために一つ前のコミットとします。（実際一つ前のコミットを打ち消したいケースが多いと思うのでこれを丸暗記するだけでも一部対応可能）

## 利用シーンの想定

例えば以下ケース

- チームでGitを管理しているとブランチが複数存在し、本来入れたくないコミットを入れ込んでしまった
- DBのシークレット情報、APIキーをソースコードに含めてコミットしてしまった
- ~~恥ずかしいポエムをソースコードに入れっぱなしにしてコミットしてしまった~~

いざという時のために使い方をなんとなく把握しておきましょう

## revert

gitには以下で定義されている

`Revert some existing commits`

一つ前のコミットを打ち消す場合は以下のコマンド

```
git revert HEAD
```

コミットメッセージを編集する画面に遷移するので保存して終了。

**編集を取り消したいコミットを打ち消すコミットを作成するため、もともとHEADであったコミットはコミットは残ったままになることに注意**

## reset

gitには以下で定義されている

`Reset current HEAD to the specified state`

一つ前のコミットを打ち消す場合は以下のコマンド

```
git reset --soft HEAD^
```

コマンドを実行すると一つ前のコミット内容がstageに現れ、打ち消したいコミットがlogから消える

## rebase

gitには以下で定義されている

`Reapply commits on top of another base tip`
（なるほど、わからん）

一つ前のコミットを打ち消す場合は以下の操作

```
git rebase -i
```

を実行すると

```
pick ${commit hash} ${commit message}
pick ${commit hash} ${commit message}
pick ${commit hash} ${commit message}
pick ${commit hash} ${commit message}
pick ${commit hash} ${commit message}
```

ここで削除したいコミットをエディタで行を削除することでコミットの削除が可能
