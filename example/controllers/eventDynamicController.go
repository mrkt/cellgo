package controllers

import (
	//	/"library/service"

	"github.com/mrkt/cellgo"
)

type EventDynamicController struct {
	cellgo.Controller
}

func (this *EventDynamicController) Before() {
	//init service
	//this.GetService(&service.UserService{})
}

func (this *EventDynamicController) Run() {

}
