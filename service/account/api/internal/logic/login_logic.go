package logic

import (
	"context"
	"errors"
	"log"

	"go-template/common"
	"go-template/middleware"
	"go-template/service/account/api/internal/svc"
	"go-template/service/account/api/internal/types"
	"go-template/service/account/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// LoginResp 用户登陆响应信息
type LoginResp struct {
	UserInfo model.AccountTable
	Token    string
}

// Login 登陆
func (l *LoginLogic) Login(req types.RegisterParam) (interface{}, error) {
	var (
		userInfo = model.AccountTable{}
		db       = l.svcCtx.DbEngin
		err      error
	)
	// 获取用户信息
	err = db.Where("name = ?", req.Name).Where("pass_word = ?", req.Password).Find(&userInfo).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if userInfo.ID == common.EmptyID {
		err = errors.New("账户和密码不匹配,请重试")
		log.Println(err)
		return nil, err
	}
	// 生成token
	tokenStr, err := middleware.GenerateToken(userInfo.ID, userInfo.Name)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	//
	value := LoginResp{UserInfo: userInfo, Token: tokenStr}
	return value, nil
}
