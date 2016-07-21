package controllers

import (
	"library/service"

	"github.com/mrkt/cellgo"
)

type UserController struct {
	cellgo.Controller
}

func (this *UserController) Before() {
	//init service
	this.GetService(&service.UserService{})
}

func (this *UserController) Run() {
	//param1 funcName, param2 funcParam ...
	userInfo := this.GetServiceFunc("GetUserInfo", "tommy").(map[string]string)
	this.Data["Username"] = userInfo["Username"]
	this.Data["Email"] = userInfo["Email"]
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
