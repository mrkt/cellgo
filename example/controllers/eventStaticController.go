package controllers

import (
	"fmt"
	//	"library/service"

	"github.com/mrkt/cellgo"
)

type EventStaticController struct {
	cellgo.Controller
}

func (this *EventStaticController) Before() {
	//init service
	//this.GetService(&service.UserService{})
}

func (this *EventStaticController) Begin() {
}

func (this *EventStaticController) End() {
}

func (this *EventStaticController) Run() interface{} {
	res := make(map[string]string)
	res["0"] = "test1"
	res["1"] = "test2"
	res["2"] = "test3"
	return res
}

func (this *EventStaticController) Reg(value interface{}) interface{} {
	var res map[string]string
	res = map[string]string{"CarryInfo": value.(string) + "00", "Exchange": "4"}
	return res
}

func (this *EventStaticController) Check(value interface{}) interface{} {
	var res map[string]string
	res = map[string]string{"FromInfo": "3267", "Exchange": "4"}
	return res
}

func (this *EventStaticController) Pull(value interface{}) interface{} {
	val := value.(string)
	fmt.Println(val)
	return "{id:4,price:30}"
}
