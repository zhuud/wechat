package externalwayqr

import (
	"net/http"

	"api/internal/logic/externalwayqr"
	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 企微联系人二维码列表
func ExternalWayQrListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ExternalWayQrListRequest

		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := externalwayqr.NewExternalWayQrListLogic(r.Context(), svcCtx)
		resp, err := l.ExternalWayQrList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
