package controllers

import (
	"bee01/models"
	"github.com/astaxie/beego"
	//"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
	"encoding/xml"
	"fmt"
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
	c.TplName = "food/doAdd.tpl"

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
	c.TplName = "user/pay.tpl"
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
	article := models.Article{}

	fmt.Printf("%s", article)

	//err := Unmarshaler
}

//func (c *UserController) doAdd() {
func (c *UserController) DoAdd() {

	//c.Ctx.WriteString("do add")

	p := models.Product{}

	err := c.ParseForm(&p)
	if err == nil {
		c.Ctx.WriteString("参数错误")
		return
	}

	//fmt.Println(p)
	fmt.Printf("%s", p)
	//beego.Info(p)

	//c.Ctx.WriteString(p)
	c.Ctx.WriteString("do add")

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
