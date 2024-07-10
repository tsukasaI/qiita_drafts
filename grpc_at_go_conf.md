# gRPC入門

ノリでgRPCを勉強してみようと思ったら楽しかったので書きます。

## gRPC is 何

公式ページは[こちら](https://grpc.io/)。

ざっくりサマるとハイパフォーマンスなRPC(Remote Procedure Call)のこと。

異なるサービス間のやり取りを実装するプログラミング言語に縛られずにストリーミングでの情報伝達も可能にするRPCの一種です。

REST APIのようにパスやメソッドは作らずにメソッドと引数を定義する。

データフォーマットにProtocol Bufferを使いバイナリ化され、型付けされたデータを転送できる。

使いどころとしては
- Microserviceのバックエンドサーバー間の通信
- モバイル向けがサービスの通信（通信量を節約）
- スピードを求められる通信

などがある。

またgRPCではHTTP/2を使うため、HTTP/1.1では実現できなかったサーバープッシュなどが可能になります。

| HTTP/1.1 | HTTP/2 |
| -- | -- |
| 1リクエスト1レスポンス | ストリームで1TCP接続で複数のリクエストとレスポンスが可能 |
| ヘッダーのオーバーヘッド | ヘッダーの圧縮/キャッシュで差分のみの送受信が可能 |
| - | サーバープッシュが可能 |

## Protocol Buffer

公式ページは[こちら](https://protobuf.dev/)。

スキーマ言語で構造化されたデータをシリアライズするために使われる。

gRPCでのリクエストとレスポンスに用いるデータフォーマットで実装するプログラミング言語は色々使える。

メリットとしては

- 型安全
- バイナリ化されるため文字列で扱うよりも高速で通信可能

といった点があります。

JSONとの比較をすると下のような感じです。

| JSON | Protocol Buffers |
| -- | -- |
| 広く使われている | 少数 |
| ほとんどのプログラミング言語で使える | 一部プログラミング言語のみ |
| ネストが自由自在 | 複雑な構造には不向き |
| 人が読みやすい | バイナリ化された後では人間には読めない |
| スキーマを強制できない | 型が保証される |
| データサイズは大きい | 小さいデータサイズ |

利用の流れとしては

1. スキーマ定義
1. 各言語のオブジェクト生成
1. バイナリにシリアライズする

という流れです。

### Protocol Bufferファイル

まずはサンプルのファイルを見てみましょう。

```proto
syntax = "proto3";

message Person {
    string name = 1;
    int32 id = 2;
    string email = 3;
}
```

`syntax = "proto3";`はバージョンの宣言で特段理由がなければ3を使いましょう。

messageは複数のフィールドを持つことができ、各フィールドは型情報、フィールド名、数値を設定します。

各スカラー型などは[こちら](https://protobuf.dev/programming-guides/proto3/#scalar)を参照。

数値は`1` から `536,870,911`を利用可能（ただし`19,000` から `19,999`は予約されているため使用しない）で再利用できない様になっています。

Protocol Bufferはフィールドは名前ではなくタグ番号で管理するため一度採番したら変更できないことに注意しましょう。

ちなみに数値は1~15がよく使われるらしい（1byteでいけるため）。


### Service

RCPの実装単位でRPCメソッドを定義する。
1サービスに複数のメソッドを定義できる。

通信方式が4種ありそれぞれ

- Unary RPC
- Server Straming RPC
- Client Straming RPC
- Bidirectional Straming RPC

です。

#### Unary RPC

クライアントから1リクエスト、サーバーから1レスポンスの通信方式

（REST APIでもお馴染み）

#### Server Straming RPC

クライアントから1リクエスト、サーバーから複数レスポンスの通信方式

クライアントはサーバーからの送信完了まで通信を続ける。

プッシュ通知などに使える。

#### Client Straming RPC

クライアントから複数リクエスト、サーバーから1レスポンスの通信方式

クライアントの終了を持ってレスポンスを返す。

大容量ファイルのアップロード時にチャンクするような時に使う。


#### Bidirectional Straming RPC

クライアントから複数リクエスト、サーバーから複数レスポンスの通信方式

チャット、オンライン対戦ゲームで使うと良い。

## いざ実践

ここまでgRPCとProtocol Bufferのあれこれを学んだので、実際に動かすまでを書いてみる。

Protocol Bufferのコードは以下。

Goで実装を行うことを前提とします。

```proto
syntax = "proto3";

package sample;

option go_package = "./pb";

message ListSampleRequest{};

message ListSampleResponse{
    repeated string samplenames = 1;
};

message DownloadRequest {
    string samplename = 1;
};

message DownloadResponse {
    bytes data = 1;
};

message UploadRequest {
    bytes data = 1;
};

message UploadResponse {
    int32 size = 1;
};

message UploadAndNotifyProgressRequest {
    bytes data = 1;
};

message UploadAndNotifyProgressResponse {
    string message = 1;
};

service SampleService {
    rpc ListSample (ListSampleRequest) returns (ListSampleResponse);
    rpc Download(DownloadRequest) returns (stream DownloadResponse);
    rpc Upload(stream UploadRequest) returns (UploadResponse);
    rpc UploadAndNotifyProgress(stream UploadAndNotifyProgressRequest) returns (stream UploadAndNotifyProgressResponse);
}

```

いくつか補足

### package

Protocol Bufferでは複数ファイルの分割が可能になっている。

他のファイルをimportするときには`パッケージ名.型名`とし、名前空間を定義できる。

### option

メタデータをいれることができる。

`option go_package = "./pb";`はGoで生成されるコードをどのパッケージに入れるかを設定します。

### service

serviceにはRPCメソッドを定義します。

例えば`rpc ListSample (ListSampleRequest) returns (ListSampleResponse);`ではUnary RPCのメソッドを定義しています。

リクエストの方を1つ目の()内に、レスポンスの方をreturnsに続く()に入れます。

またstreamにする場合は方の前に`stream`を入れます。

これでProtocol BufferからGoのコードを生成します。

```sh
$ protoc -I. --go_out=. --go-grpc_out=. proto/*.proto
```

<!-- 残りGoのコード書く -->
