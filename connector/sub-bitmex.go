package connector

import (
	"encoding/json"

	"github.com/zulmaster/gomarkets/orderbook"

	"github.com/gorilla/websocket"
)

// BitMEX - BitMEX API
// https://www.bitmex.com/app/apiOverview
// Note that all Bitcoin quantities are returned in Satoshis: 1 XBt (Satoshi) = 0.00000001 XBT (Bitcoin)
type BitMEX struct {
}

// NewBitMex - доступ через bitmex api
func NewBitMex() SubConnector {
	m := BitMEX{}
	return &m
}

// GetWebSocketURL - адрес для подключения
func (bm *BitMEX) GetWebSocketURL() string {
	return "wss://www.bitmex.com/realtime"
}

// SendPingMessage - сообщение проверки связи
func (bm *BitMEX) SendPingMessage(c *websocket.Conn) error {
	return c.WriteMessage(websocket.TextMessage, []byte("ping"))
}

type jsonSubscribeOrderBook struct {
	Operation string   `json:"op"`
	Arguments []string `json:"args"`
}

// SendSubscribeOrderBook - сообщение для подписки на книгу
func (bm *BitMEX) SendSubscribeOrderBook(c *websocket.Conn, instrument string) error {
	req := jsonSubscribeOrderBook{
		Operation: "subscribe",
		Arguments: []string{"orderBookL2:" + instrument},
	}
	msg, err := json.Marshal(req)
	if err != nil {
		return err
	}
	return c.WriteMessage(websocket.TextMessage, msg)
}

// UpdateOrderBook - если сообщение про книгу, то обновление книги
func (bm *BitMEX) UpdateOrderBook(msg []byte, ob *orderbook.OrderBook) bool {
	return false
}
