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
//| Cellgo Framework core type
//| All rights reserved: By cellgo.cn CopyRight
//| You are free to use the source code, but in the use of the process,
//| please keep the author information. Respect for the work of others
//| is respect for their own
//|-------------------------------------------------------------------
// Author:Tommy.Jin Dtime:2016-7-15

package cellgo

import (
	"net/http"
	"reflect"
	"strings"
)

type controllerInfo struct {
	controllerTitle string
	controllerType  reflect.Type
}

type ControllerRegister struct {
	controllers []*controllerInfo
	Coredrive   *Core
}

func NewControllerRegister() *ControllerRegister {
	cr := &ControllerRegister{}
	return cr
}

func (p *ControllerRegister) Add(title string, c ControllerInterface) {
	info := &controllerInfo{}
	t := reflect.Indirect(reflect.ValueOf(c)).Type()
	info.controllerType = t
	info.controllerTitle = title
	p.controllers = append(p.controllers, info)
}

func (p *ControllerRegister) workHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	//M := r.Form["m"]
	if c, a := strings.Join(r.Form["c"], ""), strings.Join(r.Form["a"], ""); c != "" && a != "" {
		var getTtile string
		var getType reflect.Type

		//find a matching controller/action
		for _, cs := range p.controllers {
			if cs.controllerTitle == c {
				getTtile = cs.controllerTitle
				getType = cs.controllerType
				break
			}
		}
		if getTtile != "" {
			//Invoke the request handler
			vc := reflect.New(getType)
			init := vc.MethodByName("Init")
			in := make([]reflect.Value, 4)
			ct := &Netinfo{Response: w, Request: r}
			in[0] = reflect.ValueOf(ct)
			in[1] = reflect.ValueOf(getTtile)
			in[2] = reflect.ValueOf(a)
			in[3] = reflect.ValueOf(p.Coredrive)
			init.Call(in)
			in = make([]reflect.Value, 0)
			method := vc.MethodByName("before")
			method.Call(in)
			/*if r.Method == "GET" {
				method = vc.MethodByName("Get")
				method.Call(in)
			} else if r.Method == "POST" {
				method = vc.MethodByName("Post")
				method.Call(in)
			} else if r.Method == "HEAD" {
				method = vc.MethodByName("Head")
				method.Call(in)
			} else if r.Method == "DELETE" {
				method = vc.MethodByName("Delete")
				method.Call(in)
			} else if r.Method == "PUT" {
				method = vc.MethodByName("Put")
				method.Call(in)
			} else if r.Method == "PATCH" {
				method = vc.MethodByName("Patch")
				method.Call(in)
			} else if r.Method == "OPTIONS" {
				method = vc.MethodByName("Options")
				method.Call(in)
			}*/
			method = vc.MethodByName(a)
			method.Call(in)

			method = vc.MethodByName("After")
			method.Call(in)
		}
	}
	http.NotFound(w, r)
}
