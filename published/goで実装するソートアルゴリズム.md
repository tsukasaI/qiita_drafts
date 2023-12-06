goで実装する基本ソートアルゴリズム

## アルゴリズムとは

アルゴリズムはある目的を達成するための決まった操作のことです。

例えば目的を任意の3つの整数の中で最大のものを取得するとした場合、以下の操作で目的が達成できます。

1. 各整数を任意の順に並べる
1. １つ目と２つ目の整数の大小を比較して大きい方を取り出す
1. 取り出した整数と３つ目の整数の第上を比較して大きい方を取り出す

これも立派なアルゴリズムです。

上記の整数の個数が10でも100でも3の操作を整数の個数分繰り返して行くと最大が得られます。

## ソートアルゴリズム

アルゴリズムの中でデータを一定の規則に従って並べるアルゴリズムです。

複数の整数がランダムに並んだ状態からある順番に並べ直す

先ほど紹介した最大を取得するを少し発展させて昇順に並べ替えることを目的としてGoのコードを見ていきます。

## ソートアルゴリズムの種類

今回はすべて昇順になるように並べます

### バブルソート

最も基本的なソートアルゴリズムで、隣接する要素を比較して順序が間違っていれば交換するという操作をします。これを繰り返すことで、最大の要素がリストの一方の端に"浮かび上がる"ことからこの名前がついています。

Goのソースコードは以下のようになります。

```go
func bubbleSort(numbers []int) {
	for i := 0; i < len(numbers)-1; i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
}
```

バブルソートは他の配列などのメモリが必要がないin-placeアルゴリズムで、計算量は最悪/平均ともにO(n ^ 2)です。

### 選択ソート

リストの中から最小の要素を見つけ、それをリストの一方の端に移動します。移動済みの要素はソート済みとしてソートしていないリストに対してこれを繰り返すことで、リスト全体がソートされます。

Goのソースコードは以下のようになります。

```go
func selectionSort(numbers []int) {
	lenNumbers := len(numbers)
	for start := 0; start < lenNumbers; start++ {
		tmpIndex := start
		for i := start; i < lenNumbers; i++ {
			if numbers[tmpIndex] > numbers[i] {
				tmpIndex = i
			}
		}
		if tmpIndex != start {
			numbers[tmpIndex], numbers[start] = numbers[start], numbers[tmpIndex]
		}
	}
}
```

選択ソートはin-placeアルゴリズムで、計算量は最悪/平均ともにO(n ^ 2)です。

### 挿入ソート

リストをソート済み部分と未ソート部分に分け、未ソート部分の要素を一つずつソート済み部分の適切な位置に挿入していきます。

Goのソースコードは以下のようになります。

```go
func insertionSort(numbers []int) {
	lenNumbers := len(numbers)
	for i := 1; i < lenNumbers; i++ {
		temp := numbers[i]
		j := i - 1
		for j >= 0 && numbers[j] > temp {
			numbers[j+1] = numbers[j]
			j--
		}
		numbers[j+1] = temp
	}
}
```

挿入ソートはin-placeアルゴリズムで、計算量は最悪/平均ともにO(n ^ 2)です。

## 動作チェック

[Playground](https://go.dev/play/p/KnSMaJH-kZZ)にソースを置いておきました。

興味ある方は動かしてみてください。
