package main

import (
	"testing"

	"github.com/shopspring/decimal"

	"github.com/zulmaster/gomarkets/orderbook"
)

func TestNewOrderBook(t *testing.T) {
	ob := orderbook.NewOrderBook()
	if ob == nil {
		t.Error("NewOrderBook вернул nil")
	}
}

func TestAsksOrder(t *testing.T) {
	ob := orderbook.NewOrderBook()
	ob.Asks.Set(decimal.NewFromFloat(100.1), orderbook.Row{})
	ob.Asks.Set(decimal.NewFromFloat(100.5), orderbook.Row{})
	ob.Asks.Set(decimal.NewFromFloat(100.3), orderbook.Row{})
	validElements := []decimal.Decimal{
		decimal.NewFromFloat(100.5),
		decimal.NewFromFloat(100.3),
		decimal.NewFromFloat(100.1),
	}
	validElementsPointer := 0
	element := ob.Asks.Front()
	for element != nil {
		if validElements[validElementsPointer].Cmp(element.Key().(decimal.Decimal)) != 0 {
			t.Errorf("Asks в книге заявок отсортированы не верно, ожидалось %v, пришло %v", validElements[validElementsPointer], element.Key().(decimal.Decimal))
		}
		validElementsPointer++
		element = element.Next()
	}
}

func TestBidsOrder(t *testing.T) {
	ob := orderbook.NewOrderBook()
	ob.Bids.Set(decimal.NewFromFloat(100.1), orderbook.Row{})
	ob.Bids.Set(decimal.NewFromFloat(100.5), orderbook.Row{})
	ob.Bids.Set(decimal.NewFromFloat(100.3), orderbook.Row{})
	validElements := []decimal.Decimal{
		decimal.NewFromFloat(100.1),
		decimal.NewFromFloat(100.3),
		decimal.NewFromFloat(100.5),
	}
	validElementsPointer := 0
	element := ob.Bids.Front()
	for element != nil {
		if validElements[validElementsPointer].Cmp(element.Key().(decimal.Decimal)) != 0 {
			t.Errorf("Bids в книге заявок отсортированы не верно, ожидалось %v, пришло %v", validElements[validElementsPointer], element.Key().(decimal.Decimal))
		}
		validElementsPointer++
		element = element.Next()
	}
}
