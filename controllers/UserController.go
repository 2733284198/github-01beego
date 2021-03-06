package controllers

import (
	"bee01/models"
	"github.com/astaxie/beego"

	//"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
	"encoding/xml"
	"fmt"
	"strconv"
	//"github.com/astaxie/beego/config/xml"
)

// UserController operations for Wmfood
type UserController struct {
	beego.Controller
}

// URLMapping ...
func (c *UserController) URLMapping() {
	c.Mapping("Get", c.Get)
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

//func (c *MainController) Get() {
func (c *UserController) Get() {
	id := c.GetString("id")
	beego.Info(id)

	c.Data["title"] = "你好beego getone === "
	c.TplName = "user/doAdd.tpl"

	//c.Ctx.WriteString("food world 1")

}

// Post ...
// @Title Create
// @Description create Wmfood
// @Param	body		body 	models.Wmfood	true		"body for Wmfood content"
// @Success 201 {object} models.Wmfood
// @Failure 403 body is empty
// @router / [post]

// pay,xml支付宝，微信返回类型
func (c *UserController) Pay() {
	c.TplName = "user/doAdd.tpl"
	//c.TplName = "user/pay.tpl"
	//c.TplName = "user/pay.xml"
	//c.TplName = "user/pay.html"
}

/******************************
 gorm:begin
*******************************/

func (c *UserController) Gedit() {

	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	beego.Warning(id)

	user := models.User{
		//Id: 2,
		Id: id,
	}
	models.Db.Debug().Find(&user)

	user.Username = "user222"
	models.Db.Save(&user)

	models.Db.Debug().Find(&user)

	c.Data["json"] = user
	c.ServeJSON()
}

func (c *UserController) Gadd() {
	user := models.User{
		Username: "manlan1",
		Password: "m1p1",
		Age:      33,
	}

	models.Db.Debug().Create(&user)
	//c.Gormget()

	//c.Ctx.WriteString("添加数据成功")

	// 2. 查询所有数据
	users := []models.User{}
	models.Db.Debug().Find(&users)

	c.Data["json"] = users
	c.ServeJSON()
}

func (c *UserController) GormWhere() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	beego.Warning(id)

	//var n1 = 5
	var n1 = id
	var users []models.User
	//models.Db.Debug().Where("id > ?", n1).Find(&users) // 第一条
	models.Db.Debug().Where("id>?", n1).Find(&users) // 第一条

	//models.Db.Take(&user) // 第一条

	c.Data["json"] = users
	c.ServeJSON()

}

func (c *UserController) Gormget() {

	// 1. 查询一条数据

	//user := models.User{
	//	Id: 2,
	//}
	//models.Db.Find(&user)

	// 2. 查询所有数据

	users := []models.User{}
	models.Db.Debug().Find(&users)
	//

	// 3. 各种查询
	user := models.User{}
	models.Db.Debug().First(&user) // 第一条
	//models.Db.Take(&user) // 第一条

	c.Data["json"] = users
	c.ServeJSON()

	//c.TplName = "user/doAdd.tpl"
	//c.TplName = "user/pay.tpl"
	//c.TplName = "user/pay.xml"
	//c.TplName = "user/pay.html"
}

/******************************
gorm:end
*******************************/

func (c *UserController) Testsession(string, string) {
	c.SetSession("username", "manlan1")

	session1 := c.GetSession("username")

	beego.Info(session1)

	fmt.Println(session1)
	c.Ctx.WriteString("session1")
}

func (c *UserController) Jump() {
	//c.Redirect("www.baidu.com", 303)
	//c.Redirect("/user/pay", 200)
	c.Redirect("www.baidu.com", 200)

	//c.TplName = "user/pay.xml"
	//c.TplName = "user/pay.html"
}

// 处理xml内容，支付宝，微信返回结果处理

/* postman: 发送 post, xml的内容，
	http://localhost:8089/user/xml

	<?xml version="1.0" encoding="UTF-8"?>
<article>
    <title type="string"> 标题 </title>
    <content type="string"> 我是张三的内容 </content>
</article>

*/
func (c *UserController) Xml() {
	xmlstr := string(c.Ctx.Input.RequestBody)

	beego.Info(xmlstr)
	//c.Ctx.WriteString(xmlstr)

	p := models.Article{}
	err := xml.Unmarshal(c.Ctx.Input.RequestBody, &p)

	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()

		//fmt.Println("xml参数错误")
	} else {
		c.Data["json"] = p
		c.ServeJSON()
	}
}

func (c *UserController) PayCallback() {
	user := models.User{}

	fmt.Printf("%s", user)

	//err := Unmarshaler
}

//func (c *UserController) doAdd() {
func (c *UserController) AddForm() {
	c.Data["title"] = "AddForm"
	c.TplName = "user/doAdd.tpl"
}

func (c *UserController) Doadd() {
	u := models.User{}

	err := c.ParseForm(&u)
	if err != nil {
		c.Ctx.WriteString("参数错误")
		return
	}

	fmt.Printf("%s", u)
	c.Data["json"] = u
	c.ServeJSON()
}

// GetOne ...
// @Title GetOne
// @Description get Wmfood by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Wmfood
// @Failure 403 :id is empty
// @router /:id [get]
func (c *UserController) GetOne() {

	//c.Data["title"] = "你好beego getone"

	//c.TplName = "food/doAdd.tpl"

	u := models.User{
		Username: "张三",
		Password: "123456",
	}

	c.Data["json"] = u
	c.ServeJSON()
}

func (c *UserController) Getuid() {
	//uid := c.Ctx.Input.Param(":id")
	uid := c.Ctx.Input.Param(":id") // :id，要加签名的:

	//u := models.User{
	//	Username: "张三",
	//	Password: "123456",
	//}

	// todo:返回 json的值
	c.Ctx.WriteString(uid)

	//c.Data["json"] = uid
	//c.ServeJSON()
}

func (c *UserController) Getuidhtml() {
	//uid := c.Ctx.Input.Param(":id")
	uid := c.Ctx.Input.Param(":id") // :id，要加签名的:

	// todo:返回 json的值
	c.Ctx.WriteString(uid)
}

func (c *UserController) AddFood() {

	c.Data["title"] = "AddFood"

	c.TplName = "food/doAdd.tpl"
}

// GetAll ...
// @Title GetAll
// @Description get Wmfood
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Wmfood
// @Failure 403
// @router / [get]
func (c *UserController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Wmfood
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Wmfood	true		"body for Wmfood content"
// @Success 200 {object} models.Wmfood
// @Failure 403 :id is not int
// @router /:id [put]
func (c *UserController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Wmfood
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *UserController) Delete() {

}
