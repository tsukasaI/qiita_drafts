残り
SQLとの直接的な対話とORMを使用するときの違いを説明。
type User struct { の部分のDDLと各メソッドにおけるDMLを一通り
リレーションPreload

# タイトル: ORM入門 Go言語 Gorm パッケージ 解説

はじめに

IGD LXU所属の井上です。今回は我々のプロジェクトで利用している技術の一つであるORMを紹介します。

## ORMとは

ORM（Object-Relational Mapping）は、プログラムのオブジェクトとリレーショナルデータベースのデータをマッピングする技術です。

これにより、データベースの操作をプログラムのオブジェクトを操作するような形で行うことができます。

例えばMySQLのあるテーブルの1レコードをGoのある構造体の実態としてマッピングして構造体で扱うことができるようになります。

ORMの主なメリットは以下の通りです：

- 生のSQLを書く必要がない：ORMはデータベース操作を抽象化し、開発者がSQLクエリを直接書く必要を減らします。これにより、コードの可読性と保守性が向上する。

- データベースの変更が容易：ORMを使用すると、使用しているデータベースの種類を変更するのが容易になる。ORMがデータベースの詳細を抽象化しているため、データベース固有のSQLを書く必要がなく、データベースの種類を変更してもコードの大部分を変更する必要がなくなる。

- コードの再利用性と組織化：ORMはモデルベースのアプローチを採用しており、これによりコードの再利用性と組織化が向上します。モデルはビジネスロジックをカプセル化し、これによりコードの整理と再利用が容易になります。

- 開発速度の向上：ORMはデータベース操作を大幅に簡素化し、開発者がより高速にアプリケーションを開発できるようにします。

これらの利点により、ORMはWebアプリケーションの開発において広く使用されARISEのプロジェクトでも採用されています。

そんなバックエンドエンジニアならぜひ覚えておきたい技術をGo言語（以下Go）で説明をしていきます。

## GoとGorm

ここからはORMの操作をGoで利用可能なORMパッケージの一つであるGormで紹介をします。

## Gormを使ったデータベース操作

MySQLにセッションの貼る方法は以下のようになります。

```go
package main

import (
    "fmt"
    "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
    dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

    if err != nil {
        fmt.Println("Failed to connect to database")
        panic(err)
    }

    fmt.Println("Database connection successfully opened")
}
```

続いてプログラム内で操作する構造体の定義を紹介します。

ここでは公式ドキュメントの構造体をそのまま引用します。

```go

type User struct {
  ID           uint           // 主キーの標準フィールド
  Name         string         // 通常の文字列フィールド
  Email        *string        // 文字列へのポインタ、nullを許可
  Age          uint8          // 符号なし8ビット整数
  Birthday     *time.Time     // time.Timeへのポインタ。nullを許可
  MemberNumber sql.NullString // sql.NullStringを使用して、null可能な文字列をハンドリング
  ActivatedAt  sql.NullTime   // sql.NullTimeを使用したnull可能な時間フィールド
  CreatedAt    time.Time      // GORMによって自動的に管理される作成時間
  UpdatedAt    time.Time      // GORMによって自動的に管理される更新時間
}

```

Gormを使った基本的なデータベース操作（CRUD操作）の例を示します。

Select操作：Gormを使ってデータを取得する方法を説明します。

```go

var user User
result := db.First(&user, "name = ?", "John")
if result.Error != nil {
    // エラーハンドリング
    fmt.Println(result.Error)
    return
}
fmt.Println(user)

var users []User
result := db.Find(&users, "age >= ?", 18)
if result.Error != nil {
    // エラーハンドリング
    fmt.Println(result.Error)
    return
}
for _, user := range users {
    fmt.Println(user)
}

```

Insert操作：Gormを使って新しいレコードを挿入する方法を説明します。

```go
user := User{Name: "John", Age: 25}
result := db.Create(&user)
if result.Error != nil {
    // エラーハンドリング
    fmt.Println(result.Error)
    return
}
fmt.Println(user)

users := []User{
    {Name: "John", Age: 25},
    {Name: "Jane", Age: 30},
    {Name: "Bob", Age: 20},
}

result := db.Create(&users)
if result.Error != nil {
    // エラーハンドリング
    fmt.Println(result.Error)
    return
}

for _, user := range users {
    fmt.Println(user)
}

```

update

```go
var user User
db.First(&user, "name = ?", "John")

user.Age = 30
result := db.Save(&user)
if result.Error != nil {
    // エラーハンドリング
    fmt.Println(result.Error)
    return
}
fmt.Println(user)

var user User
db.First(&user, "name = ?", "John")

result := db.Model(&user).Update("Age", 30)
if result.Error != nil {
    // エラーハンドリング
    fmt.Println(result.Error)
    return
}
fmt.Println(user)
```

このコードでは、まず名前が"John"の最初のUserを取得します。次に、UserのAgeフィールドを更新し、Saveメソッドを使用してこの変更をデータベースに保存します。

Updateメソッドを使用する場合：

Updateメソッドは、指定したフィールドのみを更新します。

このコードでは、まず名前が"John"の最初のUserを取得します。次に、Updateメソッドを使用してUserのAgeフィールドのみを更新します。

これらのメソッドはどちらも更新操作を行いますが、Saveメソッドは全フィールドを更新し、Updateメソッドは指定したフィールドのみを更新するという違いがあります。したがって、使用するメソッドは更新したいフィールドによります。

```go
var user User
db.First(&user, "name = ?", "John")

result := db.Delete(&user)
if result.Error != nil {
    // エラーハンドリング
    fmt.Println(result.Error)
    return
}
fmt.Println("User deleted successfully")
```

## Eager Loadingとは

Eager Loadingの概念とその利点を説明します。

Eager Loadingは、ORM（Object-Relational Mapping）におけるデータ取得戦略の一つで、あるエンティティとそれに関連するエンティティを一度のクエリで取得する方法を指します。

例えば、UserとOrderのような2つの関連するエンティティがあるとします。UserごとにそのOrderを取得するためには、通常、各Userに対して別々のクエリを発行する必要があります。これはN+1問題として知られており、パフォーマンスの問題を引き起こす可能性があります。

Eager Loadingを使用すると、一度のクエリでUserとそれに関連するすべてのOrderを一度に取得できます。これにより、データベースへのクエリ数が大幅に減少し、アプリケーションのパフォーマンスが向上します。

ただし、Eager Loadingは必要以上に多くのデータを取得する可能性があるため、使用する際には注意が必要です。必要なデータだけを取得するために、適切なクエリ戦略を選択することが重要です。

## GormでのEager Loadingの使用

Gormを使ったEager Loadingの例を示します。

```go
type User struct {
    gorm.Model
    Name   string
    Orders []Order
}

type Order struct {
    gorm.Model
    UserID uint
    Price  float64
}

var user User
db.Preload("Orders").First(&user, "name = ?", "John")
for _, order := range user.Orders {
    fmt.Println(order)
}

```

このコードでは、Preloadメソッドを使用してOrders関連をEager Loadingします。これにより、名前が"John"のUserとそのすべてのOrderを一度のクエリで取得します。

このように、GormのPreloadメソッドを使用すると、関連するエンティティを効率的に取得することができます。

# まとめ

本記事ではORMの概要とGoのORMであるGormを使った基本操作を解説しました。

バックエンドアプリケーションを構築する際にDBの操作の一つとして知っておきたい技術だと考えています。

ぜひ参考にしてみなさんのプロジェクトに役立ててください！
