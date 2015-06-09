package session

import (
	"container/list"// Package list implements a doubly linked list. 
	"time"
	"sync"
	// "fmt"
)
func init() {
    pder.sessions = make(map[string]*list.Element, 0)
}

var pder = &ProviderMemory{list: list.New()}

//Provider接口，用以表征session管理器底层存储结构
type Provider interface {
    SessionInit(sid string) (Session, error)
    SessionRead(sid string) (Session, error)
    SessionDestroy(sid string) error
    SessionGC(maxLifeTime int64)
}

//ProviderMemory implements provider interface
//sid is a unique string for every connection
//it contains a map from sid to a session instance

type ProviderMemory struct {
    lock     sync.Mutex               //用来锁
    sessions map[string]*list.Element //list.Element is struct made up of interface
    list     *list.List               //list.List is a doubly linked list, 用来做gc
}

//init a session instance by session id, put it into provider's map and list
func (pder *ProviderMemory) SessionInit(sid string) (Session, error) {
    pder.lock.Lock()
    defer pder.lock.Unlock()
    v := make(map[interface{}]interface{}, 0)
    newsess := &SessionMemory{sid: sid, timeAccessed: time.Now(), value: v }
    element := pder.list.PushBack(newsess)
    pder.sessions[sid] = element
    return newsess, nil
}

//get session by session id
func (pder *ProviderMemory) SessionRead(sid string) (Session, error) {
    if element, ok := pder.sessions[sid]; ok {
        return element.Value.(*SessionMemory), nil
    } else {
        sess, err := pder.SessionInit(sid)
        return sess, err
    }
    return nil, nil
}

//delete session from provider by session id
func (pder *ProviderMemory) SessionDestroy(sid string) error {
    if element, ok := pder.sessions[sid]; ok {
        delete(pder.sessions, sid)
        pder.list.Remove(element)
        return nil
    }
    return nil
}

//if session has not been used for maxlifetime, remove it
func (pder *ProviderMemory) SessionGC(maxlifetime int64) {
    pder.lock.Lock()
    defer pder.lock.Unlock()

    for {
        element := pder.list.Back()
        if element == nil {
            break
        }
        if (element.Value.(*SessionMemory).timeAccessed.Unix() + maxlifetime) < time.Now().Unix() {
            pder.list.Remove(element)
            delete(pder.sessions, element.Value.(*SessionMemory).sid)
        } else {
            break
        }
    }
}

//update session by session id, change timeAccessed to lasted, move element to front
func (pder *ProviderMemory) SessionUpdate(sid string) error {
    pder.lock.Lock()
    defer pder.lock.Unlock()
    if element, ok := pder.sessions[sid]; ok {
        element.Value.(*SessionMemory).timeAccessed = time.Now()
        pder.list.MoveToFront(element)
        return nil
    }
    return nil
}