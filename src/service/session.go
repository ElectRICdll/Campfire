package service

import (
	"campfire/entity"
	. "campfire/log"
	"encoding/json"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

type SessionSender func(*websocket.Conn, int, []byte) error

type SessionService interface {
	NewSession(w http.ResponseWriter, r *http.Request, h http.Header, userid int) error
}

func NewSessionService() SessionService {
	res := &sessionService{
		generator: websocket.Upgrader{},
		pool: entity.SessionPool{
			Pool: make(map[entity.ID]*websocket.Conn),
		},
		messageHandler: messageService{},
	}
	res.handlers = []MessageHandler{
		res.messageHandler.unknownMessageHandler,
		res.messageHandler.textMessageHandler,
		res.messageHandler.binaryMessageHandler,
		res.messageHandler.binaryMessageHandler,
		res.messageHandler.textMessageHandler,
		res.messageHandler.textMessageHandler,
		res.messageHandler.textMessageHandler,
		res.messageHandler.binaryMessageHandler,
		res.messageHandler.eventMessageHandler,
	}
	return res
}

type sessionService struct {
	generator      websocket.Upgrader
	pool           entity.SessionPool
	messageHandler MessageService
	handlers       []MessageHandler
}

func (s *sessionService) NewSession(w http.ResponseWriter, r *http.Request, h http.Header, senderId int) error {
	conn, err := s.generator.Upgrade(w, r, h)
	if err != nil {
		return err
	}
	execute := s.pool.SessionExecution((entity.ID)(senderId), conn)
	go execute(conn, s.handle)
	return nil
}

func (s *sessionService) handle(conn *websocket.Conn, wsType int, payload []byte) {
	Log.Info("Received new message")

	if wsType != websocket.TextMessage {
		Log.Infof("Other type message received: %s", payload)
		return
	}

	Log.Infof("Received text message: %s\n", payload)
	s.sendText(
		conn,
		websocket.TextMessage,
		"["+time.Now().String()+"]Received your message.",
	)

	var message entity.Message
	err := json.Unmarshal(payload, &message)
	if err != nil {
		s.badData(conn, err)
	}
	s.transmit(conn, message)
	return
}

func (s *sessionService) transmit(conn *websocket.Conn, message entity.Message) {
	res, users, err := s.handlers[message.Type](message)
	if err != nil {
		if _, ok := err.(entity.ExternalError); ok {
			Log.Errorf("External error: %s", err.Error())
			s.badData(conn, err)
			return
		}
		Log.Errorf(err.Error())
		return
	}

	for _, value := range users {
		if connRes, ok := s.pool.Pool[value.ID]; ok {
			s.send(
				connRes,
				websocket.TextMessage,
				res,
			)
		}
		s.inQueue()
	}

	Log.Infof("Message has transmitted to the other.")
}

func (s *sessionService) sendText(conn *websocket.Conn, wsType int, message string) {
	s.send(conn, wsType, []byte(message))
}

func (s *sessionService) send(conn *websocket.Conn, wsType int, data []byte) {
	err := conn.WriteMessage(wsType, data)
	if err != nil {
		Log.Errorf("Error replying to client: %s", err)
	}
}

func (s *sessionService) badData(conn *websocket.Conn, err error) {
	s.send(conn, websocket.TextMessage, []byte(err.Error()))
}

func (s *sessionService) inQueue() {

}
