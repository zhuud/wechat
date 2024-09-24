package externalwayqr

import (
	"net/http"

	"api/internal/logic/externalwayqr"
	"api/internal/svc"
	"api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 企微联系人二维码删除
func ExternalWayQrDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ExternalContactWayDelRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := externalwayqr.NewExternalWayQrDeleteLogic(r.Context(), svcCtx)
		resp, err := l.ExternalWayQrDelete(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
