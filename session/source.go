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
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
)

// Object contains all data for one session process with specific id.
type Object interface {
	Set(key, value interface{}) error //set session value
	Get(key interface{}) interface{}  //get session value
	Delete(key interface{}) error     //delete session value
	SessionID() string                //back current sessionID
	Cutoff() error                    //Cutoff all data
	SessionOut(w http.ResponseWriter) // open the resource & save data to provider & return the data

}

// Source contains global session methods and saved SessionObject.
// it can operate a SessionObject by its id.
type Source interface {
	SessionInit(int64, string) (Object, error)
	SessionRead(string) (Object, error)
	SessionDestroy(string) error
	SessionUpdate(string) error
	SessionGC(int64)
	//SessionAll() int //get all active session
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
	sourcename  string
}

// NewSource Create new Source with Source name.
func NewHandle(sourceName, objectName string, maxlifetime int64) (*Handle, error) {
	source, ok := sources[sourceName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", sourceName)
	}
	return &Handle{source: source, objectName: objectName, maxlifetime: maxlifetime, sourcename: sourceName}, nil
}

func (h *Handle) sessionId() (string, error) {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

//cookie base json config
type config struct {
	HashKey    string `json:"hashKey"`    //安全密钥 hash string
	CookieName string `json:"cookieName"` //cookie name
	Secure     bool   `json:"secure"`     //安全与否
	Maxage     int    `json:"maxage"`     //cookie max life time
}

func (h *Handle) SessionStart(w http.ResponseWriter, r *http.Request) (session Object, err error) {
	h.lock.Lock()
	defer h.lock.Unlock()
	if h.sourcename == "cookie" {
		cf := new(config)
		err := json.Unmarshal([]byte(h.objectName), cf)
		if err != nil {
			return nil, err
		}
		cookie, err := r.Cookie(cf.CookieName)
		if err != nil || cookie.Value == "" {
			session, _ = h.source.SessionInit(h.maxlifetime, h.objectName)
			//cookie := http.Cookie{Name: cf.CookieName, Value: url.QueryEscape(""), Path: "/", HttpOnly: true, MaxAge: int(h.maxlifetime)}
			//http.SetCookie(w, &cookie)
		} else {
			cookieVal, _ := url.QueryUnescape(cookie.Value)
			session, _ = h.source.SessionRead(cookieVal)
		}
	}
	if h.sourcename == "session" {
		cookie, err := r.Cookie(h.objectName)
		if err != nil || cookie.Value == "" {
			sid, errs := h.sessionId()
			if errs != nil {
				return nil, errs
			}
			session, _ = h.source.SessionInit(h.maxlifetime, sid)
			cookie := http.Cookie{Name: h.objectName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(h.maxlifetime)}
			http.SetCookie(w, &cookie)
		} else {
			sid, _ := url.QueryUnescape(cookie.Value)
			session, _ = h.source.SessionRead(sid)
		}
	}
	return
}

// SessionRelease Write cookie session to http response cookie
func (h *Handle) SessionRelease(w http.ResponseWriter) {
	//用于cookie w out 数据
}

// GC Start session gc process.
// it can do gc in times after gc lifetime.
func (h *Handle) GC() {
	h.source.SessionGC(h.maxlifetime)
}
