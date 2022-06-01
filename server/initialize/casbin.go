package initialize

import (
	"admin/common"
	"github.com/casbin/casbin"
	"github.com/casbin/casbin/util"
	gormadapter "github.com/casbin/gorm-adapter"
	_ "github.com/go-sql-driver/mysql" //这一行代码千万不能取消，否则连接不上
	"strings"
)

func InitCasbin() {
	//Casbin连接数据库方式
	a := gormadapter.NewAdapterByDB(common.DB)
	e := casbin.NewEnforcer(common.CONFIG.Casbin.ModelPath, a)

	e.EnableLog(true)

	//将自定义函数添加到权限认证器中
	//自定义函数必须为以下类型，所以需要二次封装:  func(args ...interface{}) (interface{}, error)
	e.AddFunction("ParamsMatch", ParamsMatchFunc)

	//
	common.CASBIN = e
	return
}

func ParamsMatch(fullNameKey1 string, key2 string) bool {
	key1 := strings.Split(fullNameKey1, "?")[0]
	return util.KeyMatch2(key1, key2)
}

func ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return (bool)(ParamsMatch(name1, name2)), nil
}
