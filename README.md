# first-go-api

## About
GoでRest APIサーバーを作成する際のサンプルソースです。

## Tool
Language: Go 1.22  
Framework: Echo v4  
DB: MySQL 8.3.0  

## install to local
Dockerを使わずlocalで実行したい場合  
Go: https://go.dev/doc/install  
DB: https://dev.mysql.com/downloads/installer/  

## usage
起動
```sh
$ docker-compose build
$ docker-compose up -d
``` 

テストデータ_最新のものを取得
```sh
$ sh curl/get_test.sh
```

テストデータ_データ登録
```sh
$ sh curl/post_test.sh {登録したい文字列}
```