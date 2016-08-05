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
//| ------------------------------------------------------------------
//| Cellgo Framework socketio/socket file
//| All rights reserved: By cellgo.cn CopyRight
//| You are free to use the source code, but in the use of the process,
//| please keep the author information. Respect for the work of others
//| is respect for their own
//|-------------------------------------------------------------------
// Author:Tommy.Jin Dtime:2016-08-06

package socket

//socket room interface
type Room interface {
	CreateRoom(name string, number string) (bool, error)             //Create a room
	RenewRoom(style int, value map[string]interface{}) (bool, error) //Renew room data
	DestroyRoom(number string) (bool, error)                         //Destroy a room
	IncreaseRunner(runner *Runner, carryInfo string) (bool, error)   //Allow an runner to enter the room
}

type Runner interface {
	RegRunner(interface{}) (interface{}, error) //Register a Runner
	IncreasePushed(string) error                //Increase push record
	DetectPushed(string) (bool, error)          //Detecting push record
}
