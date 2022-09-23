package log

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-admin-learn/api/internal/logic/sys/log"
	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"
)

func LoginLogDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteLoginLogReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := log.NewLoginLogDeleteLogic(r.Context(), svcCtx)
		resp, err := l.LoginLogDelete(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
