package handler

import (
	"go-zero-demo/common/response"
	"net/http"

	"go-zero-demo/api-demo/user/api_v2/internal/logic"
	"go-zero-demo/api-demo/user/api_v2/internal/svc"
)

func userInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		/*if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}*/
		response.Response(r, w, resp, err)
	}
}
