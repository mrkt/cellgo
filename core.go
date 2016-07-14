//|------------------------------------------------------------------
//|           __
//|        __/  \
//|     __/  \__/_
//|  __/  \__/ /  \
//| /  \__/  go\__/_
//| \__/_cell  __/  \
//|   /  \  __/  \__/
//|   \__/_/  \__/
//|     /  \__/
//|     \__/
//| ------------------------------------------------------------------
//| Cellgo Framework core type
//| All rights reserved: By cellgo.cn CopyRight
//| You are free to use the source code, but in the use of the process,
//| please keep the author information. Respect for the work of others
//| is respect for their own
//|-------------------------------------------------------------------
// Author:Tommy.Jin Dtime:2016-7-15

package cellgo

var (
	// CellCore is an core instance
	CellCore *Core
)

func init() {
	// create cellgo core
	CellCore = NewCore()
}

type Core struct {
	Handlers *ControllerRegister
	Server   *http.Server
}

// NewCore returns a new cellgo core.
func NewCore() *Core {
	cr := NewControllerRegister()
	core := &Core{Handlers: cr, Server: &http.Server{}}
	return Core
}

// Run cellgo core.
func (core *Core) Run() {
	//Listening.... && from config
	//loop:
	//		Handlers.routers
}
