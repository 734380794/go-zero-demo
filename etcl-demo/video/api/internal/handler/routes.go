// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"go-zero-demo/etcl-demo/video/api/internal/svc"
	"net/http"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/videos/:id",
				Handler: getVideoHandler(serverCtx),
			},
		},
	)
}
