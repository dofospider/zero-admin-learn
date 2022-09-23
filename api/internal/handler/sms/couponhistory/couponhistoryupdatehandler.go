package couponhistory

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-admin-learn/api/internal/logic/sms/couponhistory"
	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"
)

func CouponHistoryUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateCouponHistoryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := couponhistory.NewCouponHistoryUpdateLogic(r.Context(), svcCtx)
		resp, err := l.CouponHistoryUpdate(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
