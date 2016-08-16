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
	"fmt"
)

var (
	Queues *Queue
)

func init() {
	Queues = &Queue{}
}

//Queue Operation type
type Queue struct {
	FromInfo  interface{} //Queue's enter identification information
	CarryInfo interface{} //Queue's carrying identification information
	Pushed    []string    //Queue's p ushed log
}

func (q *Queue) RegQueue(tcpType int, value interface{}) (interface{}, error) {
	res, err := Bind[tcpType].BindMaps["Reg"].handler("Reg", value)
	if err != nil {
		return false, err
	}
	info := res.(map[string]string)
	fmt.Println(info)
	return "", errors.New("")
}

func (q *Queue) CheckQueue(tcpType int, value interface{}) (interface{}, error) {
	res, err := Bind[tcpType].BindMaps["Check"].handler("Check", value)
	if err != nil {
		return false, err
	}
	info := res.(map[string]string)
	fmt.Println(info)
	return "", errors.New("")
}

func (q *Queue) IncreasePushed(string) error {
	return nil
}

func (q *Queue) DetectPushed(string) (bool, error) {
	return true, nil
}
