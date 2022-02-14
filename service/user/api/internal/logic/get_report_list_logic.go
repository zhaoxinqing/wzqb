package logic

import (
	"context"

	"go-template/service/user/api/internal/svc"
	"go-template/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetReportListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetReportListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetReportListLogic {
	return GetReportListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetReportListLogic) GetReportList(req types.GetReportListParam) error {
	// todo: add your logic here and delete this line

	return nil
}
