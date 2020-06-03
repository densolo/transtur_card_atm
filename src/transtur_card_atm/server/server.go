package server

import (
	"log"
	"time"
	"github.com/kuznetsovin/go_tachograph_card/tachocard_reader"
	"transtur_card_atm/readesm"
	"transtur_card_atm/config"
)


func ServeCardFiles() {
	appConfig := config.GetAppConfig()
	reader := appConfig.ReaderName
	log.Printf("Serving cards (reader: %s)", reader)

	for {
		cardData, err := WaitAndReadCard(reader)
		if err != nil {
			log.Printf("Failed to read a card data: " + err.Error())
			GlobalStateHandler.SendRedState(err.Error())
			time.Sleep(15 * time.Second)
			continue
		}

		GlobalStateHandler.SendBlueState("Parsing Card Data")
		ef, err := readesm.ParseData(cardData)
		if (err != nil) {
			log.Printf("Failed to parse DDD: %s", err)
			time.Sleep(15 * time.Second)
			continue
		}

		fileName := CreateFileName(ef, time.Now())

		if appConfig.FtpServer != "" {
			err = UploadFtp(cardData, fileName)
		} else {
			err = SaveDdd(cardData, fileName)
		}

		time.Sleep(15 * time.Second)
	}
}


func WaitAndReadCard(reader string) (dddFile []byte, err error) {

	GlobalStateHandler.SendGreyState("Checking card readers")
	time.Sleep(time.Second)
	if err = tachocard_reader.CheckEnableReaders(); err != nil {
		return []byte{}, err
	}

	GlobalStateHandler.SendGreyState("Waiting a card")
	time.Sleep(time.Second)
	indexReader, err := tachocard_reader.WaitCard(reader)
	if err != nil {
		return []byte{}, err
	}

	GlobalStateHandler.SendBlueState("Reading the card")
	return tachocard_reader.ReadСard("", indexReader)
}
