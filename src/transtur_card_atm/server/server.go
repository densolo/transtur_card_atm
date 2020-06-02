package server

import (
	"log"
	"time"
	"github.com/kuznetsovin/go_tachograph_card/tachocard_reader"
	"transtur_card_atm/readesm"
)


func ServeCardFiles(reader string) {
	log.Printf("Serving cards (reader: %s)", reader)

	for {
		cardData, err := WaitAndReadCard(reader)
		if err != nil {
			log.Printf("Failed to read a card data: " + err.Error())
			SendRedState(err.Error())
			time.Sleep(15 * time.Second)
			continue
		}

		SendBlueState("Parsing Card Data")
		ef, err := readesm.ParseData(cardData)
		if (err != nil) {
			log.Printf("Failed to parse DDD: %s", err)
			time.Sleep(15 * time.Second)
			continue
		}

		SendBlueState("Saving Card Data")
		fileName := CreateFileName(ef, time.Now())
		err = SaveDdd(cardData, fileName)
		if (err != nil) {
			log.Printf("Failed to save file %s: %s", fileName, err)
			time.Sleep(15 * time.Second)
			continue
		}

		SendBlueState("Completed")
		time.Sleep(15 * time.Second)
	}
}


func WaitAndReadCard(reader string) (dddFile []byte, err error) {

	SendGreyState("Checking card readers")
	if err = tachocard_reader.CheckEnableReaders(); err != nil {
		return []byte{}, err
	}

	SendGreyState("Waiting a card")
	indexReader, err := tachocard_reader.WaitCard(reader)
	if err != nil {
		return []byte{}, err
	}

	SendBlueState("Reading the card")
	return tachocard_reader.ReadСard("", indexReader)
}
