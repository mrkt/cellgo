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
	"github.com/mrkt/cellgo/interfaces"
)

type SocketIO struct {
	b interfaces.B_SocketIO
}

func (s *SocketIO) SetB(b interfaces.B_SocketIO) {
	s.b = b
}

func (s *SocketIO) RunSocketIO() {

}

type Tcpinter struct {
	b interfaces.B_Tcp
}

func (t *Tcpinter) SetB(b interfaces.B_Tcp) {
	t.b = b
}

func (t *Tcpinter) RegisterTcp(tcpType int, addr string, route string, tcpName string, tcpConf string) error {
	return nil
}
