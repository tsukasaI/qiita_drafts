# Dockerコマンドでコンテナ、イメージ、ボリューム、ネットワーク、キャッシュをまとめて削除する

Dockerを用いた開発をすると多くのイメージやコンテナ、ボリューム、ネットワーク、キャッシュが残りがちです。

これらを削除するコマンドたちを紹介します。

## コンテナの削除

コンテナの削除は `docker rm <container_id>` で行う。

引数には複数のコンテナIDを指定することができ、全てのコンテナを削除するには `docker rm $(docker ps -a -q)` とする。

## イメージの削除

一方でイメージの削除は `docker rmi <image_id>` で行う。

こちらも引数には複数のイメージIDを指定することができ、全てのイメージを削除するには `docker rmi $(docker images -q)` とする。

## ボリュームの削除

ボリュームの削除は `docker volume rm <volume_name>` で行う。

ほぼ同じことを書いているが、引数には複数のボリューム名を指定することができ、全てのボリュームを削除するには `docker volume rm $(docker volume ls -q)` とする。

以上のようにDockerの各オブジェクトに対して削除コマンドが用意されている。

が使っては破棄が可能なためこれらのオブジェクトがどんどん溜まっていってはいないだろうか？

（筆者は溜まっている）

そのためまとめて削除するコマンドを紹介する。

## prune

`docker system prune` で全てのコンテナ、イメージ、ボリューム、ネットワークを削除することができる。

https://docs.docker.jp/config/pruning.html#id5

コマンドを実行すると以下のように確認を求められる。

```sh
$ docker system prune

WARNING! This will remove:
        - all stopped containers
        - all networks not used by at least one container
        - all dangling images
        - all build cache
Are you sure you want to continue? [y/N]
```

ここでは停止中のコンテナ、コンテナに紐づいていないネットワーク、未使用のイメージ、ビルドキャッシュを削除するけど大丈夫？と聞かれる。

ボリュームは上のコマンドでは削除されないので、ボリュームも含めて削除したい場合は `docker system prune --volumes` とする。

```sh
$ docker system prune --volumes

        - all stopped containers
        - all networks not used by at least one container
        - all volumes not used by at least one container
        - all dangling images
        - all build cache
Are you sure you want to continue? [y/N]
```

ボリュームを含めてまとめて削除ができる。
