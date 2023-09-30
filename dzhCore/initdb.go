package dzhCore

import (
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gres"
	"github.com/gzdzh/dzhgo/dzhCore/coredb"
	"gorm.io/gorm"
)

// 初始化数据库连接供gorm使用
func InitDB(group string) (*gorm.DB, error) {
	// var ctx context.Context
	var db *gorm.DB
	// 如果group为空，则使用默认的group，否则使用group参数
	if group == "" {
		group = "default"
	}
	defer func() {
		if err := recover(); err != nil {
			panic("failed to connect database")
		}
	}()
	config := g.DB(group).GetConfig()
	db, err := coredb.GetConn(config)
	if err != nil {
		panic(err.Error())
	}

	GormDBS[group] = db
	return db, nil
}

// 根据entity结构体获取 *gorm.DB
func getDBbyModel(model IModel) *gorm.DB {

	group := model.GroupName()
	// 判断是否存在 GormDBS[group] 字段，如果存在，则使用该字段的值作为DB，否则初始化DB
	if _, ok := GormDBS[group]; ok {
		return GormDBS[group]
	} else {

		db, err := InitDB(group)
		if err != nil {
			panic("failed to connect database")
		}
		// 把重新初始化的GormDBS存入全局变量中
		GormDBS[group] = db
		return db
	}
}

// 根据entity结构体创建表
func CreateTable(model IModel) error {
	if Config.AutoMigrate {
		db := getDBbyModel(model)
		return db.AutoMigrate(model)
	}
	return nil
}

// FillInitData 数据库填充初始数据
/**
* reset含true，二次更新表，追加数据
 */
func FillInitData(ctx g.Ctx, moduleName string, model IModel, reset ...bool) error {

	// 是否重写已经初始化过的表
	shouldReset := false
	if len(reset) > 0 {
		shouldReset = reset[0]
	}

	mInit := g.DB("default").Model("base_sys_init")
	n, err := mInit.Clone().Where("group", model.GroupName()).Where("table", model.TableName()).Count()
	if err != nil {
		g.Log().Error(ctx, "读取表 base_sys_init 失败 ", err.Error())
		return err
	}
	if !shouldReset {
		if n > 0 {
			g.Log().Debug(ctx, "分组", model.GroupName(), "中的表", model.TableName(), "已经初始化过,跳过本次初始化.")
			return nil
		}
	}

	// 根目录锁文件模块名下存在指定的表名，证明该模块已经初始化了，跳过二次初始化
	rootPath := gfile.MainPkgPath()
	lockFile := rootPath + "/lock.json"
	g.Log().Debug(ctx, "lockFile", lockFile)
	if gfile.Exists(lockFile) {

		// 存在文件，
		fileInfo := gfile.GetContents(lockFile)

		if fileInfo == "" {

			if err = updateData(ctx, mInit, moduleName, model, shouldReset); err == nil {
				insertMap := g.MapStrAny{
					moduleName: g.Slice{model.TableName()},
				}
				gfile.PutContents(lockFile, gjson.MustEncodeString(insertMap))
				g.Log().Debug(ctx, "分组", model.GroupName(), "中的表", model.TableName(), "第一次lock.json写入")

			}

		} else {

			jsonObj, _ := gjson.LoadContent(fileInfo)

			fileInfoMap := gmap.NewStrAnyMapFrom(jsonObj.Map())
			arr := fileInfoMap.Get(moduleName)

			fileInfoArr := garray.NewStrArrayFrom(gconv.SliceStr(arr))

			// 找是否更新过指定模块表
			if fileInfoArr.Contains(model.TableName()) {

				g.Log().Debug(ctx, "分组", model.GroupName(), "中的表", model.TableName(), "存在lock,跳过本次初始化.")
				return nil

			} else {

				if err = updateData(ctx, mInit, moduleName, model, shouldReset); err == nil {
					fileInfoMap.Remove(moduleName)
					fileInfoArr.Append(model.TableName())
					fileInfoMap.Set(moduleName, fileInfoArr)
					gfile.PutContents(lockFile, fileInfoMap.String())

					g.Log().Debug(ctx, "分组", model.GroupName(), "中的表", model.TableName(), "lock.json更新")
				}

			}

		}

	} else {

		if err = updateData(ctx, mInit, moduleName, model, shouldReset); err == nil {
			insertMap := g.MapStrAny{
				moduleName: g.Slice{model.TableName()},
			}
			gfile.PutContents(lockFile, gjson.MustEncodeString(insertMap))
			g.Log().Debug(ctx, "分组", model.GroupName(), "中的表", model.TableName(), "第一次lock.json写入")

		}
	}

	g.Log().Info(ctx, "分组", model.GroupName(), "中的表", model.TableName(), "初始化完成.")
	return nil
}

func updateData(ctx g.Ctx, mInit *gdb.Model, moduleName string, model IModel, shouldReset bool) error {

	m := g.DB(model.GroupName()).Model(model.TableName())
	path := "modules/" + moduleName + "/resource/initjson/" + model.TableName() + ".json"
	jsonData, _ := gjson.LoadContent(gres.GetContent(path))

	g.Log().Debugf(ctx, "model.TableName(): %v,path:%v", model.TableName(), path)

	if jsonData.Var().Clone().IsEmpty() && shouldReset {
		g.Log().Debug(ctx, "分组", model.GroupName(), "中的表", model.TableName(), "无可用的初始化数据,跳过本次初始化. jsonData:", jsonData)
		return gerror.New("无可用的初始化数据,跳过本次初始化")
	}
	_, err := m.Data(jsonData).Insert()
	if err != nil {
		g.Log().Error(ctx, err.Error())
		return err
	}
	if !shouldReset {
		_, err = mInit.Insert(g.Map{"group": model.GroupName(), "table": model.TableName()})
		if err != nil {
			g.Log().Error(ctx, err.Error())
			return err
		}
	}

	return nil
}
