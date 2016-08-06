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
//| Cellgo Framework socketio/socket  file
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
	TcpType int //SOCKET/SOCKETIO/WEBSOCKET/ICMP
	Addr    string
	Route   string
	Handle  interface{} //conn handle
}

func RegisterTcp(tcpType int, addr string, route string, tcpName string) error {
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
		Tcp[SOCKETIO] = append(Tcp[SOCKETIO], &TcpRun{TcpName: tcpName, TcpType: tcpType, Addr: addr, Route: route, Handle: handle})
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

//socket room interface
type Room interface {
	CreateRoom(name string, number string) (bool, error)             //Create a room
	RenewRoom(style int, value map[string]interface{}) (bool, error) //Renew room data
	DestroyRoom(number string) (bool, error)                         //Destroy a room
	IncreaseRunner(runner *Runner, carryInfo string) (bool, error)   //Allow an runner to enter the room
}

type Runner interface {
	RegRunner(interface{}) (interface{}, error) //Register a Runner
	IncreasePushed(string) error                //Increase push record
	DetectPushed(string) (bool, error)          //Detecting push record
}
