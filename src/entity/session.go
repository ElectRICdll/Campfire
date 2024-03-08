package entity

import (
	. "campfire/log"
	"github.com/gorilla/websocket"
	"sync"
)

type DataHandler func(*websocket.Conn, int, []byte)
type SessionExecutor func(conn *websocket.Conn, handler DataHandler)

type SessionPool struct {
	sync.Mutex
	Pool map[ID]*websocket.Conn
}

func (p *SessionPool) SessionExecution(senderId ID, conn *websocket.Conn) SessionExecutor {
	p.Pool[senderId] = conn

	return func(conn *websocket.Conn, handle DataHandler) {
		defer func() {
			err := conn.Close()
			if err != nil {
				Log.Error(err.Error())
			}
		}()
		for {
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				Log.Error(err.Error())
				return
			}

			handle(conn, messageType, p)
		}
	}
}
