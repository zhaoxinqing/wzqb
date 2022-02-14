package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-template/service/user/api/internal/logic"
	"go-template/service/user/api/internal/svc"
)

func UserListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserListLogic(r.Context(), svcCtx)
		err := l.UserList()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
