package controllers

import (
	"library/service"

	"github.com/mrkt/cellgo"
)

type UserController struct {
	cellgo.Controller
}

func (this *UserController) Before() {
	//test service & dao
	this.GetService(&service.UserService{})
}

func (this *UserController) Run() {
	this.Data["Username"] = "tommy.jin"
	this.Data["Email"] = "tommy.jin@aliyun.com"
	this.Data["URI"] = this.Ni.Input.Site() + this.Ni.Input.URI()
	this.TplName = "index.html"
}

func (this *UserController) Add() {
	username := this.Ni.Input.Param("username")
	email := this.Ni.Input.Param("email")
	this.Data["Username"] = username
	this.Data["Email"] = email
	this.Data["URI"] = this.Ni.Input.Site() + this.Ni.Input.URI()
	this.TplName = "index.html"
}
