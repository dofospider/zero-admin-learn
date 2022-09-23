package memberprice

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-admin-learn/api/internal/logic/product/memberprice"
	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"
)

func MemberPriceListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListMemberPriceReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := memberprice.NewMemberPriceListLogic(r.Context(), svcCtx)
		resp, err := l.MemberPriceList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
