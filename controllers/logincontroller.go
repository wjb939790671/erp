package controllers

import (
	"erp/common"
	"erp/models"
	"erp/utils"
	"time"

	"github.com/astaxie/beego"
)

type LoginController struct {
	rmsg ResultMessage
	beego.Controller
}
type ResultMessage struct {
	Code int         `json:"code"`
	Text string      `json:"text"`
	Data interface{} `json:"data"`
}

func (c *LoginController) Index() {
	c.TplName = "login/login.html"
}
func (c *LoginController) Login() {
	name := c.GetString("Name")
	pwd := c.GetString("Pwd")
	if len(name) > 0 || len(pwd) > 0 {
		db := &models.DataBase{}
		employee := models.Employee{}
		employee.Name = name
		if error := db.QueryOne(&employee, "Name"); error == nil {
			if employee.Pwd == pwd {
				c.rmsg.Code = utils.LOGIN_OK
				guid := utils.Guid()
				common.Bm.Put(guid, employee, 2*60*time.Second)
				c.SetSession("sessionid", guid)
				c.rmsg.Data = "/index"

			}
		}
	} else {
		c.rmsg.Code = utils.LOGIN_FAIL
	}
	c.rmsg.Text = utils.GetCodeText(c.rmsg.Code)
	c.Data["json"] = c.rmsg
	c.ServeJSON()
}
