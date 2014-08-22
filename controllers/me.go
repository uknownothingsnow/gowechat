package controllers

type MeController struct {
	AdminController
}

func (this *MeController) Get() {
	this.Layout = "layout/admin.html"
	this.TplNames = "me/default.html"
}
