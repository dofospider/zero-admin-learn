package growthchangehistory

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-admin-learn/api/internal/logic/member/growthchangehistory"
	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"
)

func GrowthChangeHistoryDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteGrowthChangeHistoryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := growthchangehistory.NewGrowthChangeHistoryDeleteLogic(r.Context(), svcCtx)
		resp, err := l.GrowthChangeHistoryDelete(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
