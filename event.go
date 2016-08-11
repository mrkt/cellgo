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
//| ------------------------------------------------------------------
//| Cellgo Framework envent type
//| All rights reserved: By cellgo.cn CopyRight
//| You are free to use the source code, but in the use of the process,
//| please keep the author information. Respect for the work of others
//| is respect for their own
//|-------------------------------------------------------------------
// Author:Tommy.Jin Dtime:2016-08-02

package cellgo

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"reflect"
	"sync"
	"time"

	ctcpip "github.com/mrkt/cellgo/tcpip"
	"github.com/mrkt/tcpip"
)

var Events = make(map[string]*Event)

//happen type
type happen struct {
	controllerTitle string
	controllerType  reflect.Type
	param           []string
	begin           int64
	end             int64
	coreData        interface{}
}

/*func (h *happen) Set(key, value interface{}) {
}

func (h *happen) Get(key interface{}) {
}

func (h *happen) Delete(key interface{}) {
}*/

//Event type
type Event struct {
	lock     sync.Mutex
	Happen   map[string]*happen
	Happened map[string]*happen
	waitTime int64
	EventId  string
}

//Initialization event, and initial happen, set waitTime
func (e *Event) EventInit(waitTime int64) (bool, error) {
	e.lock.Lock()
	defer e.lock.Unlock()
	e.Happen = make(map[string]*happen)
	e.Happened = make(map[string]*happen)
	e.waitTime = waitTime
	id, err := e.OnlyId()
	if err != nil {
		return false, err
	}
	e.EventId = id
	return true, nil
}

//Generate only ID
func (e *Event) OnlyId() (string, error) {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

//Call the begin happen's controller and execute its function
func (e *Event) EventRead(title string, act string) (interface{}, error) {

	var getTitle string
	var getType reflect.Type
	var getParam string
	var getCore interface{}
	for _, v := range e.Happened {
		if v.controllerTitle == title {
			for _, pr := range v.param {
				if act == pr {
					getParam = pr
				}
			}
			if getParam == "" {
				return nil, errors.New("Parame that are not defined.")
			}
			getTitle = v.controllerTitle
			getType = v.controllerType
			getCore = v.coreData
			break
		}
	}
	if getTitle != "" {
		return nil, errors.New("Controller that are not defined.")
	}
	vc := reflect.New(getType)
	init := vc.MethodByName("Init")
	in := make([]reflect.Value, 4)
	//Assignment parameter
	in[0] = reflect.ValueOf(nil)
	in[1] = reflect.ValueOf(getTitle)
	in[2] = reflect.ValueOf(getParam)
	in[3] = reflect.ValueOf(getCore)
	init.Call(in)

	in = make([]reflect.Value, 0)
	method := vc.MethodByName("Before")
	method.Call(in)

	method = vc.MethodByName(getParam)
	resEvent := method.Call(in)

	method = vc.MethodByName("After")
	method.Call(in)

	return resEvent, nil
}

//Destroy Event's happen & happened
func (e *Event) EventDestroy(title string, hp bool) {
	if hp {
		if e.Happened[title] != nil {
			delete(e.Happened, title)
		}
	} else {
		if e.Happen[title] != nil {
			delete(e.Happen, title)
		}
	}
}

//If Happe begin time is up to GC
func (e *Event) EventON() {
	for {
		if e.Happen != nil {
			for k, v := range e.Happen {
				if v.begin < time.Now().Unix() {
					e.Happened[k] = v
					vc := reflect.New(v.controllerType)
					init := vc.MethodByName("Init")
					in := make([]reflect.Value, 4)
					//Assignment parameter with begin
					in[0] = reflect.ValueOf(nil)
					in[1] = reflect.ValueOf(v.controllerTitle)
					in[2] = reflect.ValueOf("Begin") //event on function
					in[3] = reflect.ValueOf(v.coreData)
					init.Call(in)
					in = make([]reflect.Value, 0)
					method := vc.MethodByName("Begin ")
					method.Call(in)
					delete(e.Happen, k)
				}
			}
		}
		time.Sleep(time.Second * 1) //stop 1 sec check
		//fmt.Println("ON")
	}
}

//If Happend failure time is up to GC
func (e *Event) EventGC() {
	for {
		if e.Happened != nil {
			for k, v := range e.Happened {
				if v.end < time.Now().Unix() {
					vc := reflect.New(v.controllerType)
					init := vc.MethodByName("Init")
					in := make([]reflect.Value, 4)
					//Assignment parameter with begin
					in[0] = reflect.ValueOf(nil)
					in[1] = reflect.ValueOf(v.controllerTitle)
					in[2] = reflect.ValueOf("End") //event on function
					in[3] = reflect.ValueOf(v.coreData)
					init.Call(in)

					in = make([]reflect.Value, 0)
					method := vc.MethodByName("End")
					method.Call(in)
					delete(e.Happened, k)
				}
			}
		}
		time.Sleep(time.Second * 1) //stop 1 sec check
		//fmt.Println("GC")
	}
}

//An instance of happen is added to the event instance
func (e *Event) EventAdd(title string, c ControllerInterface, fc []string, begin int64, end int64, coreDate interface{}) {
	info := &happen{}
	t := reflect.Indirect(reflect.ValueOf(c)).Type()
	info.controllerType = t
	info.controllerTitle = title
	info.param = fc
	info.begin = begin
	info.end = end
	info.coreData = coreDate
	e.Happen[title] = info
}

// RegisterEvent makes a Event by the Event name.
func RegisterEvent(name string, waitTime int64) {
	if Events[name] != nil {
		panic("Event: Register Events not nil")
	}
	event := &Event{}
	event.EventInit(waitTime)
	Events[name] = event
}

// Register TCP from TCP package
func RegisterTcp(tcpType int, addr string, route string, tcpName string, tcpConf string) {
	a_tcp := new(ctcpip.Tcpinter)
	b_tcp := new(tcpip.Tcpinter)
	a_tcp.SetB(b_tcp)
	a_tcp.RegisterTcp(tcpType, addr, route, tcpName, tcpConf)
}
