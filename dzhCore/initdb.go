package dzhCore

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
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

// 数据库填充初始数据
func FillInitData(ctx g.Ctx, moduleName string, model IModel) error {

	mInit := g.DB("default").Model("base_sys_init")
	value, err := mInit.Clone().Where("group", model.GroupName()).Where("module", moduleName).Value("tables")
	if err != nil {
		g.Log().Error(ctx, "读取表 base_sys_init 失败 ", err.Error())
		return err
	}

	// 第一次
	if value.IsEmpty() {

		if err = updateData(ctx, mInit, moduleName, model); err == nil {
			_, err = mInit.Insert(g.Map{"group": model.GroupName(), "module": moduleName, "tables": model.TableName()})
			if err != nil {
				g.Log().Error(ctx, err.Error())
				return err
			}
			g.Log().Debug(ctx, "分组", model.GroupName(), "中的表", model.TableName(), "第一次写入")

		}
		return nil
	}

	tableArr := strings.Split(value.String(), ",")
	tableGarr := garray.NewStrArrayFrom(tableArr)
	if tableGarr.Contains(model.TableName()) {
		g.Log().Debug(ctx, "分组", model.GroupName(), "中的表", model.TableName(), "已经初始化过,跳过本次初始化.")
		return nil
	}

	if err = updateData(ctx, mInit, moduleName, model); err == nil {

		tableGarr.Append(model.TableName())
		str := strings.Join(tableGarr.Slice(), ",")
		mInit.Where("group", model.GroupName()).Where("module", moduleName).Data(g.Map{"tables": str}).Update()
		g.Log().Debug(ctx, "分组", model.GroupName(), "中的表", model.TableName(), "更新写入")

	}

	g.Log().Info(ctx, "分组", model.GroupName(), "中的表", model.TableName(), "初始化完成.")
	return nil
}

// 更新文件
func updateData(ctx g.Ctx, mInit *gdb.Model, moduleName string, model IModel) error {

	m := g.DB(model.GroupName()).Model(model.TableName())
	path := "modules/" + moduleName + "/resource/initjson/" + model.TableName() + ".json"
	jsonData, _ := gjson.LoadContent(gres.GetContent(path))

	g.Log().Debugf(ctx, "model.TableName(): %v,path:%v", model.TableName(), path)

	if jsonData.Var().Clone().IsEmpty() {
		g.Log().Debug(ctx, "分组", model.GroupName(), "中的表", model.TableName(), "无可用的初始化数据,跳过本次初始化. jsonData:", jsonData)
		return gerror.New("无可用的初始化数据,跳过本次初始化")
	}
	_, err := m.Data(jsonData).Insert()
	if err != nil {
		g.Log().Error(ctx, err.Error())
		return err
	}

	return nil
}

// 当前文件运行路径
func GetRunPath(ctx g.Ctx) (string, error) {

	// 获取当前文件的路径
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		err := gerror.New("Failed to GetRunPath")
		return "", err
	}

	// 获取项目的根目录路径
	rootPath := filepath.Dir(filepath.Dir(filename))

	return rootPath, nil
}

// 获取当前项目根目录路径
func GetRootPath(ctx g.Ctx) string {

	wd, _ := os.Getwd()

	return wd

	// filePath, _ := exec.LookPath(os.Args[0])
	// absFilePath, _ := filepath.Abs(filePath)
	// rootDir := path.Dir(absFilePath)

	// return rootDir
}
