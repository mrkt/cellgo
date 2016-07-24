package controllers

import (
	"github.com/mrkt/cellgo"
)

type IndexController struct {
	cellgo.Controller
}

func (this *IndexController) Before() {
	this.Data["Username"] = "Hello my friend"
}

func (this *IndexController) Run() {
	this.Data["URI"] = "http://www.cellgo.cn"
	this.TplName = "index.html"
}
