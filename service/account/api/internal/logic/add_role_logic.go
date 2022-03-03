package logic

import (
	"context"

	"go-template/service/account/api/internal/svc"
	"go-template/service/account/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddRoleLogic {
	return AddRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddRole 添加角色
func (l *AddRoleLogic) AddRole(req types.AddRoleParam) error {
	// todo: add your logic here and delete this line

	return nil
}
