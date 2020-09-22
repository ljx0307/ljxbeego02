package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "www.badu.com"
	c.Data["Email"] = "2632507285@qq.com"
	c.TplName = "index.tpl"
}
