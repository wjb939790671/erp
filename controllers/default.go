package controllers

import (
	"erp/bll"
	"erp/models"
	"erp/utils"

	"github.com/astaxie/beego/cache"

	"github.com/astaxie/beego"
)

type MainController struct {
	BaseController
}
type BaseController struct {
	beego.Controller
	cache    cache.Cache
	rmsg     ResultMessage
	employee models.Employee
}

func (this *MainController) Prepare() {
	if sessionId := this.GetSession("sessionid"); sessionId != nil {
		if bll.IsLogin(sessionId.(string)) {
			this.employee = bll.GetLoginEmployee(sessionId.(string))
			return
		}
	}
	this.Ctx.Redirect(302, "/login")
	this.Finish()
	return
}

func (c *MainController) Get() {
	m := bll.GetMainMenu(c.employee)
	//beego.Info(m)
	c.Data["Meun"] = m
	c.TplName = "index.html"
}
func (this *MainController) LoadMenu() {
	menuId, _ := this.GetInt("meunId")
	if menuId <= 0 {
		this.Finish()
		return
	}
	roleList := utils.GetRoles(this.employee)
	//this.rmsg.Code=utils.OK
	this.Data["json"] = utils.GetMenuTree(menuId, roleList)
	this.ServeJSON()

}
