package logic

import (
	"context"

	"go-template/service/user/api/internal/svc"
	"go-template/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditReportLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditReportLogic(ctx context.Context, svcCtx *svc.ServiceContext) EditReportLogic {
	return EditReportLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditReportLogic) EditReport(req types.EditReportParam) error {
	// todo: add your logic here and delete this line

	return nil
}
