package logic

import (
	"Go-SparrowRecSys/app/rpc/user/internal/svc"
	"Go-SparrowRecSys/app/rpc/user/model"
	"Go-SparrowRecSys/app/rpc/user/pb"
	"Go-SparrowRecSys/app/rpc/user/user"
	"Go-SparrowRecSys/pkg/xerr"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	userinfo, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetUserInfo find user db err , id:%d , err:%v", in.Id, err)
	}
	if userinfo == nil {
		return nil, errors.Wrapf(ErrUserNoExistsError, "id:%d", in.Id)
	}
	respUser := user.UserInfo{}
	_ = copier.Copy(&respUser, userinfo)
	return &pb.GetUserInfoResp{User: &respUser}, nil
}
