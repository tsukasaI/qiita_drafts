# はじめに
本投稿はDynamoDBの入門記事です。
Webエンジニアが初めてDynamoDBを利用したサービスを構築した際に学習したことを記する。

# DynamoDBとは
`1 桁ミリ秒単位で規模に応じたパフォーマンスを実現する高速で柔軟な NoSQL データベースサービス`です。

Oracle, MySQLなどのリレーショナルデータベース(RDB)に対してNoSQLは"Not Only SQL"と言われていてRDB以外を指す。
例えば以下のタイプがある

## Key Value Store(KVS)
その名の通りKeyとValueのデータを管理するデータストアでDynamoDBはこのタイプ。
高度なパーティション化に対応しており、他のタイプのデータベースでは達成できない大規模な水平スケーリングが可能です。

## ドキュメント型
JSONやXMLなどの形式で書かれたデータを管理するストア。
構造を柔軟に決めることが可能で開発者からしたら直感的に操作できる。
MongoDBがこのタイプ。

## インメモリ
ディスクベースのデータストアではなくメインメモリ上にデータをストアする。基本的に早い
Amazon ElastiCache は、1 秒あたり数億回の操作までスケーリングが可能。

---

# 設計
RDBとは異なる概念となるので、RDBに慣れている人は特に注意されたし。

DynamoDBはテーブルのデータをパーティションと呼ばれる領域に分けて保持する。

データがどのパーティションへ配置されるかは”パーティションキー”によって決定されます。
パーティションにデータを格納する場合は”ソートキー”として指定した順序でデータを格納する。
データが格納されすぎているパーティションがある場合（そのようなパーティションをホットパーティションといいます）、データのクエリ速度が遅くなる場合があります。また、データの検索手段として、フルスキャンと、パーティションキーを指定した上でのソートキーによる範囲を指定したスキャンの方法があり、ソートキーを指定している場合は後者の検索が可能となり、検索速度が向上する。
（パーティションとソートキーの設計が超重要）

またDynamoDBはインデックスとして、LSI（Local Secondary Index）とGSI（Global Secondary Index）の2種類がある。LSI は、Partition Keyが同一なアイテムを、ほかのアイテムからの検索するために利用する。


## 注意

データ構造だけ定義しクエリーを柔軟に発行できた RDBMS と違い、DynamoDB はクエリーが柔軟ではありません。
また、データ集計にも強くありません。そのため、設計段階で各種データへどのようなアクセスがされるかをきちんと考えておく必要があります。

また、システム全体をDynamoDBだけで無理やり実現しようとせずに、DynamoDBに向いている機能に限定して採用しようとする考え方が大切です。

 ~残念ながら~ 銀の弾丸にはなっていない。

# 操作

## テーブル作成
```
aws dynamodb create-table \
    --table-name kddi-cms-dashboard \
    --attribute-definitions \
        AttributeName=CorporateCode,AttributeType=S \
        AttributeName=TargetGenderAgeDate,AttributeType=S \
    --key-schema \
        AttributeName=CorporateCode,KeyType=HASH \
        AttributeName=TargetGenderAgeDate,KeyType=RANGE \
    --provisioned-throughput \
        ReadCapacityUnits=5,WriteCapacityUnits=4000 \
    --table-class STANDARD --endpoint-url http://172.17.0.1:8000
```

## データ投入
```
aws dynamodb put-item --table-name 'kddi-cms-dashboard' --item '{"CorporateCode": { "S": "80000011" },"TargetGenderAgeDate": {"S": "step_avg#female#10歳未満#2022-12-01"}, "Value": {"N": "5000"}}' --endpoint-url http://172.17.0.1:8000
```

## クエリ

```
aws dynamodb query --table-name kddi-cms-dashboard \
    --key-condition-expression 'CorporateCode = :code and begins_with (TargetGenderAgeDate, :sub)' \
    --expression-attribute-values '{":code": { "S": "00000000" },":sub": {"S": "step"}}' \
    --endpoint-url http://localhost:7790
```
