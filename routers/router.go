package routers

import (
	"bee04/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/register",&controllers.RegisterController{})
    beego.Router("/ ", &controllers.MainController{})
}
