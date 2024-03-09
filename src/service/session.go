package service

import (
	"campfire/dao"
	"campfire/entity"
	. "campfire/log"
	"encoding/json"
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
	query          dao.CampDao
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

}

func (s *sessionService) handle(conn *websocket.Conn, wsType int, payload []byte) {
	Log.Info("Received new message")

	if wsType != websocket.TextMessage {
		Log.Infof("Other type message received: %s", payload)
		return
	} else {
		Log.Infof("Received text message: %s\n", payload)
		s.sendText(
			conn,
			websocket.TextMessage,
			"["+time.Now().String()+"]Received your message.",
		)
	}

	var message entity.Message
	err := json.Unmarshal(payload, &message)
	if err != nil {
		s.badData(conn, err)
	}

	members, err := s.query.MemberList(10, message.CampID)
	if err != nil {
		Log.Error(err.Error())
		return
	}
	users := []entity.User{}
	for _, member := range members {
		users = append(users, entity.User{
			ID: member.UserID,
		})
	}
	s.transmit(conn, users, message)
	return
}

func (s *sessionService) transmit(conn *websocket.Conn, users []entity.User, message entity.Message) {
	res, err := s.handlers[message.Type](message)
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
			s.sendText(
				conn,
				websocket.TextMessage,
				"["+time.Now().String()+"]Received new message.",
			)
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
