package controllers

import (
	"library/service"

	"github.com/mrkt/cellgo"
)

type EventPushController struct {
	cellgo.Controller
}

func (this *EventPushController) Before() {
	//init service
	this.GetService(&service.UserService{})
}

func (this *EventPushController) Run() {

}
