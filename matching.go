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

type ServeHTTP struct {
}

type ControllerRegister struct {
	routers map[string]*ServeHTTP
}

func NewControllerRegister() *ControllerRegister {
	cr := &ControllerRegister{
		routers: make([]string{}),
	}
	return cr
}

func (p *ControllerRegister) Add(pattern string) {
	server := &ServeHTTP{}
	p.routers[pattern] = server
}

func (s *ServeHTTP) workHTTP(rw http.ResponseWriter, r *http.Request) {
	//matching......
}
