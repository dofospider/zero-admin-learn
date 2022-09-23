package brand

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-admin-learn/api/internal/logic/product/brand"
	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"
)

func ProductBrandDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteProductBrandReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := brand.NewProductBrandDeleteLogic(r.Context(), svcCtx)
		resp, err := l.ProductBrandDelete(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
