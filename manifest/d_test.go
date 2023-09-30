package test

import (
	"fmt"

	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gconv"

	_ "github.com/gzdzh/dzhgo/contrib/drivers/mysql"

	"testing"
)

var (
// ctx context.Context
)

func Test(t *testing.T) {
	rootPath := gfile.MainPkgPath()
	lockFile := rootPath + "/lock.json"
	fmt.Printf("lockFile: %v\n", lockFile)

	modulesName := "dzhMember"
	searchName := "task"

	if gfile.Exists(lockFile) {

		// 存在文件，
		fileInfo := gfile.GetContents(lockFile)
		j, _ := gjson.LoadContent(fileInfo)

		fileInfoMap := gmap.NewStrAnyMapFrom(j.Map())
		arr := fileInfoMap.Get(modulesName)

		fileInfoArr := garray.NewStrArrayFrom(gconv.SliceStr(arr))

		// 找是否更新过指定模块表
		if fileInfoArr.Contains(searchName) {

			fmt.Println("存在，跳过")

		} else {
			fileInfoMap.Remove(modulesName)
			fileInfoArr.Append(searchName)
			fileInfoMap.Set(modulesName, fileInfoArr)
			gfile.PutContents(lockFile, fileInfoMap.String())
		}

	} else {
		insertMap := g.MapStrAny{
			modulesName: g.Slice{searchName},
		}
		gfile.PutContentsAppend(lockFile, gjson.MustEncodeString(insertMap))

	}
}
