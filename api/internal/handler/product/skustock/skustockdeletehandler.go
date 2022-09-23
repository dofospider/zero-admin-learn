package skustock

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-admin-learn/api/internal/logic/product/skustock"
	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"
)

func SkuStockDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteSkuStockReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := skustock.NewSkuStockDeleteLogic(r.Context(), svcCtx)
		resp, err := l.SkuStockDelete(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
