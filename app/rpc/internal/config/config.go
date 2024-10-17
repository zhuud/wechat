package config

import (
	"time"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zhuud/go-library/svc/fasthttp"
)

const (
	LocalCacheExpire = time.Second * 60
	CropYx           = "yx"
)

// Config 总配置
type Config struct {
	zrpc.RpcServerConf
	Kafka      kq.KqConf           `json:",optional"`
	WechatDb   MysqlConf           `json:",optional"`
	CacheRedis redis.RedisConf     `json:",optional"`
	ModelRedis cache.CacheConf     `json:",optional"`
	FastHttp   fasthttp.ClientConf `json:",optional"`
	WeCom      []WeComConf         `json:",optional"`
}

// Config 子配置
type (
	MysqlConf struct {
		WechatDataSource  string `json:",optional"`
		BizUserDataSource string `json:",optional"`
		Timeout           string `json:",default=1s"`
		ReadTimeout       string `json:",default=5s"`
		WriteTimeout      string `json:",default=5s"`
		MaxIdleConn       int    `json:",optional"`
		MaxOpenConn       int    `json:",optional"`
		ConnMaxLifeTime   int    `json:",optional"`
		ConnMaxIdleTime   int    `json:",optional"`
	}

	WeComConf struct {
		CorpKey               string
		CorpName              string
		CorpId                string
		ContactSecret         string
		ExternalContactSecret string
		ChatSecret            string
		AgentId               string
		AgentSecret           string
		HttpDebug             bool `json:",default=false"`
	}
)

// 配置中心配置格式
type (
	dbConf struct {
		Driver   string `json:"driver"`
		Host     string `json:"host"`
		Dbname   string `json:"dbname"`
		Port     string `json:"port"`
		Username string `json:"username"`
		Password string `json:"password"`
		Charset  string `json:"charset"`
	}

	redisConf struct {
		Host string `json:"host"`
		Port string `json:"port"`
		Auth string `json:"auth"`
	}
)
