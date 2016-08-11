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

package interfaces

type A_SocketIO interface {
	RunSocketIO()
}

type B_SocketIO interface {
	RunSocketIO()
}

type A_Tcp interface {
	RegisterTcp(tcpType int, addr string, route string, tcpName string, tcpConf string) error
}

type B_Tcp interface {
	RegisterTcp(tcpType int, addr string, route string, tcpName string, tcpConf string) error
}

// Register TCP from TCP package
//func RegisterTcp(tcpType int, addr string, route string, tcpName string, tcpConf string) {
//	tcpip.RegisterTcp(tcpType, addr, route, tcpName, tcpConf)
//}

// Run SocketIO
//func RunSocketIO() {
//	tcpip.RunSocketIO()
//}
