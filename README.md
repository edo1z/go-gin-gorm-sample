# go-gin-gorm-sample

## Install 

```shell
git clone https://github.com/web3ten0/go-gin-gorm-sample.git
cd go-gin-gorm-sample/local
docker-compose up -d
```

## Config

```shell
cp config.json.example config.json
```

## Migration

- マイグレーションに下記を利用している。
  - https://github.com/golang-migrate/migrate
- `docker-compose.yml`の`services.api.environment`で、`POSTGRES_URL`を設定している。
- apiコンテナにアタッチして、下記を実行するとマイグレーションできる。

```shell
migrate -database ${POSTGRES_URL} -path ./repository/db/postgres/migration up
```

- ローカル環境で、`local/scripts/migrate`を実行すると、apiコンテナ内で上記マイグレーションを実行できる。

```shell
cd local/scripts
sh migrate up
sh migrate down
sh migrate force 202203281529 up
```

## Swagger
### swagger-uiのURL
- http://localhost:8001

### OpenAPIのBundle方法

- [swagger-cli](https://github.com/APIDevTools/swagger-cli)を使ってbundleする。

```shell
> npm install -g @apidevtools/swagger-cli
> swagger-cli bundle -o openapi.yaml -t yaml openapi/base.yaml
```

#### bundle script

- 上記のswagger-cliコマンドを書いた。下記の`openapi/bundle`を実行したら、openapi.yamlが最新状態になる。

```shell
> sh openapi/bundle
```

## Test

```shell
go test ./...
```