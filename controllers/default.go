package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Ctx.WriteString("hello world")

	c.Data["Website"] = "beego.me - mac"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
