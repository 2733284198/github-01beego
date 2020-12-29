package beegoctl

import (
	"github.com/astaxie/beego"

	//"github.com/astaxie/beego/config/xml"

	"bee01/beegomodel"
	_ "bee01/beegomodel"
	"bee01/models"
)

// UserController operations for Wmfood
type ArticlesController struct {
	beego.Controller
}

func (c *ArticlesController) Get() {

	// 1. 查询一条数据

	//user := models.User{
	//	Id: 2,
	//}
	//models.Db.Find(&user)

	// 2. 查询所有数据

	//users := []models.User{}
	articles := []beegomodel.Articles{}
	models.Db.Debug().Find(&articles)
	//

	// 3. 各种查询
	//user := models.User{}
	//models.Db.Debug().First(&user) // 第一条
	//models.Db.Take(&user) // 第一条

	c.Data["json"] = articles
	c.ServeJSON()

	//c.TplName = "user/doAdd.tpl"
	//c.TplName = "user/pay.tpl"
	//c.TplName = "user/pay.xml"
	//c.TplName = "user/pay.html"
}
