package main

import (
	"conf"

	"github.com/mrkt/cellgo"
)

func main() {
	conf.SetController()
	conf.SetEvent()
	conf.SetTcp()
	conf.BindTcp()
	conf.RunSocketIO()
	cellgo.Run()
}
