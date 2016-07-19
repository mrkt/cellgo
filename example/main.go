package main

import (
	"conf"

	"github.com/mrkt/cellgo"
)

func main() {
	conf.SetController()
	cellgo.Run()
}
