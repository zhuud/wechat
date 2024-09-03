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
https://go-zero.dev/docs/tutorials/cli/api  
https://go-zero.dev/docs/tutorials/cli/model  
> goctl api go -api wechat.api -dir . --style=go_zero
> goctl model mysql ddl --cache=true --style=go_zero --src=sql/external_user.sql
> goctl rpc protoc wechat.proto  --go_out=. --go-grpc_out=. --zrpc_out=. -m --style go_zero

> rpc client 全部目录都可以删除重新生成   
> rpc server 目录register.go 要保留，其他都可以删除重新生成

### wechat sdk
https://powerwechat.artisan-cloud.com/zh/wecom/contacts.html

### 常用代码
```
# 协程组管理
group := threading.NewRoutineGroup()
...
group.RunSafe(listener)
...
group.Wait()
```
```
# 程序关闭注册
proc.AddShutdownListener(func() {
    client.Close()
})
```

```
# 批处理协程处理 定时定量执行
executors.NewBulkExecutor(func(tasks []any) {
    for _, task := range tasks {
        ...
    }
}, bulkOpts...)
```
```
# 熔断器
brk := breaker.NewBreaker()
brk.Do(req func() error) error
```
```
# 日志耗时
tn := time.Now()
logx.WithDuration(time.Since(tn)).Info("hhh")
```
