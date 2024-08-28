# wechat
Wechat Private Deployment 

## 启动
### rpc
> cd app/rpc  
> go run main.go -f etc/config.local.yaml

### api
> cd ../api
> go run main.go -f etc/config.local.yaml

### curl
```
curl --location --request POST 'http://127.0.0.1:8080/externaluser/list' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
--header 'Content-Type: application/json' \
--header 'Accept: */*' \
--header 'Host: 127.0.0.1:8080' \
--header 'Connection: keep-alive' \
--data-raw '{"external_userid_list":["wmYYltDAAAlg093GN65jtwLAn1VqOi5g"]}'
```

### cmd
> go run main.go -f etc/config.local.yaml CmdSyncExternalUser

### mq
> go run main.go -f etc/config.local.yaml ConsumerSyncExternalUser

### 代码生成
https://go-zero.dev/docs/tutorials/cli/rpc
> rpc client 全部目录都可以删除重新生成
> rpc server 目录register.go 要保留，其他都可以删除重新生成
https://go-zero.dev/docs/tutorials/cli/api
https://go-zero.dev/docs/tutorials/cli/model

### wechat sdk
https://powerwechat.artisan-cloud.com/zh/wecom/contacts.html
