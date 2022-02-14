package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-template/service/user/api/internal/logic"
	"go-template/service/user/api/internal/svc"
	"go-template/service/user/api/internal/types"
)

func UserRegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterParam
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserRegisterLogic(r.Context(), svcCtx)
		err := l.UserRegister(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
