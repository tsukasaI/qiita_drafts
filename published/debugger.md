# VSCodeのデバッガを動かす

プログラミングをしている時には不可欠なデバッギングですが、VSCodeのデバッガを使うことで効率的にデバッグを行うことができます。

いつもprintデバッグをしている人も、VSCodeのデバッガを使うことでデバッグの効率が上がるかもしれません。

（かく言う私もGoのprintで様々なデバッグをしてきたので自分に向けて書いた記事でもある）

次からはPythonを使ったデバッグの方法を書いていきます。

## 準備

extensionでPython Debugger、もしくはPythonをインストールします。

.vscode/launch.jsonに以下の記述をする。

（デフォルトの設定でこうなります）

```json:.vscode/launch.json
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Python Debugger: Current File",
            "type": "debugpy",
            "request": "launch",
            "program": "${file}",
            "console": "integratedTerminal"
        }
    ]
}
```

これでF5を押すことでデバッグを開始できるようになります。

## ブレークポイント

ブレークポイントを設定することで、その行でプログラムが止まります。

ブレークポイントを設定するには行番号の左側をクリックする、もしくはF9を押します。（赤い丸が表示されます）

デバッグを開始するとブレークポイントで止まります。

## デバッグしてみる

以下のコードを書いてデバッグしてみます。

```python:main.py
def append_abc(s: str) -> str:
    a = append_a(s)
    b = append_b(a)
    c = append_c(b)
    return c

def append_a(s: str) -> str:
    return s + 'a'

def append_b(s: str) -> str:
    return s + 'b'

def append_c(s: str) -> str:
    return s + 'c'


def main():
    s = 'start'
    a = append_a(s)
    b = append_b(s)
    c = append_c(s)
    abc = append_abc(s)

if __name__ == '__main__':
    main()
```

任意の場所にブレークポイントを設定してデバッグを開始します。

ブレークポイントを設定した行でプログラムが止まり、変数やコールスタックを確認することができます。

このときのサイドバーには以下のような情報が表示されます。

- Variables: 変数の値を確認できます。
- Watch: 変数の値を監視できます。
- Call Stack: コールスタックを確認できます。

## デバッグ中の操作

デバッグ中には以下の操作ができます。

| アクション | 説明 |
| --- | --- |
| 続行 / 一時停止 (F5) | 続行: 次のブレークポイントまで通常のプログラム/スクリプトの実行を再開する。 一時停止: 現在の行で実行中のコードを検査し、行ごとにデバッグする。 |
| ステップオーバー (F10) | 次のメソッドをそのコンポーネントのステップを検査せずに単一のコマンドとして実行する。 |
| ステップイン (F11) | 次のメソッドに入り、行ごとにその実行を追跡する。 |
| ステップアウト (⇧F11) | メソッドまたはサブルーチン内部にいる場合、そのメソッドの残りの行を1つのコマンドとして完了して、元の実行コンテキストに戻ります。 |
| 再起動 (⇧⌘F5) | 現在のプログラムの実行を終了し、現在の実行構成を使用してデバッグを再開する。 |
| 停止 (⇧F5) | 現在のプログラムの実行を終了する。 |

ここまででVSCodeのデバッガの使い方について書いてきました。

さらにブレークポイントに条件をつけたり可能なので[ドキュメント](https://code.visualstudio.com/docs/editor/debugging)を見ながら使ってみてください。
