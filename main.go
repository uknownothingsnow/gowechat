package main

import (
	"github.com/astaxie/beego"
	"github.com/chanxuehong/wechat/message/request"
	"github.com/chanxuehong/wechat/message/response"
	"github.com/chanxuehong/wechat/server"
	"net/http"
)

const (
	TOKEN = "your_weixin_token"
)

// 一般一个应用维护一个实例即可
var wechatHandler *server.Handler

func main() {
	// 初始化 wechatHandler
	setting := &server.HandlerSetting{
		Token:              TOKEN,
		TextRequestHandler: TextRequestHandler,
	}
	wechatHandler = server.NewHandler(setting) // 并发安全，一般一个应用只用一个实例即可

	beego.Handler("/wechat", wechatHandler)
	beego.Run()
}

// 自定义文本消息处理函数
func TextRequestHandler(w http.ResponseWriter, r *http.Request, text *request.Text) {
	//TODO: 添加你的代码，下面只是示例代码！
	// 把用户发送过来的文本原样的回复过去
	resp := response.NewText(text.FromUserName, text.ToUserName, text.Content)

	w.Header().Set("Content-Type", "application/xml; charset=utf-8") // 可选
	server.WriteText(w, resp)
}
