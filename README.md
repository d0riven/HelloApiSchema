## 概要

以下のAPISchemaを手を動かしてみて調査・検証するためのレポジトリ

* JSON + OpenAPI3(Swagger)
  - 参考: https://qiita.com/marnie_ms4/items/4582a1a0db363fe246f3
  - 参考: https://techblog.zozo.com/entry/openapi3/go
* gRPC + Protocol Buffers
  - 参考: https://qiita.com/marnie_ms4/items/4582a1a0db363fe246f3
* GraphQL

## 準備


## API

単一のユーザ情報のみを扱う。  
この検証の本質ではないので、認証等のセキュリティや個人情報の分割は一切気にしない。

### ユーザ

#### 要素

+ ユーザID:int
+ メールアドレス:string
+ 姓:string
+ 名:string
+ 生年月日:Date
+ 住所:string

#### API

+ 取得
  - GET
+ 更新
  - PUT
+ 登録
  - POST

※ 本当はGraphQLの長所を確認するならもっと複数の読み込み系のAPIがあればいいのだろうけど複数用意するので面倒なので妥協。