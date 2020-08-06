package controllers

type EmployeeController struct {
	BaseController
}

func (this *EmployeeController) Index() {
	this.TplName = "employee/index.html"
}
func (this *EmployeeController) Add() {
	this.TplName = "employee/add.html"
}
