package connector

import (
	"reflect"
	"testing"
)

func TestNewBitMEX(t *testing.T) {
	m := NewBitMex()
	if reflect.TypeOf(m).String() != "*connector.BitMEX" {
		t.Errorf("NewBitMex вернул не BitMEX тип: %s", reflect.TypeOf(m).String())
	}
}

func TestIsConnected(t *testing.T) {
	m := NewBitMex()
	if m.IsConnected() != false {
		t.Errorf("IsConnected вернул true")
	}
}
