package tag

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-admin-learn/api/internal/logic/member/tag"
	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"
)

func MemberTagDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteMemberTagReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := tag.NewMemberTagDeleteLogic(r.Context(), svcCtx)
		resp, err := l.MemberTagDelete(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
