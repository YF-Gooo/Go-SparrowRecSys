package user

import (
	"Go-SparrowRecSys/app/api/user/internal/logic/user"
	"Go-SparrowRecSys/app/api/user/internal/svc"
	"Go-SparrowRecSys/app/api/user/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			cookie := &http.Cookie{
				Name:   "accessToken",
				Value:  resp.AccessToken,
				Domain: svcCtx.Config.Domain,
				MaxAge: int(svcCtx.Config.JwtAuth.AccessExpire),
				Path:   "/",
			}
			http.SetCookie(w, cookie)
			httpx.OkJson(w, resp)
		}
	}
}
