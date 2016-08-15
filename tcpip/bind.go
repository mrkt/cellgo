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
			"New": tb.Dispatch,
		}
		break
	case REGQUEUE:
		m = map[string]func(string, interface{}) (interface{}, error){
			"Reg": tb.Dispatch,
		}
		break
	case CHECKQUEUE:
		m = map[string]func(string, interface{}) (interface{}, error){
			"Check": tb.Dispatch,
		}
		break
	case PUSH:
		m = map[string]func(string, interface{}) (interface{}, error){
			"Push": tb.Dispatch,
		}
		break
	case PULL:
		m = map[string]func(string, interface{}) (interface{}, error){
			"Pull": tb.Dispatch,
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

func (tb *TcpBind) Dispatch(code string, value interface{}) (interface{}, error) {
	res, err := cellgo.Events[tb.BindMaps[code].eventName].EventRead(tb.BindMaps[code].controllerName, tb.BindMaps[code].funcName)
	if err != nil {
		return nil, err
	}
	return res, nil
}
