# はじめに

普段はBEをメインで開発している筆者がReactのstateの仕様を知らずに2日を溶かしたため

この記事を公開します。

# TL;DR

useStateで定義したset関数をコールしても即座にはstateに反映されない。

更新されるタイミングは再レンダリングされたあとに反映される。

set関数をコールした直後にstateを使う場合はuseEffectを使うと値を得られる。

# やりたかったこと

タイトルの通りにstateが更新されたら別の関数で新しいstateを用いて処理を行ないたい。

今回はinputに入力されたタイミングでコンソールに出力されるようなプログラムを書いてみます。

## NGパターン

set関数をコールした後にstateを直接参照するパターンを考えてみます。

```tsx
import React, { useState } from 'react'

const Sample: React.FC = () => {
  const [searchText, setSearchText] = useState<string>('')

  const handleOnChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setSearchText(e.target.value)
    console.log(searchText)
  }

  return (
    <div>
      <input
        type="text"
        value={searchText}
        onChange={handleOnChange}
      />
    </div>
  )
}
```

このように書くと入力した瞬間には「最後に入力された文字が出力されない」という問題があります。

では反映されたstateを見るためにはどうしたら良いでしょうか？

## Goodパターン

```tsx
import React, { useEffect, useState } from 'react'

const Sample: React.FC = () => {
  const [searchText, setSearchText] = useState<string>('')

  const handleOnChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setSearchText(e.target.value)
  }

  useEffect(() => {
    console.log(searchText)
  }, [searchText])

  return (
    <div>
      <input
        type="text"
        value={searchText}
        onChange={handleOnChange}
      />
    </div>
  )
}
```

このようにuseEffectの依存配列にstateを指定すると更新後の値が得られます。
