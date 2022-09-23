package operatehistory

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-admin-learn/api/internal/logic/order/operatehistory"
	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"
)

func OperateHistoryAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddOperateHistoryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := operatehistory.NewOperateHistoryAddLogic(r.Context(), svcCtx)
		resp, err := l.OperateHistoryAdd(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
