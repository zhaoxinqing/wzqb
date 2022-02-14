package logic

import (
	"context"

	"go-template/service/user/api/internal/svc"
	"go-template/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteReportLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteReportLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteReportLogic {
	return DeleteReportLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteReportLogic) DeleteReport(req types.DeleteReportParam) error {
	// todo: add your logic here and delete this line

	return nil
}
