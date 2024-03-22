package ws

import (
	"github.com/gorilla/websocket"
	"sync"
)

type DataHandler func(*websocket.Conn, int, []byte)
type SessionExecutor func(conn *websocket.Conn, handler DataHandler)

type SessionPool struct {
	sync.Mutex
	pool map[uint]struct {
		Conn  *websocket.Conn
		Token string
	}
}

func NewSessionPool() SessionPool {
	return SessionPool{
		pool: make(map[uint]struct {
			Conn  *websocket.Conn
			Token string
		}),
	}
}

func (p *SessionPool) AddSession(id uint, conn *websocket.Conn, token string) {
	p.Lock()
	defer p.Unlock()
	p.pool[id] = struct {
		Conn  *websocket.Conn
		Token string
	}{Conn: conn, Token: token}
}

func (p *SessionPool) Session() map[uint]struct {
	Conn  *websocket.Conn
	Token string
} {
	return (*p).pool
}

func (p *SessionPool) CloseSession(id uint) {
	p.Lock()
	defer p.Unlock()
	delete(p.pool, id)
}
