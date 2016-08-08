package conf

import (
	"github.com/mrkt/cellgo"
)

func SetEvent() {
	cellgo.RegisterEvent("event1", 1)
	cellgo.RegisterEvent("event2", 1)
}
