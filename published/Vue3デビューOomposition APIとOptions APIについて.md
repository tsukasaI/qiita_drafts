# はじめに

Vue2 の経験がある著者が Vue3 の書き方で大量のはてなマークを量産したためまとめる。

Vue2 の書き方と Vue3 の書き方の差について理解の一助となれば幸いです。

# Options API について

まずは Vue2 の書き方の Options API のサンプルを記述してみます。

カウントアップのアプリの script タグ内を考えてみましょう。

```vue
<script>
export default {
  data: () => {
    return {
      count: 0,
    };
  },

  methods: {
    countUp() {
      this.count++;
    },
    reset() {
      this.count = 0;
    },
  },
};
</script>
```

default 内では data メソッドでリアクティブな値、methods プロパティでリアクティブな値にアクセスするメソッドを記述します。

この問題としては **リアクティブな値にアクセスするためには this によるアクセスを必須となり、View と分離ができない** という点があります。

それ故に一つのコンポーネントファイルの巨大化、共通処理を各コンポーネントに記述するといった冗長化が発生しがちです。

# 救世主 Composition API について

そこで登場したのが Vue3 の書き方の Composition API です。サンプルを記述してみます。

```js:define.js
import { ref } from "vue";

export const useCounter = () => {
  const count = ref(0);

  const addCounter = () => {
    count.value++;
  };
  const resetCounter = () => {
    count.value = 0;
  };

  return {
    count,
    addCounter,
    resetCounter,
  };
};
```

上記のようにステートフルなロジックをカプセル化することを Composable と呼び、Composable は import を経由して呼び出しが可能です。

```vue:use.vue
<script>
  import { ref } from "difine";

  const { count, addCounter, resetCounter } = useCounter()
</script>
```

呼び出す側では import してプロパティとメソッドにアクセス可能になります。

# Options API と Composition API の比較

## data

```js:optionsApi
data: () => {
  return {
    count: 0
  };
}
```

```js:compositionApi
import { ref } from "vue"

const count = ref(0);
```

Options API では data メソッドで定義していましたが、Composition API では ref（もしくは reactive）関数で定義します。

## methods

```js:optionsApi
methods: addCounter() {
  this.count++;
}
```

```js:compositionApi
const addCounter = () => {
  count.value++;
}
```

Options API では methods プロパティ内に定義していましたが、Composition API では通常の script ブロック内で関数定義することで呼び出しが可能になります。

リアクティブな値の変更は value プロパティを利用します。

# まとめ

Composition API を利用することでロジックの分離と再利用が容易になりました。

とはいえ銀色の弾丸ではなくこれまでの Options API の方が View とセットでわかりやすい場合もあります。

作りたいものからどちらを選ぶかは別途検討する必要があります。

（個人的には自由度の高いもので十分に抽象化されている設計をできると気持ち良いので Composition API を推します。）

# 参考ページ

[Composition API とは？Options API との違いを徹底解剖【Vue 公式パートナーが解説】](https://jp.code-dict.com/media/vue/5876/)
