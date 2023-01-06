package svc

import (
	"Go-SparrowRecSys/app/rpc/user/internal/config"
	"Go-SparrowRecSys/app/rpc/user/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	UserModel     model.UserModel
	UserRoleModel model.UserRoleModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:        c,
		UserModel:     model.NewUserModel(sqlConn, c.Cache),
		UserRoleModel: model.NewUserRoleModel(sqlConn, c.Cache),
	}
}
