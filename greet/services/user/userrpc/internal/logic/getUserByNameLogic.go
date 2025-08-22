package logic

import "greet/services/user/userrpc/internal/svc"

type GetUserByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserByNameLogic()
