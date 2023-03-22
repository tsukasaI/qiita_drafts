# qiita_drafts

## future

- Github Actionsで過去のコミットを参照しようとしたらunknown revision or path not in the working tree と言われたときの解決法

actions/checkout@v3 は headのみ取ってくる
過去のものを全て取得する場合はfetch-depth = 0 とする
参考
https://github.com/actions/checkout#fetch-all-history-for-all-tags-and-branches
