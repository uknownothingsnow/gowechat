package controllers

import (
	// "github.com/astaxie/beego"
	// "wechat/models"
	// "strconv"
	"wechat/models/message"
)

const (
	PAGE_SIZE = 10
)

type MessageController struct {
	BaseController
}

// @router /message/:id([0-9]+) [get]
func (this *MessageController) Get() {
	messageTemplateId, err := this.GetInt(":id")
	if err != nil {
		this.Ctx.WriteString("id param does not exist")
		return
	}
	msg := message.GetById(int64(messageTemplateId))
	if msg == nil {
		this.Ctx.WriteString("id does not exist")
		return
	}
	this.Data["json"] = msg
	this.ServeJson()
}

// @router /message/:id([0-9]+) [put]
func (this *MessageController) Update() {
	messageTemplateId, err := this.GetInt(":id")
	if err != nil {
		this.Ctx.Output.SetStatus(401)
		return
	}
	text := this.GetString("text")
	msg := message.GetById(int64(messageTemplateId))
	if msg == nil {
		this.Ctx.WriteString("id does not exist")
		return
	}

	message.Update(msg, text)

	this.Data["json"] = msg
	this.ServeJson()
}

// @router /message [post]
func (this *MessageController) Post() {
	text := this.GetString("text")
	id, err := message.Save(text)
	if err != nil {
		this.Ctx.Output.SetStatus(500)
		return
	}
	this.Data["json"] = message.GetById(id)
	this.ServeJson()
}

// @router /message [get]
func (this *MessageController) List() {
	page := this.GetIntWithDefault("page", 0)
	pageSize := this.GetIntWithDefault("pageSize", PAGE_SIZE)
	messages := message.List(page*pageSize, pageSize)
	this.Data["json"] = messages
	this.ServeJson()
}
