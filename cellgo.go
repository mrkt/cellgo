//|------------------------------------------------------------------
//|           __
//|        __/  \
//|     __/  \__/_
//|  __/  \__/ /  \
//| /  \__/  go\__/_
//| \__/_cell  __/  \
//|   /  \  __/  \__/
//|   \__/_/  \__/
//|     /  \__/
//|     \__/
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

// Version number of the cellgo.
const VERSION = "0.0.1"

// Run cellgo framework.
func Run() {
	fmt.Println("Cellgo Version Runing:", VERSION)
	initRawData()
	fmt.Println("Cellgo RawData Runing...")
	CellCore.Run()
}

// Raw data loading
func initRawData() {
	//init Data

}
