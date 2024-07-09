# gRPC入門

gRPC Protocol Buffer スキーマ言語

マイクロサービス

宣言的なファイル

gRPCのデータフォーマット　言語は色々使える

バイナリ化可能で高速で通信

型安全

JSONに変換可能

| JSON | Protocol Buffers |
| -- | -- |
| めっちや使われる | 少数 |
| ほとんどの言語で使える | 一部言語のみ |
| ネストが自由自在 | 複雑な構造には不向き |
| 人が読みやすい | バイナリ化された後では人間には読めない |
| スキーマを強制できない | 型が保証される |
| データサイズは大きい | 小さいデータサイズ |

利用の流れ
1. スキーマ定義
1. 各言語のオブジェクト生成
1. バイナリにシリアライズして通信する


ファイル

version だいたい3を選ぶ

message

複数のフィールドを持てる
スカラ型、コンポジット型で使える

形式
```proto
message Person {
    string name = 1;
    int32 id = 2;
    string email = 3;
}
```

型 フィールド名 = タグ番号（代入してるわけではない）

[スカラー型](https://protobuf.dev/programming-guides/proto3/#scalar)

タグ

Protocol Bufferはフィールドは名前ではなくタグ番号で管理する

1~15がよく使われる（1byteでいける）


デフォルト値

列挙型: タグ番号が0

## gRPC

Googleが2015年に公開（どこでも話されているがGoogleのgではない）

RPC Remote Procedure Call

REST APIのようにパスとかメソッドは作らずにメソッドと引数を送る

データフォーマットにProtocol Buffeerを使いバイナリ化され、型付けされたデータを転送できる。

HTTP/2を使用する

特定のプログラミング言語に依存しない。

使いどころ
- Microserviceのバックエンドサーバー間通信
- モバイルユーザーが利用するサービスで通信量を節約
- スピード

| HTTP/1.1 | HTTP/2 |
| -- | -- |
| 1リクエスト1レスポンス | ストリームで1TCP接続で複数のリクエストとレスポンスが可能 |
| ヘッダーのオーバーヘッド | ヘッダーの圧縮/キャッシュで差分のみの送受信が可能 |
| - | サーバープッシュ |


### Service

RCPの実装単位
メソッドがエンドポイントになる。
1サービスに複数のメソッドを定義可能

4種の通信方式

- Unary RPC
- Server Straming RPC
- Client Straming RPC
- Bidirectional Straming RPC

### Unary RPC

1リクエスト 1レスポンス
普通の関数

### Server Straming RPC

1リクエスト複数レスポンス

クライアントはサーバーからの送信完了まで通信を続ける

プッシュ通知などに使う

### Client Straming RPC

複数リクエスト 1レスポンス

クライアントの終了を持ってレスポンスを返す

大容量ファイルのアップロード


### Bidirectional Straming RPC

複数リクエスト、複数レスポンス

チャット、オンライン対戦ゲーム
