package logic

import (
	"context"
	"errors"
	"log"

	"go-template/common"
	"go-template/service/account/api/internal/svc"
	"go-template/service/account/api/internal/types"
	"go-template/service/account/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterLogic {
	return RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Register 用户注册
func (l *RegisterLogic) Register(req types.RegisterParam) (interface{}, error) {
	var (
		newAccount = model.AccountTable{}
		db         = l.svcCtx.DbEngin
		err        error
	)
	// 用户名检查
	err = db.Where("name = ?", req.Name).Find(&newAccount).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// 创建新用户
	if newAccount.ID == common.EmptyID {
		newAccount = model.AccountTable{
			Name:     req.Name,
			PassWord: req.Password,
		}
		err = l.svcCtx.DbEngin.Create(&newAccount).Error
		if err != nil {
			log.Println(err)
			return nil, err
		}
	} else {
		err = errors.New("用户名已存在,请重试")
		return nil, err
	}
	//
	return newAccount, nil
}
