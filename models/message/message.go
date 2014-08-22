package message

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	// "wechat/g"
	// "time"
	. "wechat/models"
)

func GetById(id int64) *TextMessageTemplate {
	if id <= 0 {
		return nil
	}
	o := TextMessageTemplate{Id: id}
	err := orm.NewOrm().Read(&o)
	if err != nil {
		return nil
	}
	return &o
}

func Update(t *TextMessageTemplate, text string) error {
	if t.Id == 0 {
		return fmt.Errorf("primary key:id not set")
	}

	if text != "" && t.Text != text {
		t.Text = text
		_, e := orm.NewOrm().Update(t)
		if e != nil {
			return e
		}
		// t.UpdatedAt = time.Now().Unix()
	}
	return nil
}

func Save(message string) (int64, error) {
	o := orm.NewOrm()
	msg := &TextMessageTemplate{
		Text: message,
	}
	return o.Insert(msg)
}

func List(offset, limit int) []*TextMessageTemplate {
	var messages []*TextMessageTemplate
	Texts().Offset(offset).Limit(limit).All(&messages)
	return messages
}

func Texts() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(TextMessageTemplate))
}
