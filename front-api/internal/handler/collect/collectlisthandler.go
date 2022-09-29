package collect

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-admin-learn/front-api/internal/logic/collect"
	"zero-admin-learn/front-api/internal/svc"
	"zero-admin-learn/front-api/internal/types"
)

func CollectListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CollectListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := collect.NewCollectListLogic(r.Context(), svcCtx)
		resp, err := l.CollectList(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
