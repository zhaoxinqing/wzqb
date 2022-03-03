package handler

import (
	"net/http"

	"go-template/service/account/api/internal/logic"
	"go-template/service/account/api/internal/svc"
	"go-template/service/account/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// AddRoleHandler 添加角色
func AddRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddRoleParam
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAddRoleLogic(r.Context(), svcCtx)
		err := l.AddRole(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
