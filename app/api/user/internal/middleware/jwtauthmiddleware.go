package middleware

import (
	"Go-SparrowRecSys/pkg/response"
	"github.com/zeromicro/go-zero/rest/handler"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type JwtAuthMiddleware struct {
	secret string
}

func NewJwtAuthMiddleware(secret string) *JwtAuthMiddleware {
	return &JwtAuthMiddleware{
		secret: secret,
	}
}

func (m *JwtAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if cookie, err := r.Cookie("accessToken"); err == nil {
			r.Header.Set("Authorization", cookie.Value)
		}
		if len(r.Header.Get("Authorization")) > 0 {
			//has jwt Authorization
			authHandler := handler.Authorize(m.secret)
			authHandler(next).ServeHTTP(w, r)
			return
		} else {
			//no jwt Authorization
			httpx.WriteJson(w, http.StatusUnauthorized, response.Error(401, "jwt鉴权失败"))
			return
		}
	}
}
