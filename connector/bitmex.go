package connector

// BitMEX - BitMEX API
// https://www.bitmex.com/app/apiOverview
type BitMEX struct {
}

// NewBitMex - доступ через BitMEX API
func NewBitMex() Connector {
	m := BitMEX{}
	return &m
}

// Connect - соединяется с биржей BitMEX
func (m *BitMEX) Connect() {

}
