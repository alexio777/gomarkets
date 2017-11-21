package connector

// BitMEX - BitMEX API
// https://www.bitmex.com/app/apiOverview
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
