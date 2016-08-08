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
	CreateExchange(name string, number string) (bool, error)             //Create a Exchange
	RenewExchange(style int, value map[string]interface{}) (bool, error) //Renew Exchange data
	DestroyExchange(number string) (bool, error)                         //Destroy a Exchange
	IncreaseQueue(tcpQueue *TcpQueue, carryInfo string) (bool, error)    //Allow an Queue to enter the Exchange
}

//Queue interface
type TcpQueue interface {
	RegQueue(interface{}) (interface{}, error) //Register a Queue
	IncreasePushed(string) error               //Increase Queue record
	DetectPushed(string) (bool, error)         //Detecting Queue record
}
