package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/chanxuehong/wechat/message/request"
	"github.com/chanxuehong/wechat/message/response"
	"github.com/chanxuehong/wechat/server"
	"net/http"
	"wechat/g"
	_ "wechat/models"
	"wechat/models/message"
	_ "wechat/routers"
)

const (
	TOKEN = "your_weixin_token"
)

// 一般一个应用维护一个实例即可
var wechatHandler *server.Handler

func main() {
	g.InitEnv()
	// 初始化 wechatHandler
	setting := &server.HandlerSetting{
		Token:              TOKEN,
		TextRequestHandler: TextRequestHandler,
	}
	wechatHandler = server.NewHandler(setting) // 并发安全，一般一个应用只用一个实例即可

	beego.Handler("/wechat", wechatHandler)
	beego.Run()
	orm.RunCommand()
	orm.Debug = true
}

// 自定义文本消息处理函数
func TextRequestHandler(w http.ResponseWriter, r *http.Request, text *request.Text) {
	if msg := message.GetByText(text.Content); msg != nil {
		// 把用户发送过来的文本原样的回复过去
		resp := response.NewText(text.FromUserName, text.ToUserName, msg.Text)

		w.Header().Set("Content-Type", "application/xml; charset=utf-8") // 可选
		server.WriteText(w, resp)
	}
}
