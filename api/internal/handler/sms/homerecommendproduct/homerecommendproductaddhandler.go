package homerecommendproduct

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-admin-learn/api/internal/logic/sms/homerecommendproduct"
	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"
)

func HomeRecommendProductAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddHomeRecommendProductReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := homerecommendproduct.NewHomeRecommendProductAddLogic(r.Context(), svcCtx)
		resp, err := l.HomeRecommendProductAdd(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
