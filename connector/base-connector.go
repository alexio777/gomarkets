package connector

import (
	"container/list"
	"errors"
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/zulmaster/gomarkets/orderbook"
)

// BaseConnector - общие функции для интерфейса
type BaseConnector struct {
	Connection   *websocket.Conn
	IsConnected  bool
	SubConnector SubConnector
	Events       *list.List
}

type callbackFunction func([]byte) bool

// NewBaseConnector - создает новый доступ к бирже sub
func NewBaseConnector(sub SubConnector) *BaseConnector {
	bc := BaseConnector{
		SubConnector: sub,
		Events:       list.New(),
	}
	return &bc
}

// Connect - подключается к websocket api
func (bc *BaseConnector) Connect() (err error) {
	if bc.IsConnected {
		err = errors.New("Уже подключено")
		return
	}
	bc.Connection, _, err = websocket.DefaultDialer.Dial(bc.SubConnector.GetWebSocketURL(), nil)
	bc.IsConnected = err == nil
	// цикл получения сообщений
	heartbeatStop := make(chan bool)
	hearbeatSignal := make(chan bool)
	go func() {
		for {
			_, message, err := bc.Connection.ReadMessage()
			if err != nil {
				heartbeatStop <- true
				log.Println("read error", err)
				return
			}
			hearbeatSignal <- true
			element := bc.Events.Front()
			for element != nil {
				fn := element.Value.(callbackFunction)
				if ok := fn(message); ok {
					break
				}
				element = element.Next()
			}
			log.Printf("recv: %s", message[:80])
		}
	}()
	// цикл поддержки соединения
	// проверяется каждую секунду не прошло ли X секунд с момента последнего сообщения из websocket
	// если X секунд прошло, то посылается ping сообщение
	go func() {
		interval := time.Duration(constPingEverySeconds * time.Second)
		lastPing := time.Now()
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				if time.Since(lastPing) > interval {
					lastPing = time.Now()
					if err := bc.SubConnector.SendPingMessage(bc.Connection); err != nil {
						log.Println("ping message error", err)
						return
					}
				}
			case <-hearbeatSignal:
				lastPing = time.Now()
			case <-heartbeatStop:
				return
			}
		}
	}()
	return
}

// Close - отключается от websocket api
func (bc *BaseConnector) Close() {
	if bc.IsConnected {
		bc.Connection.Close()
		bc.IsConnected = false
	}
}

// RegisterEventCallback - регистрирует функцию для обработки сообщения
func (bc *BaseConnector) RegisterEventCallback(fn callbackFunction) {
	bc.Events.PushBack(fn)
}

// SubscribeOrderBook - подписка на книгу заявок
func (bc *BaseConnector) SubscribeOrderBook(instrument string) (ob *orderbook.OrderBook, err error) {
	err = bc.SubConnector.SendSubscribeOrderBook(bc.Connection, instrument)
	if err != nil {
		return
	}
	ob = orderbook.NewOrderBook(instrument)
	bc.RegisterEventCallback(func(msg []byte) bool {
		return bc.SubConnector.UpdateOrderBook(msg, ob)
	})
	return
}
