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
//| Cellgo Framework socketio/socketio_runner file
//| All rights reserved: By cellgo.cn CopyRight
//| You are free to use the source code, but in the use of the process,
//| please keep the author information. Respect for the work of others
//| is respect for their own
//|-------------------------------------------------------------------
// Author:Tommy.Jin Dtime:2016-08-06

package socket

import (
	"errors"
)

//Socketio Runner Operation type
type SocketioRunner struct {
	FromInfo  interface{} //socketio runner's enter identification information
	CarryInfo interface{} //socketio runner's carrying identification information
	Pushed    []string    //socketio runner's pushed log
}

func (s *SocketioRunner) RegRunner(interface{}) (interface{}, error) {
	return "", errors.New("")
}

func (s *SocketioRunner) IncreasePushed(string) error {
	return nil
}

func (s *SocketioRunner) DetectPushed(string) (bool, error) {
	return true, nil
}
