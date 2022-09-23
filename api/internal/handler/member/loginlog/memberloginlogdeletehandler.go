package loginlog

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-admin-learn/api/internal/logic/member/loginlog"
	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"
)

func MemberLoginLogDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteMemberLoginLogReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := loginlog.NewMemberLoginLogDeleteLogic(r.Context(), svcCtx)
		resp, err := l.MemberLoginLogDelete(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
