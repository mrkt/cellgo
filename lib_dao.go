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
//| Cellgo Framework dao type
//| All rights reserved: By cellgo.cn CopyRight
//| You are free to use the source code, but in the use of the process,
//| please keep the author information. Respect for the work of others
//| is respect for their own
//|-------------------------------------------------------------------
// Author:Tommy.Jin Dtime:2016-7-20

package cellgo

import (
	"reflect"
)

type Dao struct {
}

// DaoInterface is an interface to uniform all Dao handler.
type DaoInterface interface {
	Init()
	GetService(ServiceInterface)
	GetDao(DaoInterface)
}

// Init generates default values of controller operations.
func (c *Dao) Init() {}

//Service function execution.
func (c *Dao) GetService(service ServiceInterface) {
	getType := reflect.Indirect(reflect.ValueOf(service)).Type()
	vs := reflect.New(getType)
	in := make([]reflect.Value, 0)
	method := vs.MethodByName("Test")
	method.Call(in)

}

//Dao function execution.
func (c *Dao) GetDao(dao DaoInterface) {
	getType := reflect.Indirect(reflect.ValueOf(dao)).Type()
	vs := reflect.New(getType)
	in := make([]reflect.Value, 0)
	method := vs.MethodByName("Test")
	method.Call(in)
}
