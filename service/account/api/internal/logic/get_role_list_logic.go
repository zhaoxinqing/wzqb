package logic

import (
	"context"

	"go-template/service/account/api/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetRoleListLogic {
	return GetRoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetRoleList 获取角色列表
func (l *GetRoleListLogic) GetRoleList() error {
	// todo: add your logic here and delete this line

	return nil
}
