package user

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
	"github.com/zhuud/go-library/svc/conf"
	"github.com/zhuud/go-library/utils"
	"rpc/internal/svc"
	"rpc/internal/types"
	"rpc/model"
	"rpc/wechat"
)

// SetCacheYn 是否设置缓存
var SetCacheYn = false

func init() {
	// 设置环境、测试环境无cache、生产环境启动cache
	if conf.IsProd() {
		SetCacheYn = true
	}
}

type GetExternalUserCacheLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

// post 基础缓存 临时查询存储
type userUnit struct {
	Base       map[string]*wechat.ExternalUser
	Attribute  map[string][]*model.TbExternalUserFollowAttribute
	UserFollow map[string][]*wechat.ExternalUserFollowUser
}

func NewGetExternalUserCacheLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExternalUserCacheLogic {
	return &GetExternalUserCacheLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (t *GetExternalUserCacheLogic) GetUserCache(req *wechat.ExternalUserInfoReq) (externalUser map[string]*types.ExternalUser, err error) {
	externalUser = make(map[string]*types.ExternalUser)
	externalUserIdList := req.ExternalUseridList

	externalUserIdList = utils.ArrayUnique(externalUserIdList)
	if len(externalUserIdList) == 0 {
		return externalUser, errors.New(`缺少external_user id`)
	}

	// 缓存 todo
	userToCacheUserList, hasCacheUserId, err := t.getUserListCache(t.ctx, externalUserIdList)

	//diff 有缓存和没有缓存的数据
	noCacheUserIdList := utils.ArrayDiff(req.ExternalUseridList, hasCacheUserId)

	if len(noCacheUserIdList) == 0 {
		return userToCacheUserList, err
	}

	// 获取数据库信息
	userToDBUserList, err := t.getUserListByDB(t.ctx, req)
	if err != nil {
		t.Error(`getUserListByDB`, err)
		return userToCacheUserList, err
	}

	//合并缓存和数据库的数据
	userToUserList := t.mergePostListCacheDBMap(userToCacheUserList, userToDBUserList)

	return userToUserList, err
}

// mergePostListCacheDBMap map 合并
func (t *GetExternalUserCacheLogic) mergePostListCacheDBMap(sm ...map[string]*types.ExternalUser) map[string]*types.ExternalUser {
	r := map[string]*types.ExternalUser{}
	for _, m := range sm {
		for k, v := range m {
			r[k] = v
		}
	}
	return r
}

// todo 缓存
func (t *GetExternalUserCacheLogic) getUserListCache(ctx context.Context, externalUserIdList []string) (user map[string]*types.ExternalUser, hasCacheUserId []string, err error) {
	if !SetCacheYn {
		return
	}

	return
}

func (t *GetExternalUserCacheLogic) getUserListByDB(ctx context.Context, req *wechat.ExternalUserInfoReq) (user map[string]*types.ExternalUser, err error) {
	if len(req.ExternalUseridList) == 0 {
		return
	}

	externalUserIdList := req.ExternalUseridList
	uf := t.GetBaseField(req)

	group := threading.NewRoutineGroup()

	ub := types.ExternalUserUnit{}

	for _, unitField := range uf {
		group.RunSafe(func() {
			t.getUnit(ctx, externalUserIdList, unitField, &ub)
		})
	}

	group.Wait()

	for _, externalUserId := range externalUserIdList {
		if externalUser, ok := ub.ExternalUser[externalUserId]; ok {
			user[externalUserId].ExternalUser = externalUser
		}

		if externalUserFollow, ok := ub.ExternalUserFollow[externalUserId]; ok {
			user[externalUserId].ExternalUserFollow = externalUserFollow
		}

		if externalUserFollowAttrList, ok := ub.ExternalUserFollowAttribute[externalUserId]; ok {
			user[externalUserId].ExternalUserFollowAttributeDB = externalUserFollowAttrList
			user[externalUserId].ExternalUserFollowAttribute = NewGetExternalUserAttributeLogic(t.ctx, t.svcCtx).HandleAttributeFormat(externalUserFollowAttrList)
		}

	}

	return
}

func (t *GetExternalUserCacheLogic) GetBaseField(req *wechat.ExternalUserInfoReq) []string {
	uf := []string{`user`}
	if req.Opt == nil {
		return uf
	}

	if req.Opt.NeedFollow {
		uf = append(uf, `follow`)
	}

	if req.Opt.NeedAttribute {
		uf = append(uf, `follow_attribute`)
	}

	return uf
}

// getUnit 获取用户得某个对象属性，返回的信息为  external_user_id => []object
func (t *GetExternalUserCacheLogic) getUnit(ctx context.Context, externalUserIdList []string, kItem string, ub *types.ExternalUserUnit) {

	switch kItem {
	case `user`:
		// 查询用户-基础数据
		userIdToBaseInfo, err := NewGetExternalUserBaseLogic(t.ctx, t.svcCtx).GetUserListByExternalUserIdList(externalUserIdList)
		if err != nil {
			logc.Error(ctx, `cache-获取用户属性信息失败, GetUserListByExternalUserIdList_err`, err, externalUserIdList)
			return
		}
		ub.ExternalUser = userIdToBaseInfo

	case `follow`:
		// 查询用户-外部联系人属性表
		userIdToFollow, err := NewGetExternalUserFollowLogic(t.ctx, t.svcCtx).GetUserFollowListByExternalUserIdList(externalUserIdList)
		if err != nil {
			logc.Error(ctx, `cache-获取用户属性信息失败, GetUserFollowListByExternalUserIdList_err`, err, externalUserIdList)
			return
		}
		ub.ExternalUserFollow = userIdToFollow

	case `follow_attribute`:
		// 查询用户-外部联系人添加员工信息属性表
		userIdToAttributeList, err := NewGetExternalUserAttributeLogic(t.ctx, t.svcCtx).GetUserFollowAttributeByExternalUserIdList(externalUserIdList)
		if err != nil {
			logc.Error(ctx, `cache-获取用户属性信息失败, GetUserFollowAttributeByExternalUserIdList_err`, err, externalUserIdList)
			return
		}
		ub.ExternalUserFollowAttribute = userIdToAttributeList

	default:
		return
	}
	return
}
