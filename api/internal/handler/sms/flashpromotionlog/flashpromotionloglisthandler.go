package flashpromotionlog

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-admin-learn/api/internal/logic/sms/flashpromotionlog"
	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"
)

func FlashPromotionLogListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListFlashPromotionLogReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := flashpromotionlog.NewFlashPromotionLogListLogic(r.Context(), svcCtx)
		resp, err := l.FlashPromotionLogList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
