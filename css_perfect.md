CSS完全に理解したをUIライブラリで実装

# CSS完全に理解した

かつてこんなワードをこしらえたTシャツが一世を風靡した。

筆者も最初に見た時にはお腹を抱えて6時間は笑ったし、なんなら今見ても普通に吹き出す。

それくらい多くのエンジニアを笑顔にした(?)デザインがある。

そんな素敵なデザインをプロジェクトで利用しているCSSフレームワークを用いて再現をしてみたという毒にも薬にもならないコードを書いてみた。

普段はGoでバックエンド開発を生業としている筆者が自在にデザインを構築できるのかを生暖かく見守ってほしい。

## デザインの確認

「CSS完全に理解した」で検索してもらうとわかると思うがこんなイメージだ。

[例の画像]

CSSを完全に理解されていますね。

これを再現するに当たってよく見て、必要な要素を書いてみます。

- 実践のボーダーの枠
- 枠の角は丸くする
- 「CSS」と「完全に理解した」の二行の文字
- **左側の空白**
- **二行目を枠からはみ出させる**

こんな感じだろうか。

## 再現方針

まずは上で書いた要素をどうするか考えよう。

筆者ならこんな感じにする。

- divでborderをsolidにする
- border-radiusを設定する
- pタグで二行の文字を入れる
- **position: relative/absolute**で左に空白を入れる
- 改行を阻止するために**white-space: nowrap**を設定する

こんな方針でやってみる。

## 使用するCSS Framework

今回は[Tailwind CSS](https://tailwindcss.com/)を使って実装する。

Tailwind CSSは個人的に一押ししたいCSSライブラリで以下の特徴がある。

- Utility Classを使ってスタイリングを当てるためクラス名を考えなくていい
- デザインの自由度が高い
- 利用しないUtility Classは本番の環境から消すことができてファイルを軽量にできる

これらから多くのプロジェクトで利用されている

## 実装

それではTailwind CSSで方針の通りの記述をしたコードを見てみよう

```html
<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <script src="https://cdn.tailwindcss.com"></script>
  </head>
  <body>
    <div class="m-4 m-4">
      <div class="border-2 border-black rounded p-2 w-44 h-24 relative">
        <div class="absolute left-12 top-2 whitespace-nowrap">
          <p class="text-2xl font-bold">CSS</p>
          <p class="text-2xl font-bold">完全に理解した</p>
        </div>
      </div>
    </div>
  </body>
</html>
```

これでブラウザに表示するとこんな感じになった。

大枠は再現できたのではないか。
