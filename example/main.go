package main

import (
	"conf"

	"github.com/mrkt/cellgo"
)

func main() {
	conf.SetController()
	conf.SetEvent()
	conf.SetTcp()
	cellgo.Run()
}
