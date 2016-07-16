package controllers

import (
	"github.com/mrkt/cellgo"
)

type UserController struct {
	cellgo.Controller
}

func (this *UserController) Run() {
	this.Data["Username"] = "tommy.jin"
	this.Data["Email"] = "tommy.jin@aliyun.com"
	this.TplName = "index.html"
}

func (this *UserController) Add() {
	this.Data["Username"] = "tommy.jin"
	this.Data["Email"] = "tommy.jin@aliyun.com"
	this.TplName = "index.html"
}
