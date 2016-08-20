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
//| Cellgo Framework tcpip/socketio file
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
	"log"
	"net/http"

	"time"

	"github.com/googollee/go-socket.io"
)

type socketConf struct {
	Conn    string `json:"Conn"`    //Connection function name
	Disconn string `json:"Disconn"` //Disconnection function name
	Error   string `json:"Error"`   //Error function name
	Auth    string `json:"Auth"`    //Auth function name
	Push    string `json:"Push"`    //Push content function name
	Pull    string `json:"Pull"`    //Pull content function name
	Check   string `json:"Check"`   //Check content function name
}

type from struct {
	Carry string `json:"Carry"` //Connection function name
	Value string `json:"Value"` //Disconnection function name
}

type back struct {
	State   int    `json:"State"`   //Connection function name
	Message string `json:"Message"` //Disconnection function name
}

var (
	initBool bool = true
)

func RunSocketIO() {
	for _, v := range Tcp[SOCKETIO] {
		go func(v *TcpRun) {
			//CreateExchange
			CreateExchange(SOCKETIO)
			//fmt.Println(ExchangeMap[SOCKETIO].Exchanges["2"].ExchangeName)
			//Queues.RegQueue(SOCKETIO, "42")
			//Queues.CheckQueue(SOCKETIO, "326700")
			//ExchangeMap[SOCKETIO].Exchanges["4"].PullQueue(SOCKETIO, "4")

			socketConf := &socketConf{}
			err := json.Unmarshal([]byte(v.TcpConf), socketConf)
			if err != nil {
				log.Fatal("socketio [", v.TcpName, "] error:", err)
			}
			checkDefault(socketConf)

			server := v.Handle.(*socketio.Server)
			server.On(socketConf.Conn, func(so socketio.Socket) {
				//so.Join("Seckill")
				if initBool == true {
					go func(so socketio.Socket) {
						for {
							//Exchange push
							for k, v := range ExchangeMap[SOCKETIO].Exchanges {
								res, err := v.PushQueue(SOCKETIO, "")
								if err == nil {
									for _, vr := range res.([]string) {
										callback := "{\"State\":\"1\",\"Message\":" + vr + "}"
										so.BroadcastTo(k, "push", callback)

									}
								}
							}
							time.Sleep(time.Second * 1) //stop 1 sec check)
						}
					}(so)
					initBool = false
				}
				//log.Println("on connection")
				so.On(socketConf.Check, func(msg string) string {
					res, err := Queues.RegQueue(SOCKETIO, msg)
					if err != nil {
						callback, _ := callback(0, err.Error())
						return callback
					}
					result := res.(map[string]string)
					//callback, _ := callback(1, result["json"])
					callback := "{\"State\":\"1\",\"Message\":" + result["json"] + "}"
					so.Join(result["exchange"])
					return callback

				})

				so.On(socketConf.Pull, func(msg string) string {
					//Decomposition message
					fromInfo := &from{}
					err = json.Unmarshal([]byte(msg), fromInfo)
					if err != nil {
						callback, _ := callback(0, errors.New("The Data is error.").Error())
						return callback
					}
					//Exchange pull
					res, err := Queues.CheckQueue(SOCKETIO, fromInfo.Carry)
					if err != nil {
						callback, _ := callback(0, err.Error())
						return callback
					}
					info := res.(map[string]string)
					info["Value"] = fromInfo.Value
					resPull, errPull := ExchangeMap[SOCKETIO].Exchanges[info["Exchange"]].PullQueue(SOCKETIO, info)
					if errPull != nil {
						callback, _ := callback(0, err.Error())
						return callback
					}
					callback, _ := callback(1, resPull.(string))

					return callback
				})
				so.On(socketConf.Disconn, func() {
					//log.Println("on disconnect")
				})
			})
			server.On(socketConf.Error, func(so socketio.Socket, err error) {
				log.Println("error:", err)
			})
			http.Handle(v.Route, server)
			log.Println(v.TcpName, "Serving at", v.Addr, "to", v.Route)
			log.Fatal(http.ListenAndServe(v.Addr, nil))
		}(v)
	}
}

func checkDefault(s *socketConf) {
	switch {
	case s.Conn == "":
		s.Conn = "connection"
		fallthrough
	case s.Disconn == "":
		s.Disconn = "disconnection"
		fallthrough
	case s.Error == "error":
		s.Error = "error"
		fallthrough
	case s.Auth == "":
		s.Auth = "auth"
		fallthrough
	case s.Push == "":
		s.Push = "push"
		fallthrough
	case s.Pull == "":
		s.Pull = "pull"
		fallthrough
	case s.Check == "":
		s.Check = "check"
	}
}

func callback(state int, message string) (string, error) {
	callback := back{state, message}
	res, err := json.Marshal(callback)
	if err != nil {
		return "", err
	}
	return string(res), nil
}
