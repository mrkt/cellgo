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
//| Cellgo Framework socketio/socketio_room file
//| All rights reserved: By cellgo.cn CopyRight
//| You are free to use the source code, but in the use of the process,
//| please keep the author information. Respect for the work of others
//| is respect for their own
//|-------------------------------------------------------------------
// Author:Tommy.Jin Dtime:2016-08-06

package tcpip

//Socketio Room Operation type
type SocketioRoom struct {
	RoomName   string             //socketio room's name
	RoomNumber string             //socketio room's number
	Runner     map[string]*Runner //socketio room's runner
	PushedNum  int64              //socketio room's total push
	PulledNum  int64              //socketio room's total pull
}

//Socketio Create a room
func (s *SocketioRoom) CreateRoom(name string, number string) (bool, error) {
	return true, nil
}

//Socketio Renew room data
func (s *SocketioRoom) RenewRoom(style int, value map[string]interface{}) (bool, error) {
	return true, nil
}

//Socketio Destroy a room
func (s *SocketioRoom) DestroyRoom(number string) (bool, error) {
	return true, nil
}

//SocketioAllow an runner to enter the room
func (s *SocketioRoom) IncreaseRunner(runner *Runner, carryInfo string) (bool, error) {
	return true, nil
}
