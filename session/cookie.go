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
//| Cellgo Framework session/cookie file
//| All rights reserved: By cellgo.cn CopyRight
//| You are free to use the source code, but in the use of the process,
//| please keep the author information. Respect for the work of others
//| is respect for their own
//|-------------------------------------------------------------------
// Author:Tommy.Jin Dtime:2016-7-27

package session

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"net/http"
	"net/url"
	"sync"
)

var cookiequeue = &CookieQueue{}

func init() {
	RegisterSource("defaultCookie", cookiequeue)
}

// CookieSessionStore Cookie SessionStore
type Cookie struct {
	sid    string
	values map[interface{}]interface{} // session data
	lock   sync.RWMutex
}

// Set value to cookie session.
// the value are encoded as gob with hash block string.
func (c *Cookie) Set(key, value interface{}) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.values[key] = value
	return nil
}

// Get value from cookie session
func (c *Cookie) Get(key interface{}) interface{} {
	c.lock.RLock()
	defer c.lock.RUnlock()
	if v, ok := c.values[key]; ok {
		return v
	}
	return nil
}

// Delete value in cookie session
func (c *Cookie) Delete(key interface{}) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	delete(c.values, key)
	return nil
}

// Flush Clean all values in cookie session
func (c *Cookie) Cutoff() error {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.values = make(map[interface{}]interface{})
	return nil
}

// SessionID Return id of this cookie session
func (c *Cookie) SessionID() string {
	return c.sid
}

// SessionRelease Write cookie session to http response cookie
func (c *Cookie) SessionOut(w http.ResponseWriter) {
	str, err := encodeCookie(cookiequeue.block,
		cookiequeue.config.SecurityKey,
		cookiequeue.config.SecurityName,
		c.values)
	if err != nil {
		return
	}
	cookie := &http.Cookie{Name: cookiequeue.config.CookieName,
		Value:    url.QueryEscape(str),
		Path:     "/",
		HttpOnly: true,
		Secure:   cookiequeue.config.Secure,
		MaxAge:   cookiequeue.config.Maxage}
	http.SetCookie(w, cookie)
	return
}

//cookie base json config
type cookieConfig struct {
	SecurityKey  string `json:"securityKey"`  //安全密钥 hash string
	BlockKey     string `json:"blockKey"`     //AES密钥 gob encode hash string. it's saved as aes crypto.
	SecurityName string `json:"securityName"` //安全名  recognized name in encoded cookie string
	CookieName   string `json:"cookieName"`   //cookie name
	Secure       bool   `json:"secure"`       //安全与否
	Maxage       int    `json:"maxage"`       //cookie max life time
}

// CookieProvider Cookie session sources
type CookieQueue struct {
	maxlifetime int64
	config      *cookieConfig
	block       cipher.Block
}

// SessionInit Init cookie session sources with max lifetime and config json.
func (cq *CookieQueue) SessionInit(maxlifetime int64, config string) (Object, error) {
	cq.config = &cookieConfig{}
	err := json.Unmarshal([]byte(config), cq.config)
	if err != nil {
		return nil, err
	}
	if cq.config.BlockKey == "" {
		cq.config.BlockKey = string(generateRandomKey(16))
	}
	if cq.config.SecurityName == "" {
		cq.config.SecurityName = string(generateRandomKey(20))
	}
	cq.block, err = aes.NewCipher([]byte(cq.config.BlockKey))
	if err != nil {
		return nil, err
	}
	cq.maxlifetime = maxlifetime
	return nil, nil
}

// SessionRead Get Session in cooke.
// decode cooke string to map and put into Session with sid.
func (cq *CookieQueue) SessionRead(sid string) (Object, error) {
	maps, _ := decodeCookie(cq.block,
		cq.config.SecurityKey,
		cq.config.SecurityName,
		sid, cq.maxlifetime)
	if maps == nil {
		maps = make(map[interface{}]interface{})
	}
	rs := &Cookie{sid: sid, values: maps}
	return rs, nil
}

// SessionDestroy Implement method, no used.
func (cq *CookieQueue) SessionDestroy(sid string) error {
	return nil
}

// SessionGC Implement method, no used.
func (cq *CookieQueue) SessionGC(maxlifetime int64) {
	return
}

// SessionUpdate Implement method, no used.
func (cq *CookieQueue) SessionUpdate(sid string) error {
	return nil
}
