package bll

import (
	"erp/common"
	"erp/models"
	"erp/utils"
)

// judge login or unlogin
//if sessionid Isexist login or unlogin
func IsLogin(sessionId string) bool {
	return common.Bm.IsExist(sessionId)
}

//
func GetLoginEmployee(sessionId string) models.Employee {
	return common.Bm.Get(sessionId).(models.Employee)
}

// judge current visit is permission or not
//@pram employee current login
//@pram url current visit url
func IsPermission(employee models.Employee, url string) bool {
	db := models.DataBase{}
	menu := models.Menu{MenuUrl: url}
	if error := db.QueryOne(&menu, "MenuUrl"); error != nil {
		return false
	}
	db.M2MQuery(&employee, "Roles")
	for _, v := range employee.Roles {
		db.M2MQuery(v, "Powers")
		for _, p := range v.Powers {
			if p.Menu.Id == menu.Id {
				return true
			}
		}

	}
	return false

}

func GetMainMenu(employee models.Employee) []models.Menu {
	roleList := utils.GetRoles(employee)

	// sql := "select m.id, m.menu_name,m.menu_url,m.parent_id,a.power_id  from (select r.role_id,r.power_id,p.menu_id from role_powers  as r inner join power as p on r.power_id=p.id) as a  inner join menu as m on a.menu_id=m.id where m.parent_id=0 and m.is_menu=1 and a.role_id in ("
	// for _, r := range roleList {
	// 	sql += strconv.Itoa(r.Id) + ","
	// }
	// sql = sql[:len(sql)-1] + ")"
	// var m []orm.Params
	// o := orm.NewOrm()
	// o.Raw(sql).Values(&m)

	return utils.GetMainMenu(0, roleList)
}
func judgeIsExistRole(roleList []*models.Role, role *models.Role) bool {
	for _, v := range roleList {
		if v.Id == role.Id {
			return true
		}
	}
	return false
}
