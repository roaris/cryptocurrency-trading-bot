# cryptcurrency-trading-bot

## デプロイ(ECR + ECS(Fargate))
困ったら[#8](https://github.com/roaris/cryptocurrency-trading-bot/issues/8)参照

### Dockerfileを更新した場合
Dockerfileを更新した場合は、ECRにpushし直す

`aws configure`で、AdministratorAccessがアタッチされたIAMユーザのアクセスキーとシークレットアクセスキーをセットしておく

[Amazon Elastic Container Registry](https://ap-northeast-1.console.aws.amazon.com/ecr/repositories?region=ap-northeast-1) のプッシュコマンドの表示で出てくる指示に従う

### Fargateへのデプロイ
docker-compose.ymlを以下のように修正する

- `volumes`の部分を削除する
- `x-aws-vpc`で、VPC IDを指定する
- `app`の環境変数`DB_HOST`をELBのDNSにする

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
