package homenewproduct

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-admin-learn/api/internal/logic/sms/homenewproduct"
	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"
)

func HomeNewProductDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteHomeNewProductReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := homenewproduct.NewHomeNewProductDeleteLogic(r.Context(), svcCtx)
		resp, err := l.HomeNewProductDelete(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
