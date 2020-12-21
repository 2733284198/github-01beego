package controllers

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
)

type NameController struct {
	beego.Controller
}

var cpt *captcha.Captcha

func init() {
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
	cpt.ChallengeNums = 4
	cpt.StdWidth = 100
	cpt.StdHeight = 40
}

func (c *NameController) Get() {
	c.Ctx.WriteString("Name-get")

	//c.Data["Website"] = "beego.me - mac"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
}

//func (c *NameController) Getone() {
func (c *NameController) DoLogin() {

	//captchaId := c.GetString("captchaId")
	captchaId := c.GetString("captcha_id")
	captchaValue := strings.Trim(c.GetString("captcha"), "")

	beego.Warning("=======>")
	beego.Warning(captchaId)
	beego.Warning(captchaValue)

	//if !cpt.VerifyReq(captchaId, captchaValue) {
	if !cpt.Verify(captchaId, captchaValue) {
		c.Ctx.WriteString("验证码失败")
	} else {
		c.Ctx.WriteString("验证码成功")
	}
}

func (c *NameController) DoLogin2() {

	//captchaId := c.GetString("captchaId")
	//captchaValue := strings.Trim(c.GetString("captcha"), "")

	if !cpt.VerifyReq(c.Ctx.Request) {
		c.Ctx.WriteString("验证码失败")
	} else {
		c.Ctx.WriteString("验证码成功")
	}

}

func (c *NameController) Login() {
	//c.Ctx.WriteString("Name-getone")

	//c.Data["Website"] = "Name - login"
	//c.Data["Email"] = "astaxie@gmail.com"

	c.TplName = "admin/login/login.html"
	//c.TplName = "user/doAdd.tpl"
}

//beego.Any("/foo",func(ctx *context.Context){
//	ctx.Output.Body([]byte("bar"))
//})
