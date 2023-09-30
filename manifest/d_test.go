package test

import (
	"context"

	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/frame/g"

	_ "github.com/gzdzh/dzhgo/contrib/drivers/mysql"

	"testing"
)

var (
	ctx context.Context
)

func Test(t *testing.T) {
	mInit := g.DB("default").Model("base_sys_init")
	value, _ := mInit.Clone().Where("group", "default").Where("id", 111).Value("table")

	g.Dump(value)

}
