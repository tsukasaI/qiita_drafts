# Go con gopls

gopls 自動補完

Language Server Protocol
エラー、コードジャンプ、シンボル

client: Editor
server: gopls

Gopherにとってデフォルト、省エネ

もともとはdlvとか色々インストールする必要あったけどgoplsでまとまったイメージ

コンパイルなしでサジェストする？
-> コンパイルはしない。ソースコードを元に補完（gocodeというものがコンパイルしてるみたいだけどgo cacheの登場で上手くいかず）
gopls内でcache -> カーソルの位置情報で構文木を解析

どうやって保管するか？
-> Copilotとかと方式は異なる。
fmt.P と書いた瞬間に completionに渡る
- リクエスト
  - ファイルURI
  - 行数/位置
- リテラル内、関数呼び出し内では補完しない
- スコープの狭い順に探査をする
- 重みづけをして
