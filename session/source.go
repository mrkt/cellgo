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
//|------------------------------------------------------------------
//| Cellgo Framework session/source file
//| All rights reserved: By cellgo.cn CopyRight
//| You are free to use the source code, but in the use of the process,
//| please keep the author information. Respect for the work of others
//| is respect for their own
//|-------------------------------------------------------------------
// Author:Tommy.Jin Dtime:2016-7-27

package session

import (
	"fmt"
	"net/http"
	"sync"
)

// Object contains all data for one session process with specific id.
type Object interface {
	Set(key, value interface{}) error //set session value
	Get(key interface{}) interface{}  //get session value
	Delete(key interface{}) error     //delete session value
	SessionID() string                //back current sessionID
	SessionOut(w http.ResponseWriter) // open the resource & save data to provider & return the data
	Cutoff() error                    //Cutoff all data
}

// Source contains global session methods and saved SessionObject.
// it can operate a SessionObject by its id.
type Source interface {
	SessionInit(gclifetime int64, config string) error
	SessionRead(sid string) (Object, error)
	SessionRegenerate(oldsid, sid string) (Object, error)
	SessionDestroy(sid string) error
	SessionAll() int //get all active session
	SessionGC()
}

var sources = make(map[string]Source)

// RegisterSource makes a session Source available by the Source name.
// If Register is called twice with the same name or if driver is nil,
// it panics.
func RegisterSource(name string, source Source) {
	if source == nil {
		panic("session: Register Source is nil")
	}
	if _, dup := sources[name]; dup {
		panic("session: Register called twice for Source " + name)
	}
	sources[name] = source
}

type Handle struct {
	objectName  string     //private cookiename
	lock        sync.Mutex // protects session
	source      Source
	maxlifetime int64
}

// NewSource Create new Source with Source name.
func NewHandle(sourceName, objectName string, maxlifetime int64) (*Handle, error) {
	source, ok := sources[sourceName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", sourceName)
	}
	return &Handle{source: source, objectName: objectName, maxlifetime: maxlifetime}, nil
}
