package cart

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-admin-learn/front-api/internal/logic/cart"
	"zero-admin-learn/front-api/internal/svc"
	"zero-admin-learn/front-api/internal/types"
)

func CartFastAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CartFastAddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := cart.NewCartFastAddLogic(r.Context(), svcCtx)
		resp, err := l.CartFastAdd(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
