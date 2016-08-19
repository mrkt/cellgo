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
	"encoding/json"
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
	CarryInfo string      //Queue's carrying identification information
	Pushed    []string    //Queue's p ushed log
}

func (q *Queue) RegQueue(tcpType int, value interface{}) (interface{}, error) {
	res, err := Bind[tcpType].BindMaps["Reg"].handler("Reg", value)
	if err != nil {
		return false, err
	}
	if res == nil {
		return false, errors.New("The Data is not found.")
	}
	info := res.(map[string]string)
	q.FromInfo = value
	q.CarryInfo = info["CarryInfo"]
	var pushed []string
	if info["Pushed"] != "" {

		err = json.Unmarshal([]byte(info["Pushed"]), &pushed)
		if err != nil {
			fmt.Println(2)
			return false, err
		}
		q.Pushed = pushed
	}
	if ExchangeMap[SOCKETIO].Exchanges[info["Exchange"]].Queue[info["FromInfo"]] == nil {
		ExchangeMap[SOCKETIO].Exchanges[info["Exchange"]].Queue[info["FromInfo"]] = q
	}
	var result = make(map[string]string)
	result["json"] = "{\"Carry\":\"" + info["CarryInfo"] + "\"}"
	result["exchange"] = info["Exchange"]
	return result, nil
}

func (q *Queue) CheckQueue(tcpType int, value interface{}) (interface{}, error) {
	res, err := Bind[tcpType].BindMaps["Check"].handler("Check", value)
	if err != nil {
		return false, err
	}
	if res == nil {
		return false, errors.New("The Data is not found.")
	}
	info := res.(map[string]string)
	res = "{\"FromInfo\":\"" + info["FromInfo"] + "\",\"Exchange\":\"" + info["Exchange"] + "\"}"
	return res, nil
}

func (q *Queue) IncreasePushed(string) error {
	return nil
}

func (q *Queue) DetectPushed(string) (bool, error) {
	return true, nil
}
