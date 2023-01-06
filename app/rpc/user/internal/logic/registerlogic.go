package logic

import (
	"Go-SparrowRecSys/app/rpc/user/model"
	"Go-SparrowRecSys/app/rpc/user/user"
	"Go-SparrowRecSys/pkg/tool"
	"Go-SparrowRecSys/pkg/xerr"
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"Go-SparrowRecSys/app/rpc/user/internal/svc"
	"Go-SparrowRecSys/app/rpc/user/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {
	userinfo, err := l.svcCtx.UserModel.FindOneByNickname(l.ctx, in.Nickname)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Nickname:%s,err:%v", in.Nickname, err)
	}
	if userinfo != nil {
		return nil, errors.Wrapf(ErrUserAlreadyRegisterError, "Register user exists Nickname:%s,err:%v", in.Nickname, err)
	}
	//创建google验证
	secret := tool.NewGoogleAuth().GetSecret()
	var userId int64
	if err := l.svcCtx.UserModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		userdb := new(model.User)
		userdb.Nickname = in.Nickname
		if len(in.Password) > 0 {
			userdb.Password = tool.Md5Crypt(in.Password, l.svcCtx.Config.JwtAuth.AccessSecret)
		}
		userdb.Googleauth = secret
		insertResult, err := l.svcCtx.UserModel.Insert(ctx, session, userdb)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Register db user Insert err:%v,user:%+v", err, userdb)
		}
		lastId, err := insertResult.LastInsertId()
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Register db user insertResult.LastInsertId err:%v,user:%+v", err, userdb)
		}
		userId = lastId
		userRole := new(model.UserRole)
		userRole.UserId = lastId
		userRole.RoleId = 0
		if _, err := l.svcCtx.UserRoleModel.Insert(ctx, session, userRole); err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Register db user_auth Insert err:%v,userAuth:%v", err, userRole)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	//2、Generate the token, so that the service doesn't call rpc internally
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&user.GenerateTokenReq{
		UserId: userId,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "GenerateToken userId : %d", userId)
	}
	return &user.RegisterResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
		GoogleAuth:   secret,
	}, nil
}
