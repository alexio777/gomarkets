package orderbook

import (
	"testing"
)

func TestNewOrderBook(t *testing.T) {
	ob := NewOrderBook()
	if ob == nil {
		t.Errorf("NewOrderBook вернул nil")
	}
}
