# 標準出力と標準エラー出力

Goにはbuildinで定義されるprint, println関数がある。

一方でfmtパッケージにもPrint, Println関数がある。

Goを学ぶとサンプルコードではfmt.Printlnで標準出力に出力することが多いが、print, println関数も使える。

この差について気になったので本記事でまとめる。

## 結論

print, printlnは標準エラー出力に、fmt.Print, fmt.Printlnは標準出力に出力される。

## 標準出力と標準エラー出力

標準出力とは、プログラムの出力を表示するために使われるチャネルのことを指す。

コンピュータのシステムやプログラムの情報の"正常"な出力を表示するために使われ、一般的にはディスプレイに表示される。

例えば`ls`コマンドなどはファイルやディレクトリの一覧情報などを標準出力に出力する。

一方で標準エラー出力は、プログラムのエラー情報を表示するために使われるチャネルのことを指す。

例えば`ls`コマンドで存在しないディレクトリを指定するとエラー情報が標準エラー出力に出力される。

どちらも出力先はディスプレイであるが、標準出力は正常な情報、標準エラー出力はエラー情報を表示するために使われ、

標準エラー出力はログに残すなどしてエラー情報のみを残すことができる。

ちなみにLinuxでは標準出力は`1`、標準エラー出力は`2`で表される。

コマンドのリダイレクトを使って `>` で標準出力、 `2>` で標準エラー出力をファイルに出力することができる。


例えば、`ls`コマンドの標準出力を`out.txt`に、標準エラー出力を`err.txt`に出力する場合は以下のようにする。
```shell
$ ls > out.txt 2> err.txt
```
## 最後に

正常なプリントをする場合であればfmtパッケージを使えばよさそう。