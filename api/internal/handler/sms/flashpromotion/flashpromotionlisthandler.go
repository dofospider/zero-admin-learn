package flashpromotion

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-admin-learn/api/internal/logic/sms/flashpromotion"
	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"
)

func FlashPromotionListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListFlashPromotionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := flashpromotion.NewFlashPromotionListLogic(r.Context(), svcCtx)
		resp, err := l.FlashPromotionList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
