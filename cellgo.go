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
//|------------------------------------------------------------------
//| Cellgo Framework Boot file
//| All rights reserved: By cellgo.cn CopyRight
//| You are free to use the source code, but in the use of the process,
//| please keep the author information. Respect for the work of others
//| is respect for their own
//|-------------------------------------------------------------------
// Author:Tommy.Jin Dtime:2016-7-15

package cellgo

import (
	"fmt"
)

// Run cellgo framework.
func Run() {
	fmt.Println(CellConf.ServerName, " is Runing...")
	initRawData()
	fmt.Println("Cellgo RawData Runing...")
	CellCore.Run()
}

// Raw data loading
func initRawData() {
	//init Data

}
