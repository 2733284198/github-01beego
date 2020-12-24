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

func (c *MainController) Jj() {

	response := make(map[string]interface{})

	book1 := models.Book{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407}
	book2 := models.Book{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407}
	book3 := models.Book{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407}

	response["book1"] = book1
	response["book2"] = book2
	response["book3"] = book3

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
