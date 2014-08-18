package controllers

import (
	"crypto/sha1"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"sort"
)

const (
	TOKEN = "bruceinpeking"
)

func CheckErr(field string, err interface{}) {
	if err != nil {
		panic(field + " is null")
	}
}

type Echo struct {
	Signature string
	Timestamp int64
	Nonce     int64
	EchoStr   string
}

type MainController struct {
	beego.Controller
}

func (this *MainController) getEcho() *Echo {
	signature := this.GetString("signature")
	timestamp, err := this.GetInt("timestamp")
	CheckErr("timestamp", err)
	nonce, err := this.GetInt("nonce")
	CheckErr("nonce", err)
	echoStr := this.GetString("echoStr")
	return &Echo{
		signature,
		timestamp,
		nonce,
		echoStr,
	}
}

func checkSignature(echo *Echo) bool {
	tmps := []string{TOKEN, string(echo.Timestamp), string(echo.Nonce)}
	sort.Strings(tmps)
	tmpStr := tmps[0] + tmps[1] + tmps[2]
	tmp := str2sha1(tmpStr)
	return tmp == echo.Signature
}

func str2sha1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

func (this *MainController) Get() {
	echo := this.getEcho()
	ret := checkSignature(echo)
	if ret {
		this.Ctx.WriteString(echo.EchoStr)
	} else {
		this.Ctx.WriteString("Error, echoStr not match!")
	}
}
