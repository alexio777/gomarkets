package main

import (
	"log"
	"time"

	"github.com/zulmaster/gomarkets/connector"
)

func main() {
	log.Println("Примеры")
	bc := connector.NewBaseConnector(connector.NewBitMex())
	err := bc.Connect()
	if err != nil {
		log.Fatalln(err)
	}
	ob, err := bc.SubscribeOrderBook("XBTUSD")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("ob", ob)
	defer bc.Close()
	time.Sleep(time.Second * 5)
}
