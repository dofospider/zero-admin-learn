package flashpromotionsession

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-admin-learn/api/internal/logic/sms/flashpromotionsession"
	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"
)

func FlashPromotionSessionDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteFlashPromotionSessionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := flashpromotionsession.NewFlashPromotionSessionDeleteLogic(r.Context(), svcCtx)
		resp, err := l.FlashPromotionSessionDelete(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
