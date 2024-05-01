残り
SQLとの直接的な対話とORMを使用するときの違いを説明。
リレーションPreload

# タイトル: ORM入門 Go言語 Gorm パッケージ 解説

はじめに

IGD LXU所属の井上です。今回は我々のプロジェクトで利用している技術の一つであるORMを紹介します。

## ORMとは

ORM（Object-Relational Mapping）は、プログラムのオブジェクトとリレーショナルデータベースのデータをマッピングする技術です。

ORMにより、データベースの操作をプログラミング言語のオブジェクトを操作するような形で行うことができます。

例えばMySQLのあるテーブルの1レコードをGoのある構造体の実態としてマッピングして構造体で扱うことができるようになります。

ORMの主なメリットは以下の通りです：

- 生のSQLを書く必要がない：ORMはデータベース操作を抽象化し、開発者がSQLクエリを直接書く必要を減らす。これによってコードの可読性と保守性が向上する。

- データベースの変更が容易：ORMを使用すると、使用しているデータベースの種類を変更するのが容易になる。ORMがデータベースの詳細を抽象化しているため、データベース固有のSQLを書く必要がなく、データベースの種類を変更してもコードの大部分を変更する必要がなくなる。

- コードの再利用性と組織化：ORMはモデルベースのアプローチを採用しておりコードの再利用性と組織化が向上する。モデルはビジネスロジックをカプセル化し、コードの整理と再利用が容易になる。

- 開発速度の向上：ORMはデータベース操作を大幅に簡素化し、開発者がより高速にアプリケーションを開発できるようになる。

これらの利点により、ORMはWebアプリケーションの開発などにおいて広く使用されARISEのプロジェクトでも採用されています。

バックエンドエンジニアならぜひ覚えておきたい技術であるORMをGo言語（以下Go）で説明をしていきます。

## GoとGorm

ここからはORMの操作をGoで利用可能なORMパッケージの一つであるGormで紹介をします。

## Gormを使ったデータベース操作

### セットアップ

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

ここで定義したdb変数を使ってこの後に紹介するデータベース操作を行います。

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

上記の構造体では以下のDDLで定義されるテーブルを期待します。

```sql
CREATE TABLE `users` (
  `id` bigint unsigned AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `email` varchar(255),
  `age` tinyint unsigned NOT NULL,
  `birthday` datetime(6),
  `member_number` varchar(255),
  `activated_at` datetime(6),
  `created_at` datetime(6) NOT NULL,
  `updated_at` datetime(6) NOT NULL,
  PRIMARY KEY (`id`)
);
```

次からはGormを使った基本的なデータベース操作の例を示します。

### Select操作

Gormを使ってデータを取得する方法を説明します。

GormではFirstとFindメソッドを使用してデータを取得できます。

```go
var user User
result := db.First(&user, "name = ?", "John")
// SELECT * FROM users WHERE name = 'John' LIMIT 1;
if result.Error != nil {
    // エラーハンドリング
    fmt.Println(result.Error)
    return
}
fmt.Println(user)

var users []User
result := db.Find(&users, "age >= ?", 18)
// SELECT * FROM users WHERE age >= 18;
if result.Error != nil {
    // エラーハンドリング
    fmt.Println(result.Error)
    return
}
for _, user := range users {
    fmt.Println(user)
}
```

Gormのメソッドの下には実際に発行されるSQLをコメントに記載しています。

Firstメソッドは、指定した条件に一致する最初のレコードを取得するためこのコードでは、名前が"John"の最初のUserを1件取得します。

一方でFindメソッドは、指定した条件に一致するすべてのレコードを取得します。

このコードでは、年齢が18以上のすべてのUserを取得します。

### Insert操作

続いてGormを使って新しいレコードを挿入する方法を説明します。

```go
user := User{Name: "John", Age: 25}
result := db.Create(&user)
// INSERT INTO users (name, age) VALUES ('John', 25);
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
// INSERT INTO users (name, age) VALUES ('John', 25), ('Jane', 30), ('Bob', 20);
if result.Error != nil {
    // エラーハンドリング
    fmt.Println(result.Error)
    return
}

for _, user := range users {
    fmt.Println(user)
}
```

GormのCreateメソッドを使用して新しいレコードを挿入でき、引数には挿入するデータ（単一の場合は構造体のインスタンスのポインタ、複数の場合はスライスのポインタ）を渡します。

### Update操作

続いてGormを使って既存のレコードを更新する方法を説明します。

GormではSaveとUpdateメソッドを使用してデータを更新できます。

```go
var user User
db.First(&user, "name = ?", "John")

user.Name = "John 2"
user.Age = 30
result := db.Save(&user)
// UPDATE users SET name = ?, age = 30 WHERE id = ?;
if result.Error != nil {
    // エラーハンドリング
    fmt.Println(result.Error)
    return
}

var user User
db.First(&user, "name = ?", "John")

result := db.Model(&user).Update("Age", 30)
// UPDATE users SET Age = 30 WHERE id = ?;
if result.Error != nil {
    // エラーハンドリング
    fmt.Println(result.Error)
    return
}
```

このコードでは、名前が"John"の最初のUserを取得し、次にUserのName, Ageフィールドをそれぞれ更新し、Saveメソッドを使用して全てのフィールドを更新します。

Updateメソッドは、指定したフィールドのみを更新します。

このコードでは、まず名前が"John"の最初のUserを取得します。次に、Updateメソッドを使用してUserのAgeフィールドのみを更新します。

これらのメソッドはどちらも更新操作を行いますが、Saveメソッドは全フィールドを更新し、Updateメソッドは指定したフィールドのみを更新するという違いがあります。

### Delete操作

最後にGormを使ってレコードを削除する方法を説明します。

```go
var user User
db.First(&user, "name = ?", "John")

result := db.Delete(&user)
// DELETE FROM users WHERE id = ?;
if result.Error != nil {
    // エラーハンドリング
    fmt.Println(result.Error)
    return
}
fmt.Println("User deleted successfully")
```

このコードでは、名前が"John"の最初のUserを取得し、Deleteメソッドを使用してそのUserを削除します。

ここまで基本の操作を紹介しました。

次からはORMで性能問題になりやすいN+1問題やEager Loadingについて説明します。

## 性能問題とEager Loadingとは

例えば、UserとOrderのような2つの関連するエンティティがあるとして、UserごとにそのOrderを取得するためには、通常、各Userに対して別々のクエリを発行する必要があります。これはN+1問題として知られており、性能問題を引き起こす可能性があります。

Eager Loadingは、ORMにおけるデータ取得戦略の一つで、あるエンティティとそれに関連するエンティティをまとめて一定のクエリで取得する方法を指します。

Eager Loadingを使用すると、一度のクエリでUserとそれに関連するすべてのOrderを一度に取得できます。これにより、データベースへのクエリ数が大幅に減少し、アプリケーションのパフォーマンスが向上します。

## GormでのEager Loadingの使用

Gormを使ったEager Loadingの例を示します。

以下のstructで定義されるUserとOrderの関連を考えます。

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
```

Eager Loadingを使用しない場合、UserごとにOrderを取得するためには、次のようにN+1問題が発生します。

```go
var users []User
db.Find(&users)
// SELECT * FROM users;
for _, user := range users {
	db.Model(&user).Related(&user.Orders)
	// SELECT * FROM orders WHERE user_id = ?;
}
```

このコードでは、まずすべてのUserを取得し、次に各Userに対してOrdersを取得するために別々のクエリを発行しています。

一方で、Eager Loadingを使用すると、Userとそれに関連するOrderを一度に取得できます。

```go
var user User
db.Preload("Orders").First(&user, "name = ?", "John")
// SELECT * FROM users WHERE name = 'John' LIMIT 1;
// SELECT * FROM orders WHERE user_id = ?;
fmt.Println(user)

```

このコードでは、Preloadメソッドを使用してOrders関連をEager Loadingします。これにより、名前が"John"のUserとそのすべてのOrderを2つのクエリで取得します。

このように、GormのPreloadメソッドを使用すると、関連するエンティティを効率的に取得することができます。

ただし、Eager Loadingは必要以上に多くのデータを取得する可能性があるため注意が必要です。必要なデータだけを取得するために適切なクエリ戦略を選択することが重要となります。

# まとめ

本記事ではORMの概要とGoのORMであるGormを使った基本操作を解説しました。

バックエンドアプリケーションを構築する際にDBの操作の一つとして知っておきたい技術だと考えています。

ぜひ参考にしてみなさんのプロジェクトに役立ててください！
