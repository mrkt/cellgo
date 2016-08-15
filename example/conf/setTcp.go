package conf

import (
	"github.com/mrkt/cellgo/tcpip"
)

func SetTcp() {
	tcpip.RegisterTcp(1, ":5000", "/socket.io/", "cellio", "{\"Auth\":\"auth\",\"Push\":\"push\",\"Pull\":\"pull\"}")

}

func BindTcp() {
	tcpip.Bind[1].RegisterHandlers(0, "event2", "EventCreate", "Run")
}

//RunTcp Tcp
func RunSocketIO() {
	tcpip.RunSocketIO()
}
