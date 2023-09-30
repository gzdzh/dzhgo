package main

import (
	_ "github.com/gzdzh/dzhgo/internal/packed"

	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	_ "github.com/gzdzh/dzhgo/contrib/drivers/mysql"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/gzdzh/dzhgo/internal/cmd"

	_ "github.com/gzdzh/dzhgo/contrib/files/oss"

	_ "github.com/gzdzh/dzhgo/contrib/files/local"

	_ "github.com/gzdzh/dzhgo/modules"
)

func main() {
	// gres.Dump()

	cmd.Main.Run(gctx.New())
}
