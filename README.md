# Comic Review Sample
Next.jsを使用したフロントエンドと連携する、echoフレームワークのサンプルプロジェクトです。当該プロジェクトの技術および使用ツールは下記になります。
- REST API
- GORM(O/Rマッパ)
- Docker
- GitHub Actions
- クリーンアーキテクチャ
- Amazon S3
- Render

## ローカル利用方法
このプロジェクトを実行するには、`.env` ファイルに下記パラメーターを設定する必要があります。

```plaintext
# Dockerを使用する場合のPostgreSQLの設定
POSTGRES_USER = portfolio
POSTGRES_PW = portfolio
POSTGRES_DB = portfolio
POSTGRES_PORT = 5434
POSTGRES_HOST = localhost

# フロントエンドのURLを指定します。
FE_URL = http://localhost:80

# 下記パラメーターはAWS S3の設定に必要です。
AWS_BUCKETNAME = YOUR_AWS_BUCKETNAME
AWS_REGION = YOUR_AWS_REGION
AWS_ACCESS_KEY = YOUR_AWS_ACCESS_KEY
AWS_SECRET_ACCESS_KEY = YOUR_AWS_SECRET_ACCESS_KEY
```

楽天ブックAPIのキー[^1]
[^1]: [https://webservice.rakuten.co.jp/documentation/kobo-ebook-search](https://webservice.rakuten.co.jp/documentation/kobo-ebook-search)
```
APPLICATION_ID = YOUR_APPLICATION_ID
```

下記コマンドを利用して、ライブラリ及びモジュールをダウンロードしてください。
```bash
# PostgreSQLコンテナの起動
$ docker compose up -d

# データベースのマイグレーション実行
$ go run migrate/migrate.go

# 依存関係のダウンロード及びアプリケーションの実行
$ GO_ENV=dev go run .
```

## 注意事項
- Mac OS Monterey 12.7.1でのみ動作確認しています。