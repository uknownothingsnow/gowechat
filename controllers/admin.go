package controllers

type AdminController struct {
	BaseController
}

func (this *AdminController) CheckLogin() {
	if !this.IsAdmin {
		this.Redirect("/login", 302)
	}
}

// @router / [get]
func (this *AdminController) Get() {
	this.Layout = "layout/admin.html"
	this.TplNames = "admin/index.html"
}
