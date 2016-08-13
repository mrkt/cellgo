package controllers

import (
	"fmt"
	"library/service"

	"github.com/mrkt/cellgo"
)

type EventCreateController struct {
	cellgo.Controller
}

func (this *EventCreateController) Before() {
	//init service
	this.GetService(&service.UserService{})
}

func (this *EventCreateController) Begin() {
}

func (this *EventCreateController) End() {
}

func (this *EventCreateController) Run() interface{} {
	res := make(map[string]string)
	res["1"] = "青铜组"
	res["2"] = "白银组"
	fmt.Println(res)
	return res
}
