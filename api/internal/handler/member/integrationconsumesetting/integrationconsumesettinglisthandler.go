package integrationconsumesetting

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-admin-learn/api/internal/logic/member/integrationconsumesetting"
	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"
)

func IntegrationConsumeSettingListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListIntegrationConsumeSettingReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := integrationconsumesetting.NewIntegrationConsumeSettingListLogic(r.Context(), svcCtx)
		resp, err := l.IntegrationConsumeSettingList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
