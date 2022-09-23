package homeadvertise

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-admin-learn/api/internal/logic/sms/homeadvertise"
	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"
)

func HomeAdvertiseUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateHomeAdvertiseReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := homeadvertise.NewHomeAdvertiseUpdateLogic(r.Context(), svcCtx)
		resp, err := l.HomeAdvertiseUpdate(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
