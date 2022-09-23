package compayaddress

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-admin-learn/api/internal/logic/order/compayaddress"
	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"
)

func CompayAddressAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddCompayAddressReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := compayaddress.NewCompayAddressAddLogic(r.Context(), svcCtx)
		resp, err := l.CompayAddressAdd(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
