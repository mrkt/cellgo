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
	"log"
	"net/http"

	"github.com/googollee/go-socket.io"
)

func init() {
	for _, v := range Tcp[SOCKETIO] {
		server := v.Handle.(*socketio.Server)
		server.On("connection", func(so socketio.Socket) {
			log.Println("on connection")
			so.On("chat message with ack", func(msg string) string {
				return msg
			})
			so.On("disconnection", func() {
				log.Println("on disconnect")
			})
		})
		server.On("error", func(so socketio.Socket, err error) {
			log.Println("error:", err)
		})

		http.Handle("/socket.io/", server)
		log.Println("Serving at localhost:5000...")
		log.Fatal(http.ListenAndServe(":5000", nil))
	}

}
