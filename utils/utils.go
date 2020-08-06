package utils

import (
	"crypto/rand"
	"erp/models"
	"fmt"
	"log"
	"strconv"
)

const (
	LOGIN_OK   = 10000
	LOGIN_FAIL = 00000
	OK         = 00001
)

func GetCodeText(code int) string {
	return map[int]string{
		LOGIN_OK:   "登录成功",
		LOGIN_FAIL: "登录失败",
		OK:         "操作成功",
	}[code]
}

func Guid() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

}

//
func GetRoles(employee models.Employee) []models.Role {
	var roleList []models.Role
	db := models.DataBase{}
	//get role form role and add roleList
	db.M2MQuery(&employee, "Roles")
	for _, v := range employee.Roles {
		if !judgeRoleIsExist(roleList, *v) {
			roleList = append(roleList, *v)
		}
	}
	//get role from employeegroup and add roleList
	db.M2MQuery(&employee, "EmployeeGroups")
	for _, v := range employee.EmployeeGroups {
		db.M2MQuery(v, "Roles")
		for _, r := range v.Roles {
			if !judgeRoleIsExist(roleList, *r) {
				roleList = append(roleList, *r)
			}
		}
	}

	return roleList
}

func judgeRoleIsExist(roleList []models.Role, role models.Role) bool {
	for _, v := range roleList {
		if v.Id == role.Id {
			return true
		}
	}
	return false
}

//get main menu son menu tree
func GetMenuTree(parentId int, roleList []models.Role) []models.MenuTree {
	sql := "select m.id, m.menu_name,m.menu_url,m.is_menu  from (select r.role_id,r.power_id,p.menu_id from role_powers  as r inner join power as p on r.power_id=p.id) as a  inner join menu as m on a.menu_id=m.id where m.parent_id=" + strconv.Itoa(parentId)
	sql += " and a.role_id in ("
	for _, r := range roleList {
		sql += strconv.Itoa(r.Id) + ","
	}
	sql = sql[:len(sql)-1] + ")"
	var m []models.Menu
	db := models.DataBase{}
	db.Raw(sql, &m)
	var tree []models.MenuTree
	for _, v := range m {
		var tr models.MenuTree
		tr.Id = v.Id
		tr.Text = v.MenuName
		tr.Attributes = map[string]interface{}{"url": v.MenuUrl}
		if v.IsMenu {
			tr.Children = GetMenuTree(v.Id, roleList)
		}
		tree = append(tree, tr)
	}
	return tree
}

//get main menu
func GetMainMenu(parentId int, roleList []models.Role) []models.Menu {
	sql := "select m.id, m.menu_name,m.menu_url,m.is_menu  from (select r.role_id,r.power_id,p.menu_id from role_powers  as r inner join power as p on r.power_id=p.id) as a  inner join menu as m on a.menu_id=m.id where m.parent_id=" + strconv.Itoa(parentId)
	sql += " and m.is_menu=1 and a.role_id in ("
	for _, r := range roleList {
		sql += strconv.Itoa(r.Id) + ","
	}
	sql = sql[:len(sql)-1] + ")"
	var m []models.Menu
	db := models.DataBase{}
	db.Raw(sql, &m)
	return m
}
