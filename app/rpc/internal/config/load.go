package config

import (
	"fmt"
	"log"
	"time"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zhuud/go-library/svc/conf"
	"github.com/zhuud/go-library/svc/conf/confx"
	"github.com/zhuud/go-library/svc/zookeeper"
)

func setup() {
	path := conf.AppConfPath
	if CmdConfigFile != nil {
		path = *CmdConfigFile
	}
	conf.MustSetUp(conf.Conf{FilePath: path})
	conf.AppendReader(confx.NewZookeeperReader(zookeeper.MustNewZookeeperClient()))
}

func MustLoad() Config {
	setup()

	c := Config{}
	err := conf.GetUnmarshal("", &c)
	if err != nil {
		log.Fatalf("config.MustLoad error: %s", err.Error())
	}

	// 密钥
	c.Secret = mustLoadSecret()

	// 数据库
	if len(c.WechatDb.WechatDataSource) == 0 {
		c.WechatDb.WechatDataSource = mustLoadMysql(c, "db_wechat")
	}
	if len(c.WechatDb.BizUserDataSource) == 0 {
		c.WechatDb.BizUserDataSource = mustLoadMysql(c, "db_user_rw")
	}
	if len(c.WechatDb.BizUserDataSource) == 0 {
		c.WechatDb.BizUserDataSource = mustLoadMysql(c, "db_event")
	}

	// 缓存
	if len(c.CacheRedis.Host) == 0 {
		c.CacheRedis = mustLoadRedis("redis")
	}
	if len(c.ModelRedis) == 0 {
		c.ModelRedis = cache.CacheConf{
			{RedisConf: c.CacheRedis, Weight: 100},
		}
	}

	c.WeCom = mustLoadWeCom()

	return c
}

func mustLoadSecret() Secret {
	sc := Secret{}
	err := conf.GetUnmarshal(fmt.Sprintf("/qconf/web-config/%s", "privatedomain_wechat_secret"), &sc)
	if err != nil {
		log.Fatalf("config.mustLoadSecret error: %s", err.Error())
	}
	if len(sc.Cryptographic) == 0 {
		log.Fatalf("config.mustLoadSecret incomplete secret config, config: %v", sc)
	}

	return sc
}

func mustLoadMysql(c Config, dbname string) string {
	dc := dbConf{}

	err := conf.GetUnmarshal(fmt.Sprintf("/qconf/web-config/%s", dbname), &dc)
	if err != nil {
		log.Fatalf("config.mustLoadMysql error: %s", err.Error())
	}
	if len(dc.Driver) == 0 || len(dc.Host) == 0 || len(dc.Port) == 0 || len(dc.Dbname) == 0 || len(dc.Username) == 0 || len(dc.Password) == 0 || len(dc.Charset) == 0 {
		log.Fatalf("config.mustLoadMysql incomplete mysl config, config: %v", dc)
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset%s&parseTime=True&loc=Local&timeout=%s&readTimeout=%s&writeTimeout=%s", dc.Username, dc.Password, dc.Host, dc.Port, dc.Dbname, dc.Charset, c.WechatDb.Timeout, c.WechatDb.ReadTimeout, c.WechatDb.WriteTimeout)
}

func mustLoadRedis(dbname string) redis.RedisConf {
	rc := redisConf{}

	err := conf.GetUnmarshal(fmt.Sprintf("/qconf/web-config/%s", dbname), &rc)
	if err != nil {
		log.Fatalf("config.mustLoadRedis error: %s", err.Error())
	}
	if len(rc.Host) == 0 || len(rc.Port) == 0 || len(rc.Auth) == 0 {
		log.Fatalf("config.mustLoadRedis incomplete redis config, config: %v", rc)
	}

	return redis.RedisConf{
		Host:        fmt.Sprintf("%s:%s", rc.Host, rc.Port),
		Pass:        rc.Auth,
		Type:        "node",
		PingTimeout: time.Second * 5,
	}
}

func mustLoadWeCom() []WeComConf {

	wcList := make([]WeComConf, 0)

	err := conf.GetUnmarshal(fmt.Sprintf("/application-secret-key/%s", "wecom_corp"), &wcList)
	if err != nil {
		log.Fatalf("config.mustLoadWeCom error: %s", err.Error())
	}
	if len(wcList) == 0 {
		log.Fatalf("config.mustLoadWeCom no config")
	}
	for _, wc := range wcList {
		if len(wc.CorpId) == 0 || len(wc.CorpKey) == 0 || len(wc.ContactSecret) == 0 || len(wc.ExternalContactSecret) == 0 || len(wc.ChatSecret) == 0 {
			log.Fatalf("config.mustLoadWeCom incomplete wecom config, config: %v", wc)
		}
	}

	return wcList
}
