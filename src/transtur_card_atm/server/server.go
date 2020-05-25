package server

import (
	"log"
	"github.com/kuznetsovin/go_tachograph_card/tachocard_reader"
)


func ServeCardFiles(reader string) {
	log.Printf("Serving cards (reader: %s)", reader)

	err := tachocard_reader.SaveLocal(reader)
	if err != nil {
		log.Printf("Failed to save a card file: " + err.Error())
	}
}