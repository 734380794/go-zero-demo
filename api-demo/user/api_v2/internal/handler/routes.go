// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"go-zero-demo/api-demo/user/api_v2/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/users/login",
				Handler: loginHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/users/info",
				Handler: userInfoHandler(serverCtx),
			},
		},
	)
}
