// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"qq-bot-307/service/http/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/",
				Handler: MessageHandler(serverCtx),
			},
		},
	)
}