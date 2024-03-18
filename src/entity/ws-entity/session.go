package wsentity

import (
	"campfire/log"
	"github.com/gorilla/websocket"
	"sync"
)

type DataHandler func(*websocket.Conn, int, []byte)
type SessionExecutor func(conn *websocket.Conn, handler DataHandler)

type SessionPool struct {
	sync.Mutex
	Pool map[uint]*websocket.Conn
}

func (p *SessionPool) SessionExecution(senderId uint, conn *websocket.Conn) SessionExecutor {
	p.Pool[senderId] = conn

	return func(conn *websocket.Conn, handle DataHandler) {
		defer func() {
			err := conn.Close()
			if err != nil {
				log.Error(err.Error())
			}
		}()
		for {
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				log.Error(err.Error())
				return
			}

			handle(conn, messageType, p)
		}
	}
}