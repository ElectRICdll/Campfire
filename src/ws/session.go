package ws

import (
	"campfire/auth"
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
		sec: auth.SecurityInstance,
		generator: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		pool:         NewSessionPool(),
		eventHandler: EventService{},
	}
	return res
}

type SessionService struct {
	sec          auth.SecurityGuard
	generator    websocket.Upgrader
	pool         SessionPool
	eventHandler EventService
}

func (s *SessionService) NewSession(w http.ResponseWriter, r *http.Request, h http.Header, token string) error {
	res, err := s.sec.WSTokenVerify(token)
	if err != nil {
		return util.NewExternalError("unauthorized")
	}

	conn, err := s.generator.Upgrade(w, r, h)
	if err != nil {
		return err
	}
	s.pool.AddSession(res, conn, token)
	go func(conn *websocket.Conn, handle DataHandler) {
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
	}(conn, s.handle)
	return nil
}

func (s *SessionService) NotifyByEvent(event Event, eType int) error {
	res := Notification{
		Timestamp: time.Now(),
		EType:     eType,
		Event:     event,
	}
	if err := s.eventHandler.HandleEvent(&res); err != nil {
		return err
	}
	s.Notify(res)
	return nil
}

func (s *SessionService) Notify(n Notification) {
	for _, value := range n.ReceiversID {
		if res, ok := s.pool.Session()[value]; ok {
			s.sendJSON(
				res.Conn,
				websocket.TextMessage,
				n,
			)
		}
		s.inQueue()
	}
}

func (s *SessionService) handle(conn *websocket.Conn, wsType int, payload []byte) {
	log.Info("Received new message")

	if wsType == websocket.TextMessage {
		log.Infof("Received text message: %s\n", payload)
		s.sendText(
			conn,
			websocket.TextMessage,
			"["+time.Now().String()+"]Received your message.",
		)
	} else {
		log.Infof("Other type message received: %s", payload)
		s.sendText(conn, websocket.PongMessage, "pong")
		return
	}

	var tempMsg = struct {
		Timestamp   time.Time `json:"timestamp"`
		OperatorID  uint      `json:"userID"`
		EType       int       `json:"eType"`
		ReceiversID []uint    `json:"-"`
		Event       []byte    `json:"eventData"`
		Token       string    `json:"token"`
	}{}
	if err := json.Unmarshal(payload, &tempMsg); err != nil {
		s.sendError(conn, err)
		return
	}

	//if s.pool.Session()[tempMsg.OperatorID].Conn != conn {
	//	s.sendError(conn, util.NewExternalError("unauthorized"))
	//	defer conn.Close()
	//	return
	//}
	if tempMsg.EType == PingEventType {
		s.sendText(conn, websocket.TextMessage, "pong")
		return
	}

	var msg = &Notification{}
	msg.EType = tempMsg.EType
	msg.Event = s.eventSelector(tempMsg.EType)
	if err := json.Unmarshal(payload, msg); err != nil {
		s.sendError(conn, err)
		return
	}
	if err := s.eventHandler.HandleEvent(msg); err != nil {
		s.sendError(conn, err)
	}
	s.Notify(*msg)
	return
}

func (s *SessionService) eventSelector(eType int) Event {
	res, _ := EventsByType[eType]()
	return res
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
