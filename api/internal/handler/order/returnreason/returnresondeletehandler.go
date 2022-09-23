package returnreason

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-admin-learn/api/internal/logic/order/returnreason"
	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"
)

func ReturnResonDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteReturnResonReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := returnreason.NewReturnResonDeleteLogic(r.Context(), svcCtx)
		resp, err := l.ReturnResonDelete(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
