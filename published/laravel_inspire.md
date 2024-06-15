# Laravelで準備されているimspireコマンドが誰得過ぎた

PHPのWebアプリケーション向けのフレームワークのLaravelは便利な機能が多いですが、その中でも特に誰得だと感じたのが`php artisan inspire`コマンドです。

試しに実行してみると

```shell
$ php artisan inspire
I have not failed. I've just found 10,000 ways that won't work. - Thomas Edison
```

このコマンドがしていることをソースコードリードをして理解してみます。

（ソースコードリードの一例として生暖かい目で見守ってください）

## Laravelのコマンド

11.xのLaravelのソースコードにはroutes/console.phpに以下のように記述されています。

```php
use Illuminate\Foundation\Inspiring;
use Illuminate\Support\Facades\Artisan;

Artisan::command('inspire', function () {
    $this->comment(Inspiring::quote());
})->purpose('Display an inspiring quote');
```

https://github.com/laravel/laravel/blob/11.x/routes/console.php


このInspiringクラスはIlluminate\Foundation\Inspiringに定義されています。

次にInspiringクラスのファイルを見てみます。

```php
<?php

namespace Illuminate\Foundation;

use Illuminate\Support\Collection;

class Inspiring
{
    public static function quote()
    {
        return static::quotes()
            ->map(fn ($quote) => static::formatForConsole($quote))
            ->random();
    }

    public static function quotes()
    {
        return Collection::make([
            'Act only according to that maxim whereby you can, at the same time, will that it should become a universal law. - Immanuel Kant',
            ~~~
            'I have not failed. I\'ve just found 10,000 ways that won\'t work. - Thomas Edison',
            ~~~
        ]);
    }

    protected static function formatForConsole($quote)
    {
        [$text, $author] = str($quote)->explode('-');

        return sprintf(
            "\n  <options=bold>“ %s ”</>\n  <fg=gray>— %s</>\n",
            trim($text),
            trim($author),
        );
    }
}
```
（途中のコメントや一部の配列内の文字列は省略しています）

staticメソッドのquoteメソッドはquotesメソッドで返されるCollectionの中からランダムに一つの引用を返します。

https://github.com/laravel/framework/blob/11.x/src/Illuminate/Foundation/Inspiring.php

（ちなみにこのファイルにはアスキーアートが含まれていた。誰得）

## 最後に

Laravel開発者の方へ、たまにはこのコマンドを実行してインスパイアされてください。
