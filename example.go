package main

import (
	"log"
	"time"

	"github.com/zulmaster/gomarkets/connector"
)

func main() {
	log.Println("Примеры")
	bm := connector.NewBitMex()
	err := bm.Connect()
	if err != nil {
		log.Fatalln(err)
	}
	time.Sleep(time.Second * 10)
	defer bm.Close()
}
