package logic

import (
	"context"
	"greet/services/user/userrpc/userrpc"
	"time"

	"greet/services/user/userapi/internal/svc"
	"greet/services/user/userapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (*types.LoginResponse, error) {
	user, err := l.svcCtx.UserRpc.GetUserByName(l.ctx, &userrpc.GetUserByNameReq{
		Username: req.Username,
	})
	if err != nil {
		return nil, err
	}

	if user.Password != hash(req.Password) {
		return nil, errors.New("invalid password")
	}

	now := time.Now().Unix()
	expire := l.svcCtx.Config.Auth.AccessExpire
	token, err := jwt.GenerateToken(user.Id, l.svcCtx.Config.Auth.AccessSecret, now, expire)
	if err != nil {
		return nil, err
	}
	return &types.LoginResponse{Token, token}, nil
}
