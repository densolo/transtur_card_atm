package server

import (
	"log"
	"time"
	"github.com/kuznetsovin/go_tachograph_card/tachocard_reader"
)


func ServeCardFiles(reader string) {
	log.Printf("Serving cards (reader: %s)", reader)

	for {
		err := tachocard_reader.SaveLocal(reader)
		if err != nil {
			log.Printf("Failed to save a card file: " + err.Error())
		}
		time.Sleep(15 * time.Second)
	}
}