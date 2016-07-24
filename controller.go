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
	Net  *Netinfo
	Data map[interface{}]interface{}

	// template data
	TplDir  string
	TplName string
	TplExt  string

	// controller info
	controllerName string
	actionName     string
	AppController  interface{}

	//service & dao
	serviceType reflect.Type
	daoType     reflect.Type
}

// ControllerInterface is an interface to uniform all controller handler.
type ControllerInterface interface {
	Init(net *Netinfo, controllerName, actionName string, app interface{})
	Before()
	After()
	GetService(ServiceInterface)
	GetDao(DaoInterface)
	GetServiceFunc(string, ...interface{}) interface{}
	GetDaoFunc(string, ...interface{}) interface{}
	Display() error
}

// Init generates default values of controller operations.
func (c *Controller) Init(net *Netinfo, controllerName, actionName string, app interface{}) {
	c.TplName = ""
	c.controllerName = controllerName
	c.actionName = actionName
	c.Net = net
	c.TplExt = CellConf.SiteConfig.TemplateExt
	c.TplDir = CellConf.SiteConfig.TemplatePath
	c.AppController = app
	c.Data = net.Input.Data()
}

// Prepare runs after Init before request function execution.
func (c *Controller) Before() {}

// Finish runs after request function execution.
func (c *Controller) After() {}

//Service type init.
func (c *Controller) GetService(service ServiceInterface) {
	t := reflect.Indirect(reflect.ValueOf(service)).Type()
	c.serviceType = t
}

//Dao type init.
func (c *Controller) GetDao(dao DaoInterface) {
	t := reflect.Indirect(reflect.ValueOf(dao)).Type()
	c.daoType = t
}

//Service function execution.
func (c *Controller) GetServiceFunc(f string, param ...interface{}) interface{} {
	vs := reflect.New(c.serviceType)
	//init before
	in := make([]reflect.Value, 0)
	before := vs.MethodByName("Before")
	before.Call(in)
	//init method
	in = make([]reflect.Value, len(param))
	for k, v := range param {
		in[k] = reflect.ValueOf(v)
	}
	method := vs.MethodByName(f)
	res := method.Call(in)

	//init after
	in = make([]reflect.Value, 0)
	after := vs.MethodByName("After")
	after.Call(in)
	return res[0].Interface()
}

//Dao function execution.
func (c *Controller) GetDaoFunc(f string, param ...interface{}) interface{} {
	vd := reflect.New(c.daoType)
	//init before
	in := make([]reflect.Value, 0)
	before := vd.MethodByName("Before")
	before.Call(in)

	//init method
	in = make([]reflect.Value, len(param))
	for k, v := range param {
		in[k] = reflect.ValueOf(v)
	}
	method := vd.MethodByName(f)
	res := method.Call(in)

	//init after
	in = make([]reflect.Value, 0)
	after := vd.MethodByName("After")
	after.Call(in)
	return res[0].Interface()
}

//Display templates
func (c *Controller) Display() error {
	if c.TplName == "" {
		c.TplName = c.Net.Request.Method + "." + c.TplExt
	}
	t, err := template.ParseFiles(path.Join(c.TplDir, c.TplName))
	if err != nil {
		log.Println("template ParseFiles err:", err)
		return err
	}

	err = t.Execute(c.Net.Response, c.Data)
	if err != nil {
		log.Println("template Execute err:", err)
		return err
	}

	return nil
}
