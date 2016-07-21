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
//| Cellgo Framework service type
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

type Service struct {
	serviceType reflect.Type
	daoType     reflect.Type
}

// ServiceInterface is an interface to uniform all service handler.
type ServiceInterface interface {
	Init()
	Before()
	After()
	GetService(ServiceInterface)
	GetDao(DaoInterface)
	GetServiceFunc(string, ...interface{}) interface{}
	GetDaoFunc(string, ...interface{}) interface{}
}

// Init generates default values of controller operations.
func (s *Service) Init() {
}

// Prepare runs after Init before request function execution.
func (s *Service) Before() {}

// Finish runs after request function execution.
func (s *Service) After() {}

//Service type init
func (s *Service) GetService(service ServiceInterface) {
	t := reflect.Indirect(reflect.ValueOf(service)).Type()
	s.serviceType = t
}

//Dao type init
func (s *Service) GetDao(dao DaoInterface) {
	t := reflect.Indirect(reflect.ValueOf(dao)).Type()
	s.daoType = t
}

//Service function execution.
func (s *Service) GetServiceFunc(f string, param ...interface{}) interface{} {
	vs := reflect.New(s.serviceType)
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
func (s *Service) GetDaoFunc(f string, param ...interface{}) interface{} {
	vd := reflect.New(s.daoType)
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
