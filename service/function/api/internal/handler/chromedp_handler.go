package handler

import (
	"net/http"

	"wzqb/common"
	"wzqb/service/function/api/internal/logic"
	"wzqb/service/function/api/internal/svc"
	"wzqb/service/function/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ChromedpHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChromedpParam
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewChromedpLogic(r.Context(), svcCtx)
		err := l.Chromedp(req)
		common.Response(w, nil, err)
	}
}
