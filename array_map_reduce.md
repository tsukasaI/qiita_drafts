# Array の処理は可能な限り map, reduce, find を使ってほしい件

いきなりですがエンジニアの皆さん、配列の処理はどうやっていますか？

例えば以下のケースではどんなコード/関数/メソッドを使うか考えてみてください。

1. ある配列から同じ長さで別の配列を生成したい
1. ある配列のうち条件を満たす要素を一つ取得したい
1. ある配列のうち条件を満たす要素を複数取得したい
1. ある配列から一つの計算結果を出したい

今回の説明においては、言語は TypeScript/JavaScript を使います。

配列は `[1, 2, 3, 4, 5, 6, 7, 8, 9]` として

1. 配列の全ての要素を 2 倍して新しい配列を生成する
1. 配列で値が 4 である要素を取得する
1. 配列の要素のうち偶数のみの配列を生成する
1. 配列の要素の合計値を出す

で解説をします。

## 1~4 を全て for で実装

最もわかりやすい例は for 文を使うことでしょう。

for で 1~4 を実装した例を示します。

Playground にサンプルコードを載せていますので動かしてみたい方は[こちらから](https://www.typescriptlang.org/play?#code/MYewdgzgLgBAhgJwXAnjAvDA2gRgDQwBMBAzAQCwECsBAbAQOwEAcBAnALoDcAUD6JFgAzEAhwAlAKYQArgBsoALhhgZAWwBGkhFg4Zs3HiIQwAFAOgxJcyWslhYIIfCSoAlDADePGDGMTpeSgAOgAHGQgAC1NrW3tYACoiN14AX35wCBAbYLkQAHNTfylZBRS+HhthUUISoOVVTW19AFocXmMzC1hYuwcYJxdkFA9vXwBLZxibPth0eZhyUZ9fPxq6hX1e+N5VmA0ESTgAa12YdPSLbMlcgqL1wLLePm61hBINpRV1LR09TF0HVEXUyPRm8QGzkQw2WEym236AFIiBgFgAGWGrYwfR4hcJRaZxBzlPYHI6nFYXDKQa63QrYz7lSqSaoIcifBo-ZqYNFAkzmUFWcH9QbQ9xeFbGdm4mAAakwCKgaWpWRyeXpomlpSg5SAA)。

```typescript
const array = [1, 2, 3, 4, 5, 6, 7, 8, 9];

const for1Result: number[] = [];
for (const element of array) {
  for1Result.push(element * 2);
}
console.log(for1Result);

let for2Result: number = -1;
for (const element of array) {
  if (element === 4) {
    for2Result = element;
    break;
  }
}
console.log(for2Result);

const for3Result: number[] = [];
for (const element of array) {
  if (element % 2 === 0) {
    for3Result.push(element);
    break;
  }
}
console.log(for3Result);

let for4Result: number = 0;
for (const element of array) {
  for4Result += element;
}
console.log(for4Result);
```

一見は問題なく実装はできていますが、個人的には

- アウトプット用の定数/変数を定義して書き換えているのが関数型っぽくない
- for の中身を見ないと期待する結果がわからない

といった理由から直したい衝動に駆られます。

## じゃあどうするか

1~4 を叶えるメソッドが JavaScript にはあるので、それらを使いたい。

ざっくりまとめると

1. ある配列から同じ長さで別の配列を生成したい: **map**メソッドを使う
1. ある配列のうち条件を満たす要素を一つ取得したい: **find**メソッドを使う
1. ある配列のうち条件を満たす要素を複数取得したい: **filter**メソッドを使う
1. ある配列から一つの計算結果を出したい: **reduce**メソッドを使う

となる。公式の解説は[こちらのページから](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Global_Objects/Array)。

それぞれ解説していきます。

### ある配列から同じ長さで別の配列を生成したい: **map**メソッドを使う

```typescript
// 配列の全ての要素を2倍して新しい配列を生成する map
const mapResult = array.map((element) => element * 2);
console.log(mapResult);
```

外部にアウトプット用の定数を定義することなく、かつ map を使っているので array と同じ長さの別の配列が来ると考えられるので意味もわかりやすくなっています。

ついでに 1 行で書けてスッキリしていますね。

### ある配列のうち条件を満たす要素を一つ取得したい: **find**メソッドを使う

```typescript
// 配列で値が4である要素を取得する find
const findResult = array.find((element) => element === 4);
console.log(findResult);
```

こちらも外部にアウトプット用の定数を定義することなく、かつ filter を使っているので array から特定の値が来ると考えられるので意味もわかりやすくなっています。

### ある配列のうち条件を満たす要素を複数取得したい: **filter**メソッドを使う

```typescript
// 配列の要素のうち偶数のみの配列を生成する filter
const filterResult = array.filter((element) => element % 2 === 0);
console.log(filterResult);
```

ry

### ある配列から一つの計算結果を出したい: **reduce**メソッドを使う

```typescript
// 配列の要素の合計値を出す reduce
const reduceResult = array.reduce(
  (accumulator, currentValue) => accumulator + currentValue
);
console.log(reduceResult);
```

ry

## まとめ

JavaScriptに限らずPythonやRustにも同様です。

for一辺倒で片付けずにどの書き方が良いかを考えるきっかけにしてみてください！
