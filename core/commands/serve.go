package commands

import (
	"github.com/jiangrx816/wechat/http/router"

	"github.com/jiangrx816/wechat/core/graceful"

	"github.com/jiangrx816/wechat/core/server"

	"github.com/urfave/cli/v2"
)

func Serve(c *cli.Context) error {
	// 运行HTTP服务
	graceful.Start(server.NewHttp(server.Addr(":8080"), server.Router(router.All())))

	graceful.Wait()
	return nil
}
