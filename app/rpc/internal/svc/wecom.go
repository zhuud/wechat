package svc

import (
    "log"
    "sync"

    "rpc/internal/config"

    "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
    "github.com/ArtisanCloud/PowerWeChat/v3/src/work"
    "github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact"
    "github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/contactWay"
    "github.com/ArtisanCloud/PowerWeChat/v3/src/work/user"
    "github.com/zeromicro/go-zero/core/logx"
    "github.com/zeromicro/go-zero/core/stores/redis"
)

type (
    WeCom struct {
        api map[string]*WeComApi
    }

    WeComApi struct {
        User         *user.Client
        ExternalUser *externalContact.Client
        ContactWay   *contactWay.Client
    }
)

var (
    wc     *WeCom
    wcOnce sync.Once
)

func NewWeCom(confList []config.WeComConf, redisConf redis.RedisConf) *WeCom {

    wcOnce.Do(func() {
        for _, conf := range confList {
            wc = &WeCom{
                api: make(map[string]*WeComApi),
            }
            wcApi := &WeComApi{}

            cache := kernel.NewRedisClient(&kernel.UniversalOptions{
                Addrs:    []string{redisConf.Host},
                Password: redisConf.Pass,
            })
            initContact(wcApi, conf, cache)
            initExternalContact(wcApi, conf, cache)

            wc.api[conf.CorpKey] = wcApi
        }
    })

    return wc
}

func initContact(wc *WeComApi, corpConf config.WeComConf, cache kernel.CacheInterface) {
    wcc, err := work.NewWork(&work.UserConfig{
        CorpID:    corpConf.CorpId,
        Secret:    corpConf.ContactSecret,
        HttpDebug: corpConf.HttpDebug,
        Cache:     cache,
        OAuth:     work.OAuth{Callback: "https://xx.com"},
    })
    if err != nil {
        log.Fatalf("svc.NewWeCom - wcc.NewWork err: %v", err)
    }
    wc.User = wcc.User
}

func initExternalContact(wc *WeComApi, corpConf config.WeComConf, cache kernel.CacheInterface) {
    wce, err := work.NewWork(&work.UserConfig{
        CorpID:    corpConf.CorpId,
        Secret:    corpConf.ExternalContactSecret,
        HttpDebug: corpConf.HttpDebug,
        Cache:     cache,
        OAuth:     work.OAuth{Callback: "https://xx.com"},
    })
    if err != nil {
        log.Fatalf("svc.NewWeCom - wce.NewWork err: %v", err)
    }
    wc.ExternalUser = wce.ExternalContact
    wc.ContactWay = wce.ExternalContactContactWay
}

func (w *WeCom) WithCorp(corpKey string) *WeComApi {
    if api, ok := w.api[corpKey]; ok {
        return api
    }
    logx.Errorf("svc.WeCom.WithCorp - not exists crop,  api: %v, crop: %s", w.api, corpKey)
    return nil
}
