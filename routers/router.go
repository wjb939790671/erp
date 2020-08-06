package routers

import (
	"erp/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/index", &controllers.MainController{})
	beego.Router("/", &controllers.MainController{})
	beego.Router("/loadmeun", &controllers.MainController{}, "post:LoadMenu")
	beego.Router("/login", &controllers.LoginController{}, "get:Index;post:Login")

	beego.Router("/employee/index", &controllers.EmployeeController{}, "get:Index")
	beego.Router("/employee/add", &controllers.EmployeeController{}, "get:Add;post:Add")

}
