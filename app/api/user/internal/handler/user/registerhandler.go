package user

import (
	"Go-SparrowRecSys/app/api/user/internal/logic/user"
	"Go-SparrowRecSys/app/api/user/internal/svc"
	"Go-SparrowRecSys/app/api/user/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			cookie := &http.Cookie{
				Name:   "accessToken",
				Value:  resp.AccessToken,
				Domain: svcCtx.Config.Domain,
				MaxAge: int(svcCtx.Config.JwtAuth.AccessExpire),
				Path:   "/", //不写这个cookie不存本地
			}
			http.SetCookie(w, cookie)
			httpx.OkJson(w, resp)
		}
	}
}
