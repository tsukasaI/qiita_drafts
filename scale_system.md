# スケールするシステムを作りたいときに考慮する設計

Webシステムをスケールさせるためにはパフォーマンスを考える必要がある。

1人だけに使われるシステムを作る場合は単一のサーバーにそのままリクエストしてレスポンスを返すだけで十分に動くことが多い。

しかし1億人に使われるシステムを作るにはそれだけでは十分ではなくなってくる。

この記事では1人に使われるシステムから1億人に使われるシステムをスケールさせる際に考えるべき点を書いていく。

## 1人に使われるシステム

フロントエンドとバックエンドサーバーとデータベースサーバーを構築することで1人（実際にはもっと人数がいてもいい）が利用する分には問題ないことが多い。

バックエンドサーバーとDBサーバーとにおける処理のおおよそのフローとしては2つ。

1. バックエンドサーバーはクライアントからリクエストを受け取る
1. バックエンドサーバーからDBサーバーにクエリを発行してクライアントにレスポンスを返す

ユーザー数が少なければこれで十分動き続ける。

## 1億人に使われるシステム

ユーザー数が増えるに連れて各サーバーの負荷が上がってくる。

サーバーのリソースが枯渇してしまうとレスポンスを返すことができなくてユーザーが困ってしまう。

こんな悲しい状態にならないために各サーバーを冗長構成にして負荷分散をすることを考えましょう。

### バックエンドサーバーの冗長化

クライアントからリクエストが来た時に単一のサーバーで処理しきらずに複数の同一の処理をするサーバーに分散させることができる。

一般的にクライアントからのリクエストは直接バックエンドサーバーに向かわせずにロードバランサー(LB)に向かわせることが多い。

LBは複数のバックエンドサーバーへ振り分けて1台あたりの負荷を減らすことでリソースが枯渇することを回避させることが可能になる。

またLBを導入するとあるサーバーが障害を起こして止まってしまった場合には別のサーバーに処理を流すように設定することができ、サービスを使い続けられる状況を維持することができるようになる。

### DBサーバーの冗長化

DBサーバーには

## 1億人に快適に使われるシステム

### キャッシュの導入

### 非同期処理の導入
