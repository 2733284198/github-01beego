package admin

import (
	_ "bee01/controllers"
	_ "github.com/astaxie/beego"
)

type LoginController struct {
	//beego.Controller
	BaseController
}

func (c *LoginController) Get() {
	c.TplName = "admin/login/login.html"
}
