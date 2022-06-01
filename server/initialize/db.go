package initialize

import (
	"admin/common"
	"admin/dao"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"os"
)

func InitDb() {
	switch common.CONFIG.System.DBType {
	case "mysql":
		InitMysql()
	default:
		InitMysql()
	}
}

func InitMysql() {
	sql := common.CONFIG.Mysql
	//fmt.Println(sql.Username + ":" + sql.Password + "@(" + sql.Path + ")/" + sql.Dbname + "?" + sql.Config)
	//root:414524@(localhost:3306)/ginadmin?charset=utf8mb4&parseTime=true&loc=Local&timeout=1s
	if db, err := gorm.Open("mysql", sql.Username+":"+sql.Password+"@("+sql.Path+")/"+sql.Dbname+"?"+sql.Config); err != nil {
		common.LOG.Error("DefaultDB 数据库启动异常", zap.Any("err", err))
		os.Exit(1)
	} else {
		common.DB = db
		db.DB().SetMaxIdleConns(sql.MaxIdleConns)
		db.DB().SetMaxOpenConns(sql.MaxOpenConns)
		db.LogMode(sql.LogMode)
		//禁止副表,即让grom转义struct名字的时候不用加上s
		db.SingularTable(true)
		DBTableMigrate()
	}
}

func DBTableMigrate() {
	common.DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").AutoMigrate(
		dao.SysRole{}, // 角色表
		dao.SysMenu{}, // 菜单表
		dao.SysUser{}, // 用户表
		dao.SysDept{},
		dao.SysOpLog{},
	)
	common.LOG.Info("register table success")
}
