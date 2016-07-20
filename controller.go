//|------------------------------------------------------------------
//|        __
//|     __/  \
//|  __/  \__/_
//| /  \__/    \
//|/\__/CellGo /_
//|\/_/NetFW__/  \
//|  /\__ _/  \__/
//|  \/_/  \__/_/
//|    /\__/_/
//|    \/_/
//|------------------------------------------------------------------
//| Cellgo Framework controller type
//| All rights reserved: By cellgo.cn CopyRight
//| You are free to use the source code, but in the use of the process,
//| please keep the author information. Respect for the work of others
//| is respect for their own
//|-------------------------------------------------------------------
// Author:Tommy.Jin Dtime:2016-7-15

package cellgo

import (
	"html/template"
	"log"
	"path"
	"reflect"
)

type Controller struct {
	// NetInfo data
	Ni   *Netinfo
	Data map[interface{}]interface{}

	// template data
	TplDir  string
	TplName string
	TplExt  string

	// controller info
	controllerName string
	actionName     string
	AppController  interface{}
}

// ControllerInterface is an interface to uniform all controller handler.
type ControllerInterface interface {
	Init(ni *Netinfo, controllerName, actionName string, app interface{})
	Before()
	After()
	GetService(ServiceInterface)
	GetDao(DaoInterface)
	Display() error
}

// Init generates default values of controller operations.
func (c *Controller) Init(ni *Netinfo, controllerName, actionName string, app interface{}) {
	c.TplName = ""
	c.controllerName = controllerName
	c.actionName = actionName
	c.Ni = ni
	c.TplExt = CellConf.SiteConfig.TemplateExt
	c.TplDir = CellConf.SiteConfig.TemplatePath
	c.AppController = app
	c.Data = ni.Input.Data()
}

// Prepare runs after Init before request function execution.
func (c *Controller) Before() {}

// Finish runs after request function execution.
func (c *Controller) After() {}

//Service function execution.
func (c *Controller) GetService(service ServiceInterface) {
	getType := reflect.Indirect(reflect.ValueOf(service)).Type()
	vs := reflect.New(getType)
	in := make([]reflect.Value, 0)
	method := vs.MethodByName("Test")
	method.Call(in)

}

//Dao function execution.
func (c *Controller) GetDao(dao DaoInterface) {
	getType := reflect.Indirect(reflect.ValueOf(dao)).Type()
	vs := reflect.New(getType)
	in := make([]reflect.Value, 0)
	method := vs.MethodByName("Test")
	method.Call(in)
}

func (c *Controller) Display() error {
	if c.TplName == "" {
		c.TplName = c.Ni.Request.Method + "." + c.TplExt
	}
	t, err := template.ParseFiles(path.Join(c.TplDir, c.TplName))
	if err != nil {
		log.Println("template ParseFiles err:", err)
		return err
	}

	err = t.Execute(c.Ni.Response, c.Data)
	if err != nil {
		log.Println("template Execute err:", err)
		return err
	}

	return nil
}
