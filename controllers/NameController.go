package controllers

import (
	"github.com/astaxie/beego"
)

type NameController struct {
	beego.Controller
}

func (c *NameController) Get() {
	c.Ctx.WriteString("Name-get")

	//c.Data["Website"] = "beego.me - mac"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
}

//func (c *NameController) Getone() {
func (c *NameController) GetOne() {
	c.Ctx.WriteString("Name-getone")

	//c.Data["Website"] = "beego.me - mac"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
}

//beego.Any("/foo",func(ctx *context.Context){
//	ctx.Output.Body([]byte("bar"))
//})
