# PUTとPATCHと冪等性

REST APIの設計に置いて更新系のリクエストはHTTPのPUTとPATCHが担う事になっている。

それらの違いが気になったので少し調べた

## 結論

PUTとPATCHの違いは冪等性にある。

PUTはリソースの完全な置き換えを行うため、同じリクエストを何度行っても同じ結果が返ってくる。

一方でPATCHはリソースの部分的な更新を行うため、同じリクエストを何度行っても同じ結果が返ってくるとは限らない。

## PUT

PUTはリソースの完全な置き換えを行うため、リクエストボディにはリソースの完全な情報を含める必要がある。

同じリクエストを何度実行しても、サーバー上のリソースの状態は最初の実行後と変わりません。（つまり冪等性を持つ）

またPUTはリソースが存在しない場合は新規作成を行う。

## PATCH

PATCHはリソースの部分的な更新を行うため、リクエストボディには更新したいリソースの一部の情報だけを含める。

同じリクエストを何度実行しても、サーバー上のリソースの状態は最初の実行後と異なる可能性があります。（つまり冪等性を持たない）

そのためPATCHはPUTと異なり、リソースの完全な情報を含める必要がないため、リクエストボディのサイズが小さくなる場合があります。


## まとめ

サービスの動きによって使い分けるのであれば

- PUT: オブジェクトストレージなどのリソースの完全な置き換えを行う場合
- PATCH: データベースなどのリソースの部分的な更新を行う場合

が適切かと思われる。

（がPUTを使ってよいのでは）
