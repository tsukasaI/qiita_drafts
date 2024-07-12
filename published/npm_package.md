# NPMとpackage.jsonの基本

この記事ではNode.jsのパッケージマネージャであるNPM（Node Package Manager）と、NPMの設定ファイルである`package.json`について解説します。

## そもそもNPM is 何

ザックリいうとNPM（Node Package Manager）はNode.jsのパッケージ管理を担当してくれるツール。

例えばReactやVue、またサーバーサイドのExpressなどのライブラリやフレームワークをインストールする際に使います。

これらのライブラリやフレームワークは複数のファイルやディレクトリで構成されているため、それらを一括でダウンロードして管理してくれるパッケージ管理を行ってくれるのがNPMです。

Pythonではpip、Rubyではgem、PHPではcomposerなど、他の言語にもパッケージ管理ツールは存在するので他の言語で使ってた方はそんなものだと思ってもらえると。

## 基本的な使い方

### 初期化

初期構築は誰かが行ってくれていることが多いですが、自分で行う場合は以下のコマンドで初期化します。

```bash
$ npm init
```

このコマンドを実行すると、対話形式で`package.json`を作成するための情報を入力することができます。

### パッケージのインストール

package.jsonがあるディレクトリで以下のコマンドを実行すると、`package.json`に記載されているパッケージをインストールします。

```bash
$ npm install
```

一方で上記のコマンドではpackage.lock.jsonが生成されるため、実際には以下のコマンドを使うことが多いです。

```bash
$ npm ci
```

### パッケージの追加

プロジェクトに新たにパッケージを追加する場合は以下のコマンドを実行します。

```bash
$ npm install -save <package-name>
```

このコマンドでは商用環境で使うパッケージの場合は`-save`オプションをつけます。

テストやモックの作成、lintなどの開発環境でのみ使うパッケージの場合は`-save-dev`オプションをつけましょう。

```bash
$ npm install -save-dev <package-name>
```

このオプションを追加すると、`package.json`の`dependencies`ではなく`devDependencies`にパッケージが追加されます。

### パッケージのアップデート

パッケージのアップデートは以下のコマンドで行います。

```bash
$ npm update
```

### スクリプトの実行

`package.json`にスクリプトを記述しておくと、以下のコマンドでスクリプトを実行することができます。

```bash
$ npm run <script-name>
```

package.jsonに記載するスクリプトの例を以下に示します。

```json
{
  "scripts": {
    "start": "node index.js",
    "test": "jest"
  }
}
```

この場合、`npm run start`で`node index.js`が実行され、`npm run test`で`jest`が実行されます。

### スクリプトのちょっとしたテクニック

npmのスクリプトにはpost, preというプレフィックスがあります。

例えば、`start`スクリプトの前に`prestart`スクリプトを実行したい場合は以下のように記述します。

```json
{
  "scripts": {
    "prestart": "echo 'prestart'",
    "start": "node index.js"
  }
}
```

この場合、`npm run start`を実行すると、まず`prestart`スクリプトが実行され、その後`start`スクリプトが実行されます。

同様に、`post`プレフィックスを使うことで、スクリプトの後に処理を追加することもできます。
