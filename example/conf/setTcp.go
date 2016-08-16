package conf

import (
	"github.com/mrkt/cellgo/tcpip"
)

func SetTcp() {
	tcpip.RegisterTcp(1, ":5000", "/socket.io/", "cellio", "{\"Auth\":\"auth\",\"Push\":\"push\",\"Pull\":\"pull\"}")

}

func BindTcp() {
	tcpip.Bind[1].RegisterHandlers(0, "event2", "EventStatic", "Run")
	tcpip.Bind[1].RegisterHandlers(1, "event2", "EventStatic", "Reg")
	tcpip.Bind[1].RegisterHandlers(2, "event2", "EventStatic", "Check")
	tcpip.Bind[1].RegisterHandlers(4, "event2", "EventStatic", "Pull")
}

//RunTcp Tcp
func RunSocketIO() {
	tcpip.RunSocketIO()
}
