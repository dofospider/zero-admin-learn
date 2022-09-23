package category

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-admin-learn/api/internal/logic/product/category"
	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"
)

func ProductCategoryAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddProductCategoryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := category.NewProductCategoryAddLogic(r.Context(), svcCtx)
		resp, err := l.ProductCategoryAdd(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
