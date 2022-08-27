// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"5feet11/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/:snowflake",
				Handler: ExpandUrlHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/redirect",
				Handler: ShortenUrlHandler(serverCtx),
			},
		},
	)
}