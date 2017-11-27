package connector

// BitMEX - BitMEX API
// https://www.bitmex.com/app/apiOverview
// Note that all Bitcoin quantities are returned in Satoshis: 1 XBt (Satoshi) = 0.00000001 XBT (Bitcoin)
type BitMEX struct {
	isConnected bool
}

// NewBitMex - доступ через BitMEX API
func NewBitMex() Connector {
	m := BitMEX{}
	return &m
}

// IsConnected - активно ли подключение к бирже?
func (m *BitMEX) IsConnected() bool {
	return m.isConnected
}
