package flashpromotionsession

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-admin-learn/api/internal/logic/sms/flashpromotionsession"
	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"
)

func FlashPromotionSessionUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateFlashPromotionSessionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := flashpromotionsession.NewFlashPromotionSessionUpdateLogic(r.Context(), svcCtx)
		resp, err := l.FlashPromotionSessionUpdate(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
