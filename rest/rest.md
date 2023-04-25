# REST API

## ◼️ REST APIとは？

REST APIは、HTTPという通信プロトコルを使用し、Webアプリケーション間でデータの送受信を行う技術。

REST APIを使うことで、Webアプリケーションが別のアプリケーションとやり取りすることができる。
例えば、Twitterアプリで新しいツイートを投稿したり、いいねしたり、フォローする場合、TwitterのWebサーバーにデータを送信する。APIを使用することで、やり取りを効率的かつスムーズに行うことができる。

送信側から受信したHTTPメソッドに応じ、受信側は自身のリソースに対して作成、更新、削除などの動作を行います。このようにリソースに対して行う主要な動作を`CRUD`という。

| CRUD    | HTTPメソッド | 説明     |
|---------|-------------|----------|
| Create  | POST / PUT  | 作成     |
| Read    | GET         | 読み取り |
| Update  | PUT / PATCH | 更新     |
| Delete  | DELETE      | 削除     |

<br>

- クライアント/サーバプロトコル通信
  - HTTPで定義するGET、POST、PUT、DELETEなどのリクエストでデータを操作する。
  - 送信側から操作方法をHTTPメソッドで表現して、受信側は処理結果をステータスコードで表現しておりクライアント/サーバ形態の通信を行う。

- リソースを一意なURIにより識別される
  - RESTfulなシステムでは全てのリソース（情報）はURI（Uniform Resource Identifier）で表現されるユニークなアドレスを持つ。
  - https://www.example.com/test
    - https → 使用するプロトコル
    - www.example.com → ホスト名(IPアドレス)
    - test → ホスト内のパス

- 汎用的なデータ形式に変換し通信するデータ形式
  - JSON（JavaScript Object Notation）- 拡張子（.json）
  - XML（Extensible Markup Language）- 拡張子（.xml）
  - YAML（YAML Ain't a Markup Language）- 拡張子（.yaml .yml）

<br>
<hr>

## ◼️ movieをリソースとして 

CRUD操作のURI、HTTPメソッドを定義

・https://myapp

| 操作 | URI          | HTTP method |
|------|--------------|-------------|
| Create | /movie      | POST        |
| Read   | /movie      | GET         |
| Read   | /movie/{id} | GET         |
| Update | /movie/{id} | PUT         |
| Delete | /movie/{id} | DELETE      |
