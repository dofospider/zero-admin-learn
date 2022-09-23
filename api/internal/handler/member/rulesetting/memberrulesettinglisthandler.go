package rulesetting

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-admin-learn/api/internal/logic/member/rulesetting"
	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"
)

func MemberRuleSettingListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListMemberRuleSettingReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := rulesetting.NewMemberRuleSettingListLogic(r.Context(), svcCtx)
		resp, err := l.MemberRuleSettingList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
