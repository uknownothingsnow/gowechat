package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type TextMessageTemplate struct {
	Id   int64
	Text string
}

type ImageMessageTemplate struct {
	Id       int64
	ImageUrl string
	Title    string
}

type Rule struct {
	Id         int64
	Text       string `orm:"unique"`
	ReplyType  string
	TemplateId int64     `orm:"index"`
	Created    time.Time `orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.Debug = true
	// 需要在init中注册定义的model
	orm.RegisterModel(new(TextMessageTemplate), new(ImageMessageTemplate), new(Rule))
}
