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
//| Cellgo Framework tcpip/bind file
//| All rights reserved: By cellgo.cn CopyRight
//| You are free to use the source code, but in the use of the process,
//| please keep the author information. Respect for the work of others
//| is respect for their own
//|-------------------------------------------------------------------
// Author:Tommy.Jin Dtime:2016-08-08

package tcpip

import (
	"errors"

	"github.com/mrkt/cellgo"
)

//const bindType
const (
	NEWEXCHANGE = iota
	REGQUEUE
	CHECKQUEUE
	PUSH
	PULL
)

var (
	Bind map[int]*TcpBind = make(map[int]*TcpBind)
)

func init() {
	Bind[SOCKETIO] = &TcpBind{TcpType: SOCKETIO, BindMaps: make(map[string]*bindInfo, 10)}
}

// TcpBind type.
type TcpBind struct {
	TcpType  int
	BindMaps map[string]*bindInfo
}

// TcpBind Handler type
type bindInfo struct {
	handler        func(string, interface{}) (interface{}, error)
	bindCode       string
	bindType       int
	eventName      string
	controllerName string
	funcName       string
}

// register Command and handle function
func (tb *TcpBind) RegisterHandlers(bindType int, eventName string, controllerName string, funcName string) {
	var m map[string]func(string, interface{}) (interface{}, error)
	switch bindType {
	case NEWEXCHANGE:
		m = map[string]func(string, interface{}) (interface{}, error){
			"New": tb.Happen,
		}
		break
	case REGQUEUE:
		m = map[string]func(string, interface{}) (interface{}, error){
			"Reg": tb.Happen,
		}
		break
	case CHECKQUEUE:
		m = map[string]func(string, interface{}) (interface{}, error){
			"Check": tb.Happen,
		}
		break
	case PUSH:
		m = map[string]func(string, interface{}) (interface{}, error){
			"Push": tb.BatchHappens,
		}
		break
	case PULL:
		m = map[string]func(string, interface{}) (interface{}, error){
			"Pull": tb.Happen,
		}
		break
	default:
		break
	}
	for e, h := range m {
		if _, ok := tb.BindMaps[e]; !ok {
			tb.ExchangeHandler(e, h, bindType, eventName, controllerName, funcName)
		}
	}
}

func (tb *TcpBind) ExchangeHandler(code string, h func(string, interface{}) (interface{}, error), Type int, eName string, cName string, fName string) {
	tb.BindMaps[code] = &bindInfo{
		handler:        h,
		bindCode:       code,
		bindType:       Type,
		eventName:      eName,
		controllerName: cName,
		funcName:       fName,
	}
}

//Perform binding event only Happen
func (tb *TcpBind) Happen(code string, value interface{}) (interface{}, error) {
	res, err := cellgo.Events[tb.BindMaps[code].eventName].EventRead(tb.BindMaps[code].controllerName, tb.BindMaps[code].funcName, value)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//Perform binding event batch Happens
func (tb *TcpBind) BatchHappens(code string, value interface{}) (interface{}, error) {
	happens, err := tb.findHappen(code)
	if err != nil {
		return nil, err
	}
	var res map[string]interface{} = make(map[string]interface{})
	for _, v := range happens {
		hs, err := tb.happens(code, v, value)
		if err != nil {
			return nil, err
		}
		res[v] = hs
	}
	return res, nil
}

//Perform binding event batch Happens's one
func (tb *TcpBind) happens(code string, title string, value interface{}) (interface{}, error) {
	res, err := cellgo.Events[tb.BindMaps[code].eventName].EventRead(title, tb.BindMaps[code].funcName, value)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//Find binding event batch Happens's title
func (tb *TcpBind) findHappen(code string) ([]string, error) {
	var res []string
	for k, _ := range cellgo.Events[tb.BindMaps[code].eventName].Happened {
		res = append(res, k)
	}
	if res == nil {
		return nil, errors.New("The Happen is not found.")
	}
	return res, nil
}
