package logic

import (
	"context"

	"go-template/service/user/api/internal/svc"
	"go-template/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddReportLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddReportLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddReportLogic {
	return AddReportLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddReportLogic) AddReport(req types.AddReportParam) error {
	// todo: add your logic here and delete this line

	return nil
}
