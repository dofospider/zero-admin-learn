package role

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-admin-learn/api/internal/logic/sys/role"
	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"
)

func RoleMenuIdsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RoleMenuIdsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := role.NewRoleMenuIdsLogic(r.Context(), svcCtx)
		resp, err := l.RoleMenuIds(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
