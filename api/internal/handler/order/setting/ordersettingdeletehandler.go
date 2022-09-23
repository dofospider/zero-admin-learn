package setting

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-admin-learn/api/internal/logic/order/setting"
	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"
)

func OrderSettingDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteOrderSettingReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := setting.NewOrderSettingDeleteLogic(r.Context(), svcCtx)
		resp, err := l.OrderSettingDelete(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
