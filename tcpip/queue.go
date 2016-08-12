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
//| Cellgo Framework tcpip/queue file
//| All rights reserved: By cellgo.cn CopyRight
//| You are free to use the source code, but in the use of the process,
//| please keep the author information. Respect for the work of others
//| is respect for their own
//|-------------------------------------------------------------------
// Author:Tommy.Jin Dtime:2016-08-06

package tcpip

import (
	"errors"
)

//Queue Operation type
type Queue struct {
	FromInfo  interface{} //Queue's enter identification information
	CarryInfo interface{} //Queue's carrying identification information
	Pushed    []string    //Queue's pushed log
}

func (q *Queue) RegQueue(interface{}) (interface{}, error) {
	return "", errors.New("")
}

func (q *Queue) IncreasePushed(string) error {
	return nil
}

func (q *Queue) DetectPushed(string) (bool, error) {
	return true, nil
}
