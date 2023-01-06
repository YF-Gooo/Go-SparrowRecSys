package logic

import "Go-SparrowRecSys/pkg/xerr"

var ErrGenerateTokenError = xerr.NewErrMsg("生成token失败")
var ErrUsernamePwdError = xerr.NewErrMsg("账号或密码不正确")
var ErrUserAlreadyRegisterError = xerr.NewErrMsg("user has been registered")
var ErrUserNoExistsError = xerr.NewErrMsg("用户不存在")
