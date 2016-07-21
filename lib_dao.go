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
	daoType reflect.Type
}

// DaoInterface is an interface to uniform all Dao handler.
type DaoInterface interface {
	Init()
	Before()
	After()
	GetDao(DaoInterface)
	GetDaoFunc(string, ...interface{}) interface{}
}

// Init generates default values of controller operations.
func (d *Dao) Init() {}

// Prepare runs after Init before request function execution.
func (d *Dao) Before() {}

// Finish runs after request function execution.
func (d *Dao) After() {}

//Dao function execution.
func (d *Dao) GetDao(dao DaoInterface) {
	t := reflect.Indirect(reflect.ValueOf(dao)).Type()
	d.daoType = t
}

//Dao function execution.
func (d *Dao) GetDaoFunc(f string, param ...interface{}) interface{} {
	vd := reflect.New(d.daoType)
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
