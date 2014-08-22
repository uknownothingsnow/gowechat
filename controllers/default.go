package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	// this.Data["Catalogs"] = catalog.All()
	this.Data["PageTitle"] = "首页"
	this.Layout = "layout/default.html"
	this.TplNames = "index.html"
}

func (this *MainController) Post() {

}
