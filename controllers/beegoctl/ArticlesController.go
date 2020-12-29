package beegoctl

import (
	"github.com/astaxie/beego"

	//"github.com/astaxie/beego/config/xml"

	"bee01/beegomodel"
	//_ "bee01/beegomodel/pay"
	//_ "bee01/beegomodel"
	"bee01/models"
)

// UserController operations for Wmfood
type ArticlesController struct {
	beego.Controller
}

func (c *ArticlesController) Getall() {
	var shops []beegomodel.Articles

	//models.Db.Debug().Table("articles").Select("title", "content").Scan(&shops).Limit(1)
	//models.Db.Debug().Table("articles").Select("title").Scan(&shops)

	models.Db.Debug().Select("title").Find(&shops)
	//models.Db.Debug().Select("title").Pluck("title", shops)

	c.Data["json"] = shops
	c.ServeJSON()
}

func (c *ArticlesController) Shopinfo() {

	//articles := []beegomodel.Articles{}

	var shops []beegomodel.Shop
	//var shops []beegomodel.Articles

	//models.Db.Debug().Select("id,shop_id").Find(&shops).Limit(2)
	//models.Db.Debug().Select("id,shop_id").Scan(&shops).Limit(2)
	//models.Db.Debug().Table("tb_wm_shop").Select([]string{"id", "shopname"}).Scan(&shops)

	//models.Db.Debug().Table("tb_wm_shop").Select("shop_id").Scan(&shops)

	models.Db.Debug().Table("articles").Select("title", "content").Scan(&shops).Limit(1)

	//models.Db.Debug().Select("id").Find(&shops)
	//models.Db.Select("id").Find(&shops)

	//models.Db.Debug().Select("id").Find(&shops)

	// Raw SQL
	//models.Db.Debug().Raw("SELECT 'id' FROM tb_wm_shop").Scan(&shops).Limit(2)

	//// SELECT name, age FROM users;ß

	//models.Db.Debug().Select([]string{"id", "shopname"}).Find(&shops)

	//// SELECT name, age FROM users;

	c.Data["json"] = shops
	c.ServeJSON()
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
}
