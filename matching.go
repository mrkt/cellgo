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
	"sync"
)

type controllerInfo struct {
	controllerTitle string
	controllerType  reflect.Type
	param           []string
}

type ControllerRegister struct {
	controllers []*controllerInfo
	Coredrive   *Core
	pool        sync.Pool //池
}

func NewControllerRegister() *ControllerRegister {
	cr := &ControllerRegister{}
	cr.pool.New = func() interface{} {
		return NewNetInfo()
	}
	return cr
}

func (p *ControllerRegister) Add(title string, c ControllerInterface, fc []string) {
	info := &controllerInfo{}
	t := reflect.Indirect(reflect.ValueOf(c)).Type()
	info.controllerType = t
	info.controllerTitle = title
	info.param = fc
	p.controllers = append(p.controllers, info)
}

func (p *ControllerRegister) workHTTP(w http.ResponseWriter, r *http.Request) {
	net := p.pool.Get().(*Netinfo)
	net.Reset(w, r)
	defer p.pool.Put(net) //销毁池
	// session init
	if CellConf.SiteConfig.SessionOn {
		var err error
		net.Input.Session, err = SESSION.SessionStart(w, r)
		if err != nil {
			CellError.ErrMaps["500"].handler(w, r)
			return
		}
		/*defer func() {
			if Netinfo.Input.Session != nil {
				Netinfo.Input.Session.SessionRelease(rw)
			}
		}()*/
	}
	//M := r.Form["m"]
	if c, a := strings.Join(r.Form["c"], ""), strings.Join(r.Form["a"], ""); c != "" && a != "" {
		var getTitle string
		var getType reflect.Type
		var getParam string

		//find a matching controller/action
		for _, cs := range p.controllers {
			if cs.controllerTitle == c {
				for _, pr := range cs.param {
					if p.IndexToUpper(a) == pr {
						getParam = pr
					}
				}
				if getParam == "" {
					goto over
				}
				getTitle = cs.controllerTitle
				getType = cs.controllerType
				break
			}
		}
		if getTitle != "" {
			//Invoke the request handler
			vc := reflect.New(getType)
			init := vc.MethodByName("Init")
			in := make([]reflect.Value, 4)
			//Assignment parameter
			for key, value := range r.Form {
				net.Input.SetParam(key, strings.Join(value, ""))
			}
			in[0] = reflect.ValueOf(net)
			in[1] = reflect.ValueOf(getTitle)
			in[2] = reflect.ValueOf(a)
			in[3] = reflect.ValueOf(p.Coredrive)
			init.Call(in)
			in = make([]reflect.Value, 0)
			method := vc.MethodByName("Before")
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

			method = vc.MethodByName(getParam)
			method.Call(in)

			method = vc.MethodByName("After")
			method.Call(in)

			if CellConf.SiteConfig.AutoDisplay {
				method = vc.MethodByName("Display")
				method.Call(in)
			}
			return
		}
	}
over:
	//http.NotFound(w, r)
	CellError.ErrMaps["404"].handler(w, r)
}

func (p *ControllerRegister) IndexToUpper(str string) string {
	strlen := len(str)
	index := strings.ToUpper(string([]byte(str)[0:1]))
	prefix := string([]byte(str)[1:strlen])

	return index + prefix
}
