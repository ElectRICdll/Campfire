package ws_service

import (
	"campfire/entity/ws-entity"
	"campfire/log"
	"campfire/util"
	"encoding/json"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

type SessionSender func(*websocket.Conn, int, []byte) error

func NewSessionService() *SessionService {
	res := &SessionService{
		generator: websocket.Upgrader{},
		pool: wsentity.SessionPool{
			Pool: make(map[uint]*websocket.Conn),
		},
		eventHandler: EventService{},
	}
	return res
}

type SessionService struct {
	generator    websocket.Upgrader
	pool         wsentity.SessionPool
	eventHandler EventService
}

func (s *SessionService) NewSession(w http.ResponseWriter, r *http.Request, h http.Header, senderId uint) error {
	conn, err := s.generator.Upgrade(w, r, h)
	if err != nil {
		return err
	}
	execute := s.pool.SessionExecution(senderId, conn)
	go execute(conn, s.handle)
	return nil
}

func (s *SessionService) Notify(n Notification) {
	for _, value := range n.ReceiversID {
		if connRes, ok := s.pool.Pool[value]; ok {
			s.sendJSON(
				connRes,
				websocket.TextMessage,
				n,
			)
		}
		s.inQueue()
	}
}

func (s *SessionService) handle(conn *websocket.Conn, wsType int, payload []byte) {
	log.Info("Received new message")

	if wsType != websocket.TextMessage {
		log.Infof("Other type message received: %s", payload)
		return
	} else {
		log.Infof("Received text message: %s\n", payload)
		s.sendText(
			conn,
			websocket.TextMessage,
			"["+time.Now().String()+"]Received your message.",
		)
	}

	msg := Notification{}
	var tempMsg = struct {
		Timestamp   time.Time `json:"timestamp"`
		EType       int       `json:"e_type"`
		ReceiversID []uint    `json:"-"`
		Event       []byte    `json:"event_info"`
	}{}
	if err := json.Unmarshal(payload, &tempMsg); err != nil {
		s.sendError(conn, err)
		return
	}
	event, err := s.eventSelector(tempMsg.EType)
	if err != nil {
		s.sendError(conn, err)
		return
	}
	msg.EType = tempMsg.EType
	msg.Event = event
	if err := json.Unmarshal(payload, &msg); err != nil {
		s.sendError(conn, err)
		return
	}

	if err := s.eventHandler.HandleEvent(&msg); err != nil {
		s.sendError(conn, err)
	}
	s.Notify(msg)
	return
}

func (s *SessionService) eventSelector(eType int) (wsentity.Event, error) {
	res, _ := wsentity.NewEventByType((wsentity.EventType)(eType))
	return res, nil
}

func (s *SessionService) sendText(conn *websocket.Conn, wsType int, msg string) {
	err := conn.WriteMessage(wsType, ([]byte)(msg))
	if err != nil {
		log.Errorf("Error replying to client: %s", err)
	}
}

func (s *SessionService) sendJSON(conn *websocket.Conn, wsType int, data any) {
	msg, err := json.Marshal(data)
	if err != nil {
		log.Errorf("Illegal transmit!")
	}

	if err := conn.WriteMessage(wsType, msg); err != nil {
		log.Errorf("Error replying to client: %s", err)
	}
}

func (s *SessionService) sendError(conn *websocket.Conn, err error) {
	if _, ok := err.(util.ExternalError); ok {
		s.sendText(conn, websocket.TextMessage, err.Error())
	} else {
		log.Error(err.Error())
	}
}

func (s *SessionService) inQueue() {

}
