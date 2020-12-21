package admin

import (
	"github.com/astaxie/beego"
)

type FocusController struct {
	beego.Controller
}

func (c *FocusController) Get() {
	c.Ctx.WriteString("后台管理系统的轮播图管理")
}
