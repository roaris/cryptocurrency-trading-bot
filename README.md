# cryptcurrency-trading-bot

bitFlyerのAPIを通じて、仮想通貨取引を行うプログラム

[wiki](https://github.com/roaris/cryptocurrency-trading-bot/wiki/cryptocurrency-trading-bot)

## 起動
### .envの用意
bitFlyer Lightningで取得したAPI KeyとAPI Secretを.envに書く

```
API_KEY = ...
API_SECRET = ...
```

### Dockerの起動
```
$ docker-compose build
$ docker-compose up
```

`docker-compose build`は初回とDockerfileを更新した時のみでOK

終了時は`docker-compose down`

## デプロイ(ECR + ECS(Fargate))
困ったら[#8](https://github.com/roaris/cryptocurrency-trading-bot/issues/8)参照

### ECRにDockerイメージをpushする
`aws configure`で、AdministratorAccessがアタッチされたIAMユーザのアクセスキーとシークレットアクセスキーをセットしておく

[Amazon Elastic Container Registry](https://ap-northeast-1.console.aws.amazon.com/ecr/repositories?region=ap-northeast-1) のプッシュコマンドの表示で出てくる指示に従う

### Fargateへのデプロイ
docker-compose.ymlを以下のように修正する

- `volumes`の部分を削除する
- `x-aws-vpc`で、VPC IDを指定する
- `app`の`build`を削除し、`image`でECRのレポジトリのURIを指定する
- `app`の環境変数`DB_HOST`をELBのDNS名にする

以下のコマンドでデプロイする

```
$ docker context use myecscontext
$ docker compose up
```

デプロイができたら、コンテキストをdefaultに戻す

```
$ docker context use default
```

### 動かなかったら
[CloudWatch](https://ap-northeast-1.console.aws.amazon.com/cloudwatch/home?region=ap-northeast-1#home:)にコンテナのログが出ているので、見る
