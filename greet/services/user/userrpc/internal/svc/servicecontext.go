package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"greet/services/user/userrpc/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource) // 连接数据库
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(conn, c.Cache.Host), // 创建 model
	}
}
