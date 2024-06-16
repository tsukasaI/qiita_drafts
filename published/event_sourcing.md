# イベントソーシング

イベントソーシングとは、データベースに保存するデータをイベントの形式で保存することです。イベントソーシングを使うことで、データベースの状態をイベントの履歴から再現することができます。


CRUDのようにデータベースの状態を直接変更するのではなく、イベントを発行してデータベースの状態を変更します。イベントソーシングを使うことで、データベースの状態を変更する操作をイベントの履歴として保存することができます。


ステートソーシングとイベントソーシングの違いは、データベースの状態を直接変更するのではなく、イベントを発行してデータベースの状態を変更する点です。イベントソーシングを使うことで、データベースの状態を変更する操作をイベントの履歴として保存することができます。

ステートソーシングにおいて履歴をトレースするためには、データベースの状態を保存するためのテーブルを作成したりロギングを行って時系列で解析する必要があります

この特徴から、イベントソーシングは、データベースの状態を変更する操作をイベントの履歴として保存することができるため特定の時点でのデータベースの状態を再現することが容易です。


## イベントソーシングの利点

イベントソーシングの利点は以下の通りです。

- データベースの状態をイベントの履歴から再現することができる
- データベースの状態を変更する操作をイベントの履歴として保存することができる
- 更新履歴があるため特定のタイミングのステートを再現することが容易

## イベントソーシングのデメリット

イベントソーシングのデメリットは以下の通りです。

- データベースの状態を変更する操作をイベントの履歴として保存するため、データベースの状態を再現するためにはイベントの履歴を再生する必要がある
- イベントの履歴を再生するためには、イベントの履歴を保存するためのストレージが必要

特に最新の状況を取得するためには、イベントの履歴を再生する必要があるため、データベースの状態を再現するためにはイベントの履歴を再生する必要があります。

一方でステートソーシングでは最新の状態はデータベースに保存されているため、データベースの状態を再現するためにはデータベースの状態を取得するだけで済みます。

## イベントソーシングの実装

ここからGoを使ってイベントソーシングを実装する方法を説明します。

イベントソーシングを実装するためには、以下の手順が必要です。

1. イベントを発行する
2. イベントを保存する
3. イベントを再生する

イベントを発行するためには、イベントを発行するための関数を定義します。イベントを保存するためには、イベントを保存するためのストレージを定義します。イベントを再生するためには、イベントを再生するための関数を定義します。

以下にイベントソーシングを実装するためのGoのコードを示します。

```go

package main

import (
    "encoding/json"
    "fmt"
    "log"
    "os"
)

type Event struct {
    Type string `json:"type"`
    Data string `json:"data"`
}

func main() {
    // イベントを発行する
    event := Event{
        Type: "user_created",
        Data: "user_id: 1, user_name: alice",
    }

    // イベントを保存する
    saveEvent(event)

    // イベントを再生する
    replayEvents()
}

func saveEvent(event Event) {
    file, err := os.OpenFile("events.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    if err := encoder.Encode(event); err != nil {
        log.Fatal(err)
    }
}

func replayEvents() {
    file, err := os.Open("events.json")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    for {
        var event Event
        if err := decoder.Decode(&event); err != nil {
            break
        }
        fmt.Printf("replay event: %+v\n", event)
    }
}

```

このコードでは、`Event`構造体を定義し、`main`関数でイベントを発行して保存し、再生する処理を行っています。

イベントを発行のためには、`Event`構造体を使ってイベントを作成し、`saveEvent`関数でイベントを保存します。イベントを保存するためには、`os.OpenFile`関数を使ってファイルを開き、`json.NewEncoder`を使ってイベントをエンコードしてファイルに書き込みます。

イベントを再生のためには、`replayEvents`関数でファイルを開き、`json.NewDecoder`を使ってイベントをデコードして再生します。

このように、イベントソーシングを実装することで、データベースの状態をイベントの履歴から再現することができます。