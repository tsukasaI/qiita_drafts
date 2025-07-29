# Functional Options Patternでstructを初期化

Goにおいてstructを初期化するためにはいくつかの表現方法があります。

この記事ではいくつかのstructの初期化方法を紹介して、各パターンの特徴を紹介します。

## structのサンプル

まずは全ての設定値を関数の引数に渡すケースを考えてみます。

例としてDBへの接続情報を保持するstructを作ってみます。

```go
package db

type DB struct {
	Host     string
	User     string
	Password string
	Port     int
}

func New(host, user, password string, port int) *DB {
	return &DB{
		Host:     host,
		User:     user,
		Password: password,
		Port:     port,
	}
}
```

さて、このDB structに対して新たな設定値の `maxConn` を追加したい場合はどのようにするかをいくつかのパターンで実装してみます。

## 関数の引数に設定値を定義する

新しく関数を作成してmaxConnを設定できるようにしてみます。

```go
package db

type DB struct {
	Host     string
	User     string
	Password string
	Port     int
	MaxConn  int
}

func New(host, user, password string, port int) *DB {
	return &DB{
		Host:     host,
		User:     user,
		Password: password,
		Port:     port,
	}
}

func NewWithMaxConn(host, user, password string, port, maxConn int) *DB {
	return &DB{
		Host:     host,
		User:     user,
		Password: password,
		Port:     port,
		MaxConn:  maxConn,
	}
}
```

このようにMaxConnに対応する引数を追加してDB structをインスタンス化できるようにします。

これを別のパッケージから呼び出してインスタンス化する場合は以下のようになります。

```go
package main

import sample-project/db

func main() {
	simpleDB := db.New("localhost", "user", "password", 5432)
	DBWithMaxConn := db,NewWithMaxConn("localhost", "user", "password", 5432, 10)

	// ...
}
```

これを使う場合は設定値が増えるたびに引数を増やしたり、設定項目を選択するために関数を増やしたりする必要があるため拡張性に難がありそうですね。


## 設定値のstructを定義して引数に渡す

それでは設定を別のstructに定義してみます。

```go
type DB struct {
	cfg Config
}

type Config struct {
	Host     string
	User     string
	Password string
	Port     int
	MaxConn  int
}

func NewWithConfig(cfg Config) *DB {
	return &DB{
		cfg: cfg,
	}
}
```

呼び出し側
```go
package main

import sample-project/db

func main() {
	simpleDB := db.NewWithConfig(db.Config{
		Host: "localhost",
		User: "user",
		Password: "password",
		Port: 5432,
	})
	DBWithMaxConn := db,NewWithConfig(db.Config{
		Host: "localhost",
		User: "user",
		Password: "password",
		Port: 5432,
		MaxConn: 10,
	})
	// ...
}
```

これの利点は設定の個数に対して引数が固定になり、さらに変更があった場合も呼び出す関数を変更する必要がなくなる点です。

## 設定値を関数で定義できるようにする

さて、この記事で一番お伝えしたいパターンを書いてみます。

言葉で説明してみると各設定値を更新する関数を定義して、呼び出し側で必要な設定のみ使用するようなイメージです。

コードを見てみましょう。

```go
package db

type DB struct {
	Host     string
	User     string
	Password string
	Port     int
	MaxConn  int
}

type OptionFunc func(*DB)

func New(options ...OptionFunc) *DB {
	instance := new(DB)
	for _, optFn := range options {
		optFn(instance)
	}
	return instance
}

func WithHost(host string) OptionFunc {
	return func(d *DB) {
		d.Host = host
	}
}

func WithUser(user string) OptionFunc {
	return func(d *DB) {
		d.User = user
	}
}

func WithPassword(password string) OptionFunc {
	return func(d *DB) {
		d.Password = password
	}
}

func WithPort(port int) OptionFunc {
	return func(d *DB) {
		if port <= 0 || port > 65535 {
			// エラーハンドリング
			return
		}
		d.Port = port
	}
}

func WithMaxConn(maxConn int) OptionFunc {
	return func(d *DB) {
		d.MaxConn = maxConn
	}
}
```

呼び出し元を見てみましょう。

```go
package main

import sample-project/db

func main() {
	simpleDB := db.New(
		db.WithHost("localhost"),
		db.WithUser("user"),
		db.WithPassword("password"),
		db.WithPort(5432),
	})
	DBWithMaxConn := db,New(db.Config{
		db.WithHost("localhost"),
		db.WithUser("user"),
		db.WithPassword("password"),
		db.WithPort(5432),
		db.WithMaxConn(10),
	})
	// ...
}
```

一気に関数の数が増えました。

一方で記述量が多くなってはいるものの設定値の変更に対して関数のシグネチャの更新が必要なく変更に柔軟になることがメリットになっています。

変更が期待されていない初期化の関数に対しては記述量が増えるだけですが、拡張される構造体を初期化するには良い選択肢の1つだと思います。

## 参考
ref: https://golang.cafe/blog/golang-functional-options-pattern.html
