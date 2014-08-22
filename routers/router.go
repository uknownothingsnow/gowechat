package routers

import (
	"github.com/astaxie/beego"
	"wechat/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{}, "get:Login;post:DoLogin")
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")
	beego.Router("/me", &controllers.MeController{})
	ns := beego.NewNamespace("/admin", beego.NSInclude(&controllers.MessageController{}))
	beego.AddNamespace(ns)
}
