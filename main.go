package main

import (
	"erp/common"
	"erp/models"
	_ "erp/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	beego.Run()

}

func init() {
	common.Bm, _ = cache.NewCache("memory", `{"interval":360}`)
	orm.RegisterDriver("mysql", orm.DRMySQL)

	//orm.RegisterDataBase("default", "mysql", "root:root@/orm_test?charset=utf8")
	// set default database
	// 参数4(可选)  设置最大空闲连接
	// 参数5(可选)  设置最大数据库连接 (go >= 1.2)
	maxIdle := 30
	maxConn := 30
	orm.RegisterDataBase("default", "mysql", "root:123@tcp(127.0.0.1:3306)/erp?charset=utf8", maxIdle, maxConn)

	// register model
	orm.RegisterModel(new(models.Employee), new(models.Role), new(models.Power), new(models.EmployeeGroup),
		new(models.PageElemet), new(models.OperationLog), new(models.Operation), new(models.Menu),
		new(models.File))

	// create table
	orm.RunSyncdb("default", false, true)
}
