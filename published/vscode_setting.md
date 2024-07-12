# VSCodeの設定どうなってる？

VSCodeユーザーの皆さんは設定はどのようにしていますか？

Macであれば`Cmd + ,`で、Windowsであれば`Ctrl + ,`で設定画面を開くことができます。

しかしVSCodeにはリッチな機能を有するため設定ファイルには大量の項目が存在します。

プライベートでも開発をする（強要はしていない）エンジニアの皆さんは業務用のPCもプライベートのPCも使っていて設定を共有したい時もありますよね。

そんなときに便利に共通の設定を使えるようにするのがこの記事の目的になります。

## 設定ファイルの場所

VSCodeの設定ファイルは、`settings.json`というファイルに記述されています。

このファイルは、どのワークスペースにも適用されるファイル（ユーザー）と各ワークスペース設定の両方をが存在します。

優先順位はユーザーよりもワークスペースの設定が優先されます。

### ユーザー設定

ユーザーの設定は`Cmd + Shift + P`でコマンドパレットを開き、`Preferences: Open Settings (JSON)`を選択すると開くことができます。

jsonなので各項目がkey-valueで記述されています。

ちなみに筆者の設定は以下の感じになっている。

```json
{
  "editor.minimap.enabled": false,
  "editor.wordSeparators": "`~!@#$%^&*()-=+[{]}\\|;:'\",.<>/?　、。！？「」【】『』（）",
  "swaggerViewer.previewInBrowser": true,
  "files.trimFinalNewlines": true,
  "files.trimTrailingWhitespace": true,
  "files.insertFinalNewline": true,
  "terminal.integrated.tabs.enabled": false,
  "[typescriptreact]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "[html]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "[json]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "go.toolsManagement.autoUpdate": true
}
```

こだわっている部分はさておいて、いくつか皆さんにも適用しておくと便利な設定を紹介します。

```
  "files.trimFinalNewlines": true,
  "files.trimTrailingWhitespace": true,
  "files.insertFinalNewline": true,
```

ここで設定しておくと、ファイル保存時に自動で改行や空白を削除してくれます。

ファイルの最後に改行がないとGitの更新時に差分として表れるので、変更したいファイル以外にも変更があるように見えてしまいます。

### ワークスペース設定

ワークスペース設定は、ワークスペースの`.vscode`ディレクトリ内に`settings.json`というファイルを作成することで設定を行うことができます。

ここでは例えば、プロジェクトごとに異なる設定を行いたい場合に使うことができます。

例えば、プロジェクトごとに異なるフォーマッタを使いたい場合などに使うことができます。

```json
{
  "editor.defaultFormatter": "esbenp.prettier-vscode"
}
```

このように設定することで、このワークスペース内でPrettierを使うことができます。

ついでに.vscodeディレクトリに使うextensionsを記述した`extensions.json`を作成しておくと、ワークスペースを開いたときに自動で拡張機能をインストールしてくれます。

```json
{
  "recommendations": [
    "esbenp.prettier-vscode"
  ]
}
```
## まとめ

統一した設定を使うことで、開発効率が上がりますので個人/チームで設定を共有することをおすすめします。
