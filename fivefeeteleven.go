package main

import (
	"flag"
	"net/http"

	"5feet11/internal/config"
	"5feet11/internal/errorx"
	"5feet11/internal/handler"
	"5feet11/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/fivefeeteleven-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		switch e := err.(type) {
		case *errorx.CodeError:
			return http.StatusBadRequest, e.Data()
		default:
			return http.StatusInternalServerError, nil
		}
	})

	logx.Infof("Starting server at http://%s:%d", c.Host, c.Port)
	server.Start()
}
