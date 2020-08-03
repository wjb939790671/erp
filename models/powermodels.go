package models

import "time"

type Employee struct {
	Id             int
	Name           string
	Pwd            string
	EmployeeGroups []*EmployeeGroup `orm:"rel(m2m)"`
	Roles          []*Role          `orm:"rel(m2m)"`
}
type EmployeeGroup struct {
	Id        int
	GroupName string
	Describe  string

	Employees []*Employee `orm:"reverse(many)"`
	Roles     []*Role     `orm:"rel(m2m)"`
}

type Role struct {
	Id       int
	RoleName string
	Describe string

	Powers         []*Power         `orm:"rel(m2m)"`
	Employees      []*Employee      `orm:"reverse(many)"`
	EmployeeGroups []*EmployeeGroup `orm:"reverse(many)"`
}

type Power struct {
	Id         int
	PowerType  int         //1:Menu 2:PageElement 3:file
	Menu       *Menu       `orm:"rel(one)"`
	PageElemet *PageElemet `orm:"rel(one)"`
	File       *File       `orm:"rel(one)"`

	Operations []*Operation `orm:"rel(m2m)"`
	Roles      []*Role      `orm:"reverse(many)"`
}

type Menu struct {
	Id       int
	MenuName string
	MenuUrl  string
	ParentId int
	IsMenu   bool
	Power    *Power `orm:"reverse(one)"`
}
type PageElemet struct {
	Id             int
	PageElemetName string
	Power          *Power `orm:"reverse(one)"`
}

type File struct {
	Id       int
	FileName string
	FilePath string
	Power    *Power `orm:"reverse(one)"`
}

type Operation struct {
	Id            int
	OperationName string
	OperationCode string
	ParentId      int

	Powers []*Power `orm:"reverse(many)"`
}

type OperationLog struct {
	Id               int
	OperationTypeId  int
	OperationUserId  int
	OperationTime    time.Time
	OperationContent string
}

//menu tree
type MenuTree struct {
	Id         int                    `json:"id"`
	Text       string                 `json:"text"`
	IconCls    string                 `json:"iconCls"`
	State      string                 `json:"State"`
	Checked    bool                   `json:"checked"`
	Attributes map[string]interface{} `json:"attributes"`
	Children   []MenuTree             `json:"children"`
}
