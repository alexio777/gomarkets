package main

import (
	"testing"

	"github.com/zulmaster/gomarkets/orderbook"
)

func TestNewOrderBook(t *testing.T) {
	ob := orderbook.NewOrderBook()
	if ob == nil {
		t.Errorf("NewOrderBook вернул nil")
	}
}
