package logic

import (
	"context"

	"greet/services/user/userrpc/internal/svc"
	"greet/services/user/userrpc/userrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *userrpc.Request) (*userrpc.Response, error) {
	// todo: add your logic here and delete this line

	return &userrpc.Response{}, nil
}
