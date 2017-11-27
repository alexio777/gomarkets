package orderbook

import (
	"github.com/huandu/skiplist"
	"github.com/shopspring/decimal"
)

// OrderBook - хранит, обновляет книгу заявок
// Продавцы - asks, сортируются по убыванию цены
// Покупатели - bids, сортируются по возрастанию цены
type OrderBook struct {
	Asks *skiplist.SkipList
	Bids *skiplist.SkipList
}

// Row - строка стакана
type Row struct {
	Price decimal.Decimal
}

// NewOrderBook - создает новую книгу
func NewOrderBook() *OrderBook {
	var greater skiplist.GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(Row).Price.GreaterThan(rhs.(Row).Price)
	}
	var less skiplist.LessThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(Row).Price.LessThan(rhs.(Row).Price)
	}
	ob := OrderBook{
		Asks: skiplist.New(greater),
		Bids: skiplist.New(less),
	}
	return &ob
}
