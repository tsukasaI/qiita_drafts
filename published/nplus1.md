# N+1問題のヤバさととその解決策

この言葉を聞いてダメ！絶対！と感じる人、どんな問題なんだろ？と感じる人さまざまいると思います。

この記事ではN+1問題のヤバさと対応方法にも言及していきます。

## N+1とは

どんな問題になるかというと性能の問題になります。

例を見ながら説明していきます。

X(旧Twitter)を簡単にしたサービスに以下のようなusersとtweeetsテーブルがあるとします。

**users**

| id | name |
| -- | -- |
| 1 | John |
| 2 | Nick |
| 3 | Mary |
| 4 | Nina |
| 5 | Kevin |
| 6 | Kate |

**tweets**

| id | user_id | tweet |
| -- | -- | -- |
| 1 | 1 | Good morning |
| 2 | 1 | Hi |
| 3 | 1 | Hello |
| 4 | 1 | Wow |
| 5 | 1 | Great |
| 6 | 1 | Fine |
| ... | ... | ... |

このDBから各tweetとtweetしたユーザーをクライアントに返却する場合はどうするでしょうか？

### N+1問題を引き起こすパターン

tweetsテーブルからレコードを全て(Nレコードとする)取得してからループ処理で各user_idでusersテーブルから1件ずつレコードを取得

という処理考えたとします。確かにtweetsテーブルのuser_id情報からusersの情報を取得することはできます。

しかしクエリの発行数に着目してみるとどうでしょうか？

クエリは

- tweetsテーブルからレコードを全件取得 : 1回
- ループ処理で各user_idでusersテーブルから1件ずつレコードを取得 : N回

から合計N+1回発行します。

Xのようなサービスでは毎日数多くのTweetがされるので、性能問題が起きるのは明らかですね。

### じゃあどうするか

対応は2パターンあると考えられます。

- RDBのJoinを使う
- Eager Loadを使う

#### RDBのJoinを使う

DBにある程度精通した方ならJoinでテーブルを結合する手法を思いつくでしょう。

```sql
select * from tweets left join users on tweets.user_id = users.id;
```

とすると1クエリで全データを取ることができます。

#### Eager Loadを使う

Eager LoadほとんどのORMが持っている機能で解決することもできます。

Eager Loadを使う場合、処理とクエリ発行は以下の流れになります。

1. tweetsテーブルのレコードを取得: 1回
1. tweetsテーブルのレコードからuser_idをリスト化: 0回（ORM内部の処理）
1. 上で生成したuser_idリストにあるusersテーブルのレコードを取得: 1回
1. tweetsのuser_idとusersのidのリレーションからまとめた構造を作る: 0回（ORM内部の処理）

Eager LoadはGoのORMパッケージのGORMでこの機能を使えます。

例を書いてみます。

```go
// usersテーブルの構造体
type User struct {
    ID    uint `gorm:"primaryKey"`
    Name  string
}

// tweetsテーブルの構造体
type Tweet struct {
    ID     uint `gorm:"primaryKey"`
    UserID uint
    User   User
    Tweet  string
}

// tweetsからusersをEager Load
func getTweetsWithUser() []Tweet {
    tweets := make([]Tweet, 0)
    db.Preload("User").Find(&tweets)
    return tweets
}
```

このように`getTweetsWithUser()`をコールするとEager Loadが実行されTweet構造体にUser情報が格納され、かつクエリは2回のみになります。

皆さんは性能問題を引き起こすことのないようにこの対応を覚えておきましょう。
