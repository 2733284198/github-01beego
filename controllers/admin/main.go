package admin

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "admin/main/index.html"
}

func (c *MainController) Welcome() {
	c.TplName = "admin/main/welcome.html"
}
