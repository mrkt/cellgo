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
)

var (
	BindExchange map[int]*TcpBind = make(map[int]*TcpBind)
	BindQueue    map[int]*TcpBind = make(map[int]*TcpBind)
)

func init() {
	BindExchange[SOCKETIO] = &TcpBind{TcpType: SOCKETIO, BindMaps: make(map[string]*bindInfo, 10)}
	BindQueue[SOCKETIO] = &TcpBind{TcpType: SOCKETIO, BindMaps: make(map[string]*bindInfo, 10)}
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
			"NewExchange": tb.NewExchange,
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

func (tb *TcpBind) NewExchange(code string, value interface{}) (interface{}, error) {
	res, err := cellgo.Events[tb.BindMaps[code].eventName].EventRead(tb.BindMaps[code].controllerName, tb.BindMaps[code].funcName)
	if err != nil {
		return nil, err
	}
	return res, nil
}
