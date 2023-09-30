package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gzdzh/dzhgo/dzhCore"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			if dzhCore.IsRedisMode {
				go dzhCore.ListenFunc(ctx)
			}

			s := g.Server()
			s.Run()
			return nil
		},
	}
)
