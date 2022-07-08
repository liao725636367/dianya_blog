package cmd

import (
	"blog/internal/service"
	"context"

	"blog/internal/controller"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(service.Middleware().MiddlewareHandlerResponse)
				group.Middleware(service.Middleware().MiddlewareCORS)
				group.Middleware(service.Middleware().Ctx)
				group.Bind(
					controller.Article,
					controller.User,
					controller.Comment,
					controller.Category,
				)
			})
			s.Run()
			return nil
		},
	}
)
