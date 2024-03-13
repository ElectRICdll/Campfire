package service

import (
	"campfire/entity"
	"campfire/log"
	"campfire/util"
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

type SessionSender func(*websocket.Conn, int, []byte) error

type SessionService interface {
	NewSession(w http.ResponseWriter, r *http.Request, h http.Header, userid uint) error

	notify(n entity.Notification)
}

func NewSessionService() SessionService {
	res := &sessionService{
		generator: websocket.Upgrader{},
		pool: entity.SessionPool{
			Pool: make(map[uint]*websocket.Conn),
		},
		messageHandler: NewMessageService(),
	}
	return res
}

type sessionService struct {
	generator      websocket.Upgrader
	pool           entity.SessionPool
	messageHandler MessageService
}

func (s *sessionService) NewSession(w http.ResponseWriter, r *http.Request, h http.Header, senderId uint) error {
	conn, err := s.generator.Upgrade(w, r, h)
	if err != nil {
		return err
	}
	execute := s.pool.SessionExecution(senderId, conn)
	go execute(conn, s.handle)
	return nil
}

func (s *sessionService) notify(n entity.Notification) {
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

func (s *sessionService) handle(conn *websocket.Conn, wsType int, payload []byte) {
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

	msg := entity.Notification{}
	var tempMsg = struct {
		Timestamp   time.Time `json:"timestamp"`
		EType       int       `json:"e_type"`
		ReceiversID []uint    `json:"-"`
		Event       []byte    `json:"event_info"`
	}{}
	if err := json.Unmarshal(payload, &tempMsg); err != nil {
		s.importError(conn, err)
		return
	}
	event, err := s.eventSelector(tempMsg.EType)
	if err != nil {
		s.importError(conn, err)
		return
	}
	msg.EType = tempMsg.EType
	msg.Event = event
	if err := json.Unmarshal(payload, &msg); err != nil {
		s.importError(conn, err)
		return
	}

	if err := s.messageHandler.eventMessageHandler(&msg); err != nil {
		s.importError(conn, err)
	}
	s.notify(msg)
	return
}

func (s *sessionService) eventSelector(eType int) (entity.Event, error) {
	var res entity.Event
	if eType >= 0 && eType < len(entity.EventTypeIndex) {
		return nil, util.NewExternalError("invalid message type.")
	}
	if eventInstance, ok := entity.EventTypeIndex[eType-1].InnerType.(entity.Event); ok {
		res = eventInstance
	} else {
		return nil, errors.New("event type assertion failed")
	}
	return res, nil
}

func (s *sessionService) sendText(conn *websocket.Conn, wsType int, msg string) {
	err := conn.WriteMessage(wsType, ([]byte)(msg))
	if err != nil {
		log.Errorf("Error replying to client: %s", err)
	}
}

func (s *sessionService) sendJSON(conn *websocket.Conn, wsType int, data any) {
	msg, err := json.Marshal(data)
	if err != nil {
		log.Errorf("Illegal transmit!")
	}

	if err := conn.WriteMessage(wsType, msg); err != nil {
		log.Errorf("Error replying to client: %s", err)
	}
}

func (s *sessionService) importError(conn *websocket.Conn, err error) {
	if _, ok := err.(util.ExternalError); ok {
		s.sendText(conn, websocket.TextMessage, err.Error())
	} else {
		log.Error(err.Error())
	}
}

func (s *sessionService) inQueue() {

}
