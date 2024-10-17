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
	Alarm      *Alarm
	FastHttp   *fasthttp.Client  // fast http
	Redis      *redis.Redis      // redis
	LocalCache *collection.Cache // local cache

	WeCom       *WeCom            // 企微api
	WechatLimit *WechatTokenLimit // 企微限流

	// model
	ModelBizUser                     model.TbUserOpenModel
	ModelExternalUser                model.TbExternalUserModel
	ModelExternalUserFollow          model.TbExternalUserFollowModel
	ModelExternalUserFollowAttribute model.TbExternalUserFollowAttributeModel
	ModelUserServiceQrcodeModel      model.TbUserServiceQrcodeModel
	ModelUserServiceQrcodeConclusion model.TbUserServiceQrcodeConclusionsModel
	ModelExternalUserTag             model.TbExternalUserTagModel
	ModelExternalUserAttribute       model.TbExternalUserAttributeModel
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
		wechatDbConn := sqlx.NewMysql(c.WechatDb.WechatDataSource)
		bizUserDbConn := sqlx.NewMysql(c.WechatDb.BizUserDataSource)
		// redis
		redisConn := redis.MustNewRedis(c.CacheRedis)

		//
		ctx = &ServiceContext{
			Config: c,

			// infra
			Alarm:      &Alarm{},
			FastHttp:   fasthttp.NewFastHttp(c.FastHttp),
			Redis:      redisConn,
			LocalCache: localCache,

			// wecom
			WeCom: NewWeCom(c.WeCom, c.CacheRedis),
			// limit
			WechatLimit: NewTokenWechatLimit(redisConn),

			// model
			ModelBizUser:                     model.NewTbUserOpenModel(bizUserDbConn),
			ModelExternalUser:                model.NewTbExternalUserModel(wechatDbConn),
			ModelExternalUserAttribute:       model.NewTbExternalUserAttributeModel(wechatDbConn),
			ModelExternalUserFollow:          model.NewTbExternalUserFollowModel(wechatDbConn),
			ModelExternalUserFollowAttribute: model.NewTbExternalUserFollowAttributeModel(wechatDbConn),
			ModelUserServiceQrcodeModel:      model.NewTbUserServiceQrcodeModel(wechatDbConn),
			ModelUserServiceQrcodeConclusion: model.NewTbUserServiceQrcodeConclusionsModel(wechatDbConn),
		}
	})

	return ctx
}

func GetSvcCtx() *ServiceContext {
	if ctx == nil {
		log.Fatalf("svc.GetSvcCtx is nil, use NewServiceContext before")
	}
	return ctx
}
