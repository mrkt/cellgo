//|------------------------------------------------------------------
//|           __
//|        __/  \
//|     __/  \__/_
//|  __/  \__/ /  \
//| /  \__/  go\__/_
//| \__/_cell  __/  \
//|   /  \  __/  \__/
//|   \__/_/  \__/
//|     /  \__/
//|     \__/
//|------------------------------------------------------------------
//| Cellgo Framework controller type
//| All rights reserved: By cellgo.cn CopyRight
//| You are free to use the source code, but in the use of the process,
//| please keep the author information. Respect for the work of others
//| is respect for their own
//|-------------------------------------------------------------------
// Author:Tommy.Jin Dtime:2016-7-15

package cellgo

type Controller struct {
	// NetInfo data
	ni   *Netinfo
	Data map[interface{}]interface{}

	// template data
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
	GetService()
	GetDao()
}

// Init generates default values of controller operations.
func (c *Controller) Init(ni *Netinfo, controllerName, actionName string, app interface{}) {
	c.TplName = ""
	c.controllerName = controllerName
	c.actionName = actionName
	c.ni = ni
	c.TplExt = "html"
	c.AppController = app
	//c.Data = ni.Input.Data()
}

// Prepare runs after Init before request function execution.
func (c *Controller) Before() {}

// Finish runs after request function execution.
func (c *Controller) After() {}

// Prepare runs after Init before request function execution.
func (c *Controller) GetService() {}

// Finish runs after request function execution.
func (c *Controller) GetDao() {}
