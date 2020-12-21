package admin

import (
	"github.com/astaxie/beego"
)

type ManagerController struct {
	beego.Controller
}

func (c *ManagerController) Get() {
	c.TplName = "admin/manager/index.html"
}

func (c *ManagerController) Add() {
	c.TplName = "admin/manager/add.html"
}

func (c *ManagerController) Edit() {
	c.TplName = "admin/manager/edit.html"
}
