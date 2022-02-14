package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-template/service/user/api/internal/logic"
	"go-template/service/user/api/internal/svc"
	"go-template/service/user/api/internal/types"
)

func EditReportHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EditReportParam
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewEditReportLogic(r.Context(), svcCtx)
		err := l.EditReport(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
