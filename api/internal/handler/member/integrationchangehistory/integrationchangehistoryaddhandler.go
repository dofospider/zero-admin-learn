package integrationchangehistory

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-admin-learn/api/internal/logic/member/integrationchangehistory"
	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"
)

func IntegrationChangeHistoryAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddIntegrationChangeHistoryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := integrationchangehistory.NewIntegrationChangeHistoryAddLogic(r.Context(), svcCtx)
		resp, err := l.IntegrationChangeHistoryAdd(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
