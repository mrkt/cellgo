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
//| Cellgo Framework session/session file
//| All rights reserved: By cellgo.cn CopyRight
//| You are free to use the source code, but in the use of the process,
//| please keep the author information. Respect for the work of others
//| is respect for their own
//|-------------------------------------------------------------------
// Author:Tommy.Jin Dtime:2016-7-27

package session

import (
	"container/list"
	"net/http"
	"sync"
	"time"
)

var sessionqueue = &SessionQueue{sessions: make(map[string]*list.Element, 0), list: list.New()}

func init() {
	RegisterSource("memorySess", sessionqueue)
}

type Session struct {
	sid          string                      //session id唯一标示
	timeAccessed time.Time                   //最后访问时间
	value        map[interface{}]interface{} //session里面存储的值
}

func (s *Session) Set(key, value interface{}) error {
	s.value[key] = value
	sessionqueue.SessionUpdate(s.sid)
	return nil
}

func (s *Session) Get(key interface{}) interface{} {
	sessionqueue.SessionUpdate(s.sid)
	if v, ok := s.value[key]; ok {
		return v
	} else {
		return nil
	}
	return nil
}

func (s *Session) Delete(key interface{}) error {
	delete(s.value, key)
	sessionqueue.SessionUpdate(s.sid)
	return nil
}

func (s *Session) SessionID() string {
	return s.sid
}

func (s *Session) Cutoff() error {
	return nil
}

func (s *Session) SessionOut(w http.ResponseWriter) {
}

type SessionQueue struct {
	lock        sync.Mutex               //用来锁
	sessions    map[string]*list.Element //用来存储在内存
	list        *list.List               //用来做gc
	maxlifetime int64
}

// SessionInit Init cookie session sources with max lifetime and sid.
// maxlifetime is ignored.
func (sq *SessionQueue) SessionInit(maxlifetime int64, sid string) (Object, error) {
	sq.lock.Lock()
	defer sq.lock.Unlock()
	v := make(map[interface{}]interface{}, 0)
	newsess := &Session{sid: sid, timeAccessed: time.Now(), value: v}
	element := sq.list.PushBack(newsess)
	sq.sessions[sid] = element
	sq.maxlifetime = maxlifetime
	return newsess, nil
}

func (sq *SessionQueue) SessionRead(sid string) (Object, error) {
	if element, ok := sq.sessions[sid]; ok {
		return element.Value.(*Session), nil
	} else {
		sess, err := sq.SessionInit(sq.maxlifetime, sid)
		return sess, err
	}
	return nil, nil
}

func (sq *SessionQueue) SessionDestroy(sid string) error {
	if element, ok := sq.sessions[sid]; ok {
		delete(sq.sessions, sid)
		sq.list.Remove(element)
		return nil
	}
	return nil
}

func (sq *SessionQueue) SessionGC(maxlifetime int64) {
	sq.lock.Lock()
	defer sq.lock.Unlock()

	for {
		element := sq.list.Back()
		if element == nil {
			break
		}
		if (element.Value.(*Session).timeAccessed.Unix() + maxlifetime) < time.Now().Unix() {
			sq.list.Remove(element)
			delete(sq.sessions, element.Value.(*Session).sid)
		} else {
			break
		}
	}
}

func (sq *SessionQueue) SessionUpdate(sid string) error {
	sq.lock.Lock()
	defer sq.lock.Unlock()
	if element, ok := sq.sessions[sid]; ok {
		element.Value.(*Session).timeAccessed = time.Now()
		sq.list.MoveToFront(element)
		return nil
	}
	return nil
}
