package svc

import (
	"log"
	"sync"

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

	WeCom       *WeCom             // 企微api
	WechatLimit *WechatPeriodLimit // 企微限流

	// model
	ModelExternalUser                model.TbExternalUserModel
	ModelExternalUserFollow          model.TbExternalUserFollowModel
	ModelExternalUserFollowAttribute model.TbExternalUserFollowAttributeModel
	ModelUserServiceQrcodeModel      model.UserServiceQrcodeModel
}

var (
	ctx  *ServiceContext
	once sync.Once
)

func NewServiceContext(c config.Config) *ServiceContext {

	once.Do(func() {
		// local cache
		localCache, err := collection.NewCache(config.LocalCacheExpire)
		logx.Must(err)
		// mysql
		msyqlConn := sqlx.NewMysql(c.WechatDb.DataSource)
		// redis
		redisConn := redis.MustNewRedis(c.CacheRedis)

		//
		ctx = &ServiceContext{
			Config: c,

			// infra
			FastHttp:   fasthttp.NewFastHttp(c.FastHttp),
			Redis:      redisConn,
			LocalCache: localCache,

			// wecom
			WeCom: NewWeCom(c.WeCom, c.CacheRedis),
			// limit
			WechatLimit: NewWechatPeriodLimit(redisConn),

			// model
			ModelExternalUser:                model.NewTbExternalUserModel(msyqlConn, c.ModelRedis),
			ModelExternalUserFollow:          model.NewTbExternalUserFollowModel(msyqlConn, c.ModelRedis),
			ModelExternalUserFollowAttribute: model.NewTbExternalUserFollowAttributeModel(msyqlConn, c.ModelRedis),
			ModelUserServiceQrcodeModel:      model.NewUserServiceQrcodeModel(msyqlConn),
		}
	})

	return ctx
}

func GetSvcCtx() *ServiceContext {
	if ctx == nil {
		log.Fatalf("svc.GetSvcCtx is nil")
	}
	return ctx
}
