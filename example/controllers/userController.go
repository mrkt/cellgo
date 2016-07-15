package controllers

import (
	"github.com/mrkt/cellgo"
)

type UserController struct {
	cellgo.Controller
}

func (this *UserController) show() {
	this.Data["Username"] = "tommy.jin"
	this.Data["Email"] = "tommy.jin@aliyun.com"
	this.TplName = "index.html"
}
