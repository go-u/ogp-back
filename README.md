![](https://github.com/go-u/ogp-back/workflows/Test/badge.svg)
![](https://github.com/go-u/ogp-back/workflows/Staging/badge.svg)
![](https://github.com/go-u/ogp-back/workflows/Production/badge.svg) 

<p align="center"><img src="https://github.com/go-u/ogp-index/blob/master/docs/systems.jpg" alt="Systems"></p>


## ハイライト
### :star2: DDD/クリーンアーキテクチャに準じた実装  

### :unlock: JWT(Json Web Token)を利用したモダンな認証方式で実装  

### :heart: APIにアクセス制限を導入しセキュリティを向上    

## システム構成について
Google系のクラウドサービスを利用しています(GCP / GAE / GCS / CloudSql / Firebase等)
- Api(Golang)はApp Engineスダンダードで運用
- データベースはCloud SQL
- ユーザ登録やJWT認証にFirebase

## セキュリティについて
- APIへのダイレクトアクセスはファイアーウォールで制限しています(DDos等の対策)  
- いたずら防止のためTorネットワークからのアクセス制限もしています

## CI/CDについて
Github ActionsでCI/CDを行っています。
- [プッシュ時に自動テスト](https://github.com/go-u/ogp-back/blob/master/.github/workflows/test.yml)
- [プルリク時に検証環境に自動デプロイ](https://github.com/go-u/ogp-back/blob/master/.github/workflows/deploy_staging.yml)
- [マージ時にプロダクション環境に自動デプロイ](https://github.com/go-u/ogp-back/blob/master/.github/workflows/deploy_production.yml)  
