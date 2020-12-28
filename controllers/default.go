package controllers

import (
	"bee01/models"
	"fmt"
	"github.com/astaxie/beego"
	. "github.com/hunterhug/go_image"
	"github.com/skip2/go-qrcode"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) ScaleImg() {
	//图像位置
	filename := "static/upload/1.png"
	//保存位置
	savepath := "static/upload/1_3.png" //按照宽度进行等比例缩放
	err := ScaleF2F(filename, savepath, 400)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	} else {
		fmt.Printf("生成按宽度缩放图%s\n", savepath)
	}

	c.Ctx.WriteString("ScaleImg")
}

func (c *MainController) Qrcode() {

	//var png []byte
	//png, err := qrcode.Encode("https://www.baidu.com", qrcode.Medium, 256)

	err := qrcode.WriteFile("https://www.baidu.com", qrcode.Medium, 256, "static/upload/qrcode.png")

	if err != nil {
		beego.Error("生成二维码失败")
	} else {
		beego.Info("生成二维码成功")
	}

	//c.Data["png"] = png
	c.TplName = "index.tpl"
}

func (c *MainController) Get() {
	c.Ctx.WriteString("hello world")

	c.Data["Website"] = "beego.me - mac"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Product() {
	data := models.ProductsList()

	c.Data["json"] = data
	c.ServeJSON()
}

/* 二位map */
func (c *MainController) Jj2() {

	//var tags map[string] map[string] interface{}
	//var tags map[string] map[string]string
	//tags := make(map[string] map[string]string)

	//tags := make(map[string] map[string]string)

	//
	//tags := make(map[string] map[string]string)
	tags := make(map[string]map[string]interface{})

	tags["err"] = make(map[string]interface{})
	//tags["err"] = "jenny"
	tags["err"]["code"] = "200"
	tags["err"]["allpage"] = 2
	tags["err"]["curpage"] = 10

	//tags["no1"] = make(map[string]string)
	tags["no1"] = make(map[string]interface{})
	tags["no1"]["name"] = "jenny"
	tags["no1"]["sex"] = "man"
	tags["no1"]["address"] = "china"

	//tags["no1"]["no1"] = {
	//	"name":"n1",
	//	"age":10,
	//}

	c.Data["json"] = tags
	c.ServeJSON()
}

func (c *MainController) Jj1() {

	var tags map[string]interface{}
	tags = make(map[string]interface{})

	var tagsLocal map[string]string
	tagsLocal = make(map[string]string)
	var tagsTest map[string]string
	tagsTest = make(map[string]string)
	var tagsProduction map[string]string
	tagsProduction = make(map[string]string)

	tagsLocal["dev.www.9178.us"] = "dev.www.9178.us"
	tagsLocal["dev.static.9178.us"] = "dev.static.9178.us"

	tagsTest["dev.www.9178.us"] = "www.ninja911.com"
	tagsTest["dev.static.9178.us"] = "static.ninja911.com"

	tagsProduction["dev.www.9178.us"] = "ipx.www.ninja911.com"
	tagsProduction["dev.static.9178.us"] = "ipx.static.ninja911.com"

	tags["local"] = tagsLocal
	tags["test"] = tagsLocal
	tags["production"] = tagsLocal

	c.Data["json"] = tags
	c.ServeJSON()
}

func (c *MainController) Jj() {

	response := make(map[string]interface{})

	book1 := models.Book{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407}
	book2 := models.Book{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407}
	book3 := models.Book{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407}

	response["book1"] = book1
	response["book2"] = book2
	response["book3"] = book3

	// 嵌套json
	bodydata := `{
		  "version": {
			  "max": 3,
			  "last": "2016-03-11",
			  "detail": [
				  {
					  "time": "2016-03-12",
					  "ops": "add my email"
				  }
				 ]
			  }
		  }
  	`

	response["bodydata"] = bodydata

	//
	reqBody := []byte(`{"学校班级信息":
	{"班级信息":
		{"班级":"0212",
		"学号":"2002-2006"}
	}
	}`)
	response["reqBody"] = reqBody

	//response1 := make(map[string]   interface{map[string] interface{}} )
	//response1 = (
	//	"book1": models.Book{Title: "Go 语言", Author: "www.runoob.com", Subject: "Go 语言教程", Bookid: 6495407},
	//	"book2":models.Book{Title: "Go 语言", Author: "www.runoob.com", Subject: "Go 语言教程", Bookid: 6495407},
	//)
	//response["books"] = response1

	response["data"] = `{"name":"Xiao mi 6","product_id":"10","number":"10000","price":"2499","is_on_sale":"true"}`
	response["msg"] = "新增失败！"
	response["code"] = 500
	//response["err"] = err.Key + " " + err.Message

	c.Data["json"] = response
	c.ServeJSON()

}

func (c *MainController) Img() {
	//c.Ctx.WriteString("hello Image")

	//c.Data["Website"] = "beego.me - mac"
	//c.Data["Email"] = "astaxie@gmail.com"

	c.TplName = "index.html"
}
