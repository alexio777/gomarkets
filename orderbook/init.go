package orderbook

// OrderBook - хранит, обновляет книгу заявок
type OrderBook struct {
}

// NewOrderBook - создает новую книгу
func NewOrderBook() *OrderBook {
	ob := OrderBook{}
	return &ob
}
