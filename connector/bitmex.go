package connector

import (
	"errors"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const constBitMEXWebSocketURL = "wss://www.bitmex.com/realtime"

// BitMEX - BitMEX API
// https://www.bitmex.com/app/apiOverview
// Note that all Bitcoin quantities are returned in Satoshis: 1 XBt (Satoshi) = 0.00000001 XBT (Bitcoin)
type BitMEX struct {
	Connection *websocket.Conn
	BaseConnector
}

// NewBitMex - доступ через bitmex api
func NewBitMex() Connector {
	m := BitMEX{}
	return &m
}

// Connect - подключается к websocket api
func (bm *BitMEX) Connect() (err error) {
	if bm.IsConnected {
		err = errors.New("")
		return
	}
	bm.Connection, _, err = websocket.DefaultDialer.Dial(constBitMEXWebSocketURL, nil)
	bm.IsConnected = err == nil
	// цикл получения сообщений
	heartbeatStop := make(chan bool)
	hearbeatSignal := make(chan bool)
	go func() {
		for {
			_, message, err := bm.Connection.ReadMessage()
			if err != nil {
				heartbeatStop <- true
				log.Println("read error", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()
	// цикл поддержки соединения
	// проверяется каждую секунду не прошло ли 5 секунд с момента последнего сообщения из websocket
	// если 5 секунд прошло, то посылается ping сообщение
	go func() {
		interval := time.Duration(5 * time.Second)
		lastPing := time.Now()
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				if time.Since(lastPing) > interval {
					lastPing = time.Now()
					if err := bm.Connection.WriteMessage(websocket.TextMessage, []byte("ping")); err != nil {
						log.Println("ping error", err)
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
func (bm *BitMEX) Close() {
	if bm.IsConnected {
		bm.Connection.Close()
		bm.IsConnected = false
	}
}
