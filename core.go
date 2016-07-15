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

import (
	"fmt"
	"net/http"
)

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
type Mux struct {
}

func (p *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		CellCore.Handlers.workHTTP(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

// NewCore returns a new cellgo core.
func NewCore() *Core {
	cr := NewControllerRegister()
	core := &Core{Handlers: cr, Server: &http.Server{}}
	return core
}

// Run cellgo core.
func (core *Core) Run() {
	fmt.Println("Cellgo Core Runing...")
	mux := &Mux{}
	http.ListenAndServe(":9090", mux)
}

func (core *Core) RegisterController(title string, c ControllerInterface) *Core {
	core.Handlers.Add(title, c)
	return core
}
