# Drizzle ORMの紹介

TypeScriptでORMを使うときに何を使いますか？

本記事ではDrizzle ORMの紹介をしていきます。

## 特徴

> Using Drizzle you can define & manage database schemas in typescript, access your data in a SQL-like or relational way, and take advantage of opt in tools to push your developer experience through the roof

Drizzleの公式ページの言葉を見るとこのように書かれています。

TypeScriptで型とスキーマの定義が可能で、SQLとリレーショナル（他のORMでよくある方法？）な手法でデータにアクセスできる。

早速次からどのように書くか見てましょう。

## サンプル

今回はPostgreSQLを使った例を示します。

### スキーマ定義

```typescript
import { integer, pgEnum, pgTable, serial, uniqueIndex, varchar } from 'drizzle-orm/pg-core';

// declaring enum in database
export const popularityEnum = pgEnum('popularity', ['unknown', 'known', 'popular']);

export const countries = pgTable('countries', {
  id: serial('id').primaryKey(),
  name: varchar('name', { length: 256 }),
}, (countries) => {
  return {
    nameIndex: uniqueIndex('name_idx').on(countries.name),
  }
});

export const countriesRelations = relations(countries, ({ many }) => ({
	cities: many(cities),
}));

export const cities = pgTable('cities', {
  id: serial('id').primaryKey(),
  name: varchar('name', { length: 256 }),
  countryId: integer('country_id').references(() => countries.id),
  popularity: popularityEnum('popularity'),
});
```

pgTableという関数を使ってテーブル名とテーブル定義をオブジェクトで受け取っていますね。ここはTypeScriptみを感じますね。

countriesRelationsについてはリレーショナルな書き方をする際に登場します。

外部キーを指定する場合はオブジェクトのvalue内に `references` メソッドで定義が可能です。

またインデックスキーは第三引数でコールバック関数として返すようです。

このテーブル定義のselectで得られる型については以下で抽出可能。

```typescript
xport type User = typeof countries.$inferSelect;
```

### データ取得

冒頭に書いたSQLライクな書き方と、リレーショナルな書き方をやってみます。

まずはSQLライクな書き方

```typescript
import { eq } from 'drizzle-orm';

const result = await db.select().from(countries).where(eq(countries.id, 42));
```

見たままSQLですね。これはSQLに慣れたエンジニアには嬉しいですね。

一方でリレーショナルな書き方

```typescript
import * as schema from './schema';
import { drizzle } from 'drizzle-orm/...';

const db = drizzle(client, { schema });

const result = await db.query.countries.findMany({
	with: {
		cities: true
	},
});
```

こちらはDBコネクションのタイミングでDBのスキーマをschemaとして渡す必要があります。

実際にデータにアクセスするためには `db.query.{テーブル}.findMany(~)` と書くことでクエリを発行できます。

更に `with` が続いていますが、ここでスキーマ定義した `countriesRelations` が効果を発揮します。

この書き方をするとリレーション関係にあるレコードを取得することができます。

## おまけでマイグレーション

Drizzleはマイグレーション機能も提供しているのでおまけとして書きます。

設定ファイル

```typescript:drizzle.config.ts
import 'dotenv/config';
import type { Config } from 'drizzle-kit';

export default {
	schema: './src/schema.ts',
	out: './drizzle/migrations',
	driver: 'pg',
	dbCredentials: {
    host: process.env.DB_HOST,
    user: process.env.DB_USER,
    password: process.env.DB_PASSWORD,
    database: process.env.DB_NAME,
	},
} satisfies Config;
```


```src/db.ts
import { drizzle } from 'drizzle-orm/mysql2';
import mysql from 'mysql2/promise';
import * as schema from './schema';

export const connection = mysql.createConnection({
  host: process.env.DB_HOST,
  user: process.env.DB_USER,
  password: process.env.DB_PASSWORD,
  database: process.env.DB_NAME,
  multipleStatements: true,
});

export const db = drizzle(connection, { schema });
```

```src/migrate.ts
import 'dotenv/config';
import { migrate } from 'drizzle-orm/mysql2/migrator';
import { db, connection } from './db';

// This will run migrations on the database, skipping the ones already applied
await migrate(db, { migrationsFolder: './drizzle' });

// Don't forget to close the connection, otherwise the script will hang
await connection.end();
```

このsrc/migrate.tsをtscで実行するとマイグレーションが実行されます。
