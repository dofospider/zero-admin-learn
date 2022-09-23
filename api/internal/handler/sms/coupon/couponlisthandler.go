package coupon

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-admin-learn/api/internal/logic/sms/coupon"
	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"
)

func CouponListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListCouponReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := coupon.NewCouponListLogic(r.Context(), svcCtx)
		resp, err := l.CouponList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
