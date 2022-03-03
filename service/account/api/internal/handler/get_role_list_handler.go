package handler

import (
	"net/http"

	"go-template/service/account/api/internal/logic"
	"go-template/service/account/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// GetRoleListHandler 获取角色列表
func GetRoleListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetRoleListLogic(r.Context(), svcCtx)
		err := l.GetRoleList()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
