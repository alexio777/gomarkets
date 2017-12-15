package orderbook

import (
	"github.com/huandu/skiplist"
	"github.com/shopspring/decimal"
)

// OrderBook - хранит, обновляет книгу заявок
// Продавцы - asks, сортируются по убыванию цены
// Покупатели - bids, сортируются по возрастанию цены
type OrderBook struct {
	Asks       *skiplist.SkipList
	Bids       *skiplist.SkipList
	Instrument string
}

// Row - строка стакана
type Row struct {
	Price  decimal.Decimal
	Volume int64
}

// NewOrderBook - создает новую книгу
func NewOrderBook(instrument string) *OrderBook {
	var greater skiplist.GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(decimal.Decimal).GreaterThan(rhs.(decimal.Decimal))
	}
	var less skiplist.LessThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(decimal.Decimal).LessThan(rhs.(decimal.Decimal))
	}
	ob := OrderBook{
		Asks:       skiplist.New(less),
		Bids:       skiplist.New(greater),
		Instrument: instrument,
	}
	return &ob
}
