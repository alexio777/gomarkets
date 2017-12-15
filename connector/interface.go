package connector

import (
	"github.com/gorilla/websocket"
	"github.com/zulmaster/gomarkets/orderbook"
)

// SubConnector - универсальный интерфейс к биржам
// специфические для биржи функции.
// Используется в BaseConnector.
type SubConnector interface {
	GetWebSocketURL() string
	SendPingMessage(*websocket.Conn) error
	SendSubscribeOrderBook(*websocket.Conn, string) error
	UpdateOrderBook([]byte, *orderbook.OrderBook) bool
}
