package routers

import (
	"bee01/controllers/index"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &index.IndexController{})
	beego.Router("/login", &index.LoginController{})
}
