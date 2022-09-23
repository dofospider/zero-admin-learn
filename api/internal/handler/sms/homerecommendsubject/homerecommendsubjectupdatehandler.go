package homerecommendsubject

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-admin-learn/api/internal/logic/sms/homerecommendsubject"
	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"
)

func HomeRecommendSubjectUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateHomeRecommendSubjectReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := homerecommendsubject.NewHomeRecommendSubjectUpdateLogic(r.Context(), svcCtx)
		resp, err := l.HomeRecommendSubjectUpdate(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
