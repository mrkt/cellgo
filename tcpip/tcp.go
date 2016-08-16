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
//| Cellgo Framework tcpip/tcp  file
//| All rights reserved: By cellgo.cn CopyRight
//| You are free to use the source code, but in the use of the process,
//| please keep the author information. Respect for the work of others
//| is respect for their own
//|-------------------------------------------------------------------
// Author:Tommy.Jin Dtime:2016-08-06

package tcpip

import (
	"github.com/googollee/go-socket.io"
)

//const TcpType
const (
	SOCKET = iota
	SOCKETIO
	WEBSOCKET
	ICMP
)

var (
	Tcp = make(map[int][]*TcpRun)
)

type TcpRun struct {
	TcpName string
	TcpType int    //SOCKET/SOCKETIO/WEBSOCKET/ICMP
	TcpConf string //json info
	Addr    string
	Route   string
	Handle  interface{} //conn handle
}

func RegisterTcp(tcpType int, addr string, route string, tcpName string, tcpConf string) error {
	var (
		handle interface{}
		err    error
	)
	switch {
	case tcpType == SOCKET:
		break
	case tcpType == SOCKETIO:
		handle, err = socketio.NewServer(nil)
		if err != nil {
			return err
		}
		Tcp[SOCKETIO] = append(Tcp[SOCKETIO], &TcpRun{TcpName: tcpName, TcpType: tcpType, TcpConf: tcpConf, Addr: addr, Route: route, Handle: handle})
		break
	case tcpType == WEBSOCKET:
		break
	case tcpType == ICMP:
		break
	default:
		break

	}
	return nil
}

//Exchange interface
type TcpExchange interface {
	NewExchange(int) (bool, error)              //Create a Exchange
	Push(int, interface{}) (interface{}, error) //Push a Queue
	Pull(int, interface{}) (interface{}, error) //Pull a Queue
}

//Queue interface
type TcpQueue interface {
	RegQueue(int, interface{}) (interface{}, error)   //Register a Queue
	CheckQueue(int, interface{}) (interface{}, error) //Check a Queue
}
