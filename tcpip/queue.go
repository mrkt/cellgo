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

var (
	Queues *Queue
)

func init() {
	Queues = &Queue{Pushed: make(map[string]bool)}
}

//Queue Operation type
type Queue struct {
	FromInfo  interface{}     //Queue's enter identification information
	CarryInfo string          //Queue's carrying identification information
	Pushed    map[string]bool //Queue's p ushed log
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
	if ExchangeMap[SOCKETIO].Exchanges[info["Exchange"]].Queue[info["FromInfo"]] == nil {
		ExchangeMap[SOCKETIO].Exchanges[info["Exchange"]].Queue[info["FromInfo"]] = &Queue{}
		ExchangeMap[SOCKETIO].Exchanges[info["Exchange"]].Queue[info["FromInfo"]].FromInfo = value
		ExchangeMap[SOCKETIO].Exchanges[info["Exchange"]].Queue[info["FromInfo"]].CarryInfo = info["CarryInfo"]
		ExchangeMap[SOCKETIO].Exchanges[info["Exchange"]].Queue[info["FromInfo"]].Pushed = make(map[string]bool)
	}
	/*var pushed []string
	if info["Pushed"] != "" {

		err = json.Unmarshal([]byte(info["Pushed"]), &pushed)
		if err != nil {
			return false, err
		}
		q.Pushed = pushed
	}*/

	var result = make(map[string]string)
	//Find first push info
	if true {
		p, err := Bind[tcpType].BindMaps["Push"].handler("Push", value)
		if err == nil {
			push := p.(map[string]map[string]string)
			var (
				firstPush     []string
				jsonFirstPush string = ""
			)
			for hk, hp := range push {
				for k, p := range hp {
					if !ExchangeMap[SOCKETIO].Exchanges[info["Exchange"]].Queue[info["FromInfo"]].Pushed[hk] {
						if info["Exchange"] == k {
							firstPush = append(firstPush, p)
						}
					}
				}
				ExchangeMap[SOCKETIO].Exchanges[info["Exchange"]].Queue[info["FromInfo"]].Pushed[hk] = true
			}
			jsonFirstPush += "["
			for k, v := range firstPush {
				if k != 0 {
					jsonFirstPush += ","
				}
				jsonFirstPush += v

			}
			jsonFirstPush += "]"
			result["json"] = "{\"Carry\":\"" + info["CarryInfo"] + "\",\"Value\":" + jsonFirstPush + "}"

		} else {
			result["json"] = "{\"Carry\":\"" + info["CarryInfo"] + "\",\"Value\":\"\"}"
		}
	} else {
		result["json"] = "{\"Carry\":\"" + info["CarryInfo"] + "\"}"
	}
	result["exchange"] = info["Exchange"]
	return result, nil
}

func (q *Queue) CheckQueue(tcpType int, value interface{}) (interface{}, error) {
	res, err := Bind[tcpType].BindMaps["Check"].handler("Check", value)
	if err != nil || res == nil {
		return false, errors.New("The Data is not found.")
	}
	return res, nil
}

func (q *Queue) IncreasePushed(string) error {
	return nil
}

func (q *Queue) DetectPushed(string) (bool, error) {
	return true, nil
}
