package main

import (
	"reflect"
	"testing"

	"github.com/zulmaster/gomarkets/connector"
)

func TestNewBitMEX(t *testing.T) {
	m := connector.NewBitMex()
	if reflect.TypeOf(m).String() != "*connector.BitMEX" {
		t.Errorf("NewBitMex вернул не BitMEX тип: %s", reflect.TypeOf(m).String())
	}
}

func TestIsConnected(t *testing.T) {
	m := connector.NewBitMex()
	if m.IsConnected() != false {
		t.Errorf("IsConnected вернул true")
	}
}
