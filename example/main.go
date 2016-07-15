package main

import (
	"conf"
	"fmt"

	"github.com/mrkt/cellgo"
)

func main() {
	conf.SetController()
	cellgo.Run()
	fmt.Println(cellgo.CellCore.Handlers)
}
