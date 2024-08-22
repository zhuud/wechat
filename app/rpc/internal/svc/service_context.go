package svc

import (
	"rpc/internal/config"
	"rpc/model"

	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhuud/go-library/svc/fasthttp"
)

type ServiceContext struct {
	Config     config.Config
	FastHttp   *fasthttp.Client  // fast http
	Redis      *redis.Redis      // redis
	LocalCache *collection.Cache // local cache

	WeCom *WeCom // 企微api

	// model
	ModelExternalUser                model.TbExternalUserModel
	ModelExternalUserFollow          model.TbExternalUserFollowModel
	ModelExternalUserFollowAttribute model.TbExternalUserFollowAttributeModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	// local cache
	localCache, err := collection.NewCache(config.LocalCacheExpire)
	logx.Must(err)
	// mysql
	conn := sqlx.NewMysql(c.WechatDb.DataSource)

	//
	return &ServiceContext{
		Config: c,

		// infra
		FastHttp:   fasthttp.NewFastHttp(c.FastHttp),
		Redis:      redis.MustNewRedis(c.CacheRedis),
		LocalCache: localCache,

		// wecom
		WeCom: NewWeCom(c.WeCom, c.CacheRedis),

		// model
		ModelExternalUser:                model.NewTbExternalUserModel(conn, c.ModelRedis),
		ModelExternalUserFollow:          model.NewTbExternalUserFollowModel(conn, c.ModelRedis),
		ModelExternalUserFollowAttribute: model.NewTbExternalUserFollowAttributeModel(conn, c.ModelRedis),
	}
}
