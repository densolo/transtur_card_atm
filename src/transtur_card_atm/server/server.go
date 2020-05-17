package server

import (
	"log"
	"github.com/kuznetsovin/go_tachograph_card/tachocard_reader"
)


func ServeCardFiles() {
	log.Printf("Serving cards")

	err := tachocard_reader.SaveLocal()
	if err != nil {
		log.Printf("Failed to save a card file: " + err.Error())
	}
}