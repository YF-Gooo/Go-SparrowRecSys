package logic

import (
	"Go-SparrowRecSys/app/rpc/user/model"
	"Go-SparrowRecSys/app/rpc/user/user"
	"Go-SparrowRecSys/pkg/tool"
	"Go-SparrowRecSys/pkg/xerr"
	"context"
	"github.com/pkg/errors"

	"Go-SparrowRecSys/app/rpc/user/internal/svc"
	"Go-SparrowRecSys/app/rpc/user/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	var userId int64
	var err error
	userId, err = l.loginByNickName(in.Nickname, in.Password, in.Gcode)
	if err != nil {
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
	return &user.LoginResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil
	return &pb.LoginResp{}, nil
}

func (l *LoginLogic) loginByNickName(nickname, password, gcode string) (int64, error) {
	user, err := l.svcCtx.UserModel.FindOneByNickname(l.ctx, nickname)
	if err != nil && err != model.ErrNotFound {
		return 0, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "根据账户查询用户信息失败，nickname:%s,err:%v", nickname, err)
	}
	if user == nil {
		return 0, errors.Wrapf(ErrUserNoExistsError, "nickname:%s", nickname)
	}
	if !(tool.Md5Crypt(password, l.svcCtx.Config.JwtAuth.AccessSecret) == user.Password) {
		return 0, errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "密码匹配出错")
	}
	if ok, _ := tool.NewGoogleAuth().VerifyCode(user.Googleauth, gcode); !ok {
		return 0, errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "动态口令匹配出错")
	}
	return user.Id, nil
}
