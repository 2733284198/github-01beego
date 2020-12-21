package admin

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Success(message string, redirect string) {
	c.Data["message"] = message
	c.Data["redirect"] = redirect
	c.TplName = "admin/public/success.html"
}

func (c *BaseController) Error(message string, redirect string) {
	c.Data["message"] = message
	c.Data["redirect"] = redirect
	c.TplName = "admin/public/error.html"
}
