package conf

import (
	"github.com/mrkt/cellgo"
)

func SetTcp() {
	cellgo.RegisterTcp(1, ":5000", "/socket.io/", "cellio", "{\"Auth\":\"auth\",\"Push\":\"push\",\"Pull\":\"pull\"}")

}
