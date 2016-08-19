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
	"fmt"
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
	carryInfo string `json:"Carry"` //Connection function name
	value     string `json:"Vlaue"` //Disconnection function name
}

type back struct {
	State   int    `json:"State"`   //Connection function name
	Massage string `json:"Massage"` //Disconnection function name
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
							fmt.Println(1)
							res, err := ExchangeMap[SOCKETIO].Exchanges["2"].PushQueue(SOCKETIO, "")
							if err == nil {
								push := res.(map[string]map[string]string)
								for _, hp := range push {
									for k, p := range hp {
										so.BroadcastTo(k, "push", p)
									}
								}
							}
							//so.BroadcastTo("2", "push", "Hello!")
							time.Sleep(time.Second * 10)
							//fmt.Println("Hello")
						}
					}(so)
					initBool = false
				}
				log.Println("on connection")
				so.On(socketConf.Check, func(msg string) string {
					res, err := Queues.RegQueue(SOCKETIO, msg)
					if err != nil {
						callback, _ := callback(0, err.Error())
						return callback
					}
					result := res.(map[string]string)
					callback, _ := callback(1, result["json"])
					fmt.Println(callback)
					so.Join(result["exchange"])
					return callback

				})

				so.On(socketConf.Pull, func(msg string) string {
					//Exchange pull
					res, err := Queues.CheckQueue(SOCKETIO, msg)
					if err != nil {
						callback, _ := callback(0, err.Error())
						return callback
					}
					callback, _ := callback(1, res.(string))
					return callback

					/*res, err := Queues.CheckQueue(SOCKETIO, msg)
					if err != nil {
						callback, _ := callback(0, err.Error())
						return callback
					}
					callback, _ := callback(1, res.(string))
					return callback*/
				})
				so.On(socketConf.Disconn, func() {
					log.Println("on disconnect")
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
