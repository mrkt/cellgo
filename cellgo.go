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
//| Cellgo Framework Run file
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

var cellShow = string(`
//|------------------------------------------------------------------
//|         __          
//|      __/  \         
//|   __/  \__/_       ♥ ♥ 
//|  /  \__/    \       Using Golang!
//| /\__/CellGo /_      -------------  
//| \/_/NetFW__/  \                   
//|   /\__ _/  \__/                   
//|   \/_/  \__/_/                      
//|     /\__/_/                         
//|     \/_/                                
//|------------------------------------------------------------------
//| Cellgo framework core file has been started
//| Feel very much your support for the cause of open source
//| Have any questions please send mail to <tommy.jin@aliyun.com>
//| Or enter the GitHub submission issue <github.com/mrkt>
//| Hope you have a nice day !
//|-------------------------------------------------------------------
`)

// Run cellgo framework.
func Run() {
	fmt.Println(cellShow)
	fmt.Println(CellConf.ServerName, " is Runing...")
	powerBoot()
	fmt.Println("Cellgo RawData Runing...")
	CellCore.Run()
}

// Raw data loading
func powerBoot() {
	//init Data whit boot
	bt := Boot{}
	bt.GCEvent()
}
