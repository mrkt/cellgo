package controllers

import (
	"fmt"
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
	this.Data["URI"] = this.Net.Input.Site() + this.Net.Input.URI()
	this.TplName = "index.html"
}

func (this *UserController) Add() {
	username := this.Net.Input.GetGP("username", true)
	email := this.Net.Input.GetGP("email", true)
	if user := this.Net.Input.Session.Get("user"); user != nil {
		fmt.Println(user.(map[string]string))
	} else {
		var user map[string]string = make(map[string]string)
		user["username"] = "tommy"
		user["email"] = "tommy.jin@aliyun.com"
		this.Net.Input.Session.Set("user", user)
	}
	this.Data["Username"] = username
	this.Data["Email"] = email
	this.Data["URI"] = this.Net.Input.Site() + this.Net.Input.URI()
	this.TplName = "index.html"
}
