package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["wechat/controllers:AdminController"] = append(beego.GlobalControllerRouter["wechat/controllers:AdminController"],
		beego.ControllerComments{
			"Get",
			"/",
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["wechat/controllers:MessageController"] = append(beego.GlobalControllerRouter["wechat/controllers:MessageController"],
		beego.ControllerComments{
			"Get",
			"/message/:id([0-9]+)",
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["wechat/controllers:MessageController"] = append(beego.GlobalControllerRouter["wechat/controllers:MessageController"],
		beego.ControllerComments{
			"Update",
			"/message/:id([0-9]+)",
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["wechat/controllers:MessageController"] = append(beego.GlobalControllerRouter["wechat/controllers:MessageController"],
		beego.ControllerComments{
			"Post",
			"/message",
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["wechat/controllers:MessageController"] = append(beego.GlobalControllerRouter["wechat/controllers:MessageController"],
		beego.ControllerComments{
			"List",
			"/message",
			[]string{"get"},
			nil})

}
