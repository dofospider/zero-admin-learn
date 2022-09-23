package cart

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-admin-learn/api/internal/logic/order/cart"
	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"
)

func CartItemDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteCartItemReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := cart.NewCartItemDeleteLogic(r.Context(), svcCtx)
		resp, err := l.CartItemDelete(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
