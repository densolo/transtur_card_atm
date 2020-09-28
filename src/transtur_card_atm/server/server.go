package server

import (
	"fmt"
	"log"
	"time"
	"transtur_card_atm/config"
	"transtur_card_atm/readesm"

	"github.com/kuznetsovin/go_tachograph_card/tachocard_reader"
)

func ServeCardFiles() {
	appConfig := config.GetAppConfig()
	reader := appConfig.ReaderName
	log.Printf("Serving cards (reader: %s)", reader)

	for {
		indexReader, cardData, err := WaitAndReadCard(reader)
		if err != nil {
			log.Printf("Failed to read a card data: " + err.Error())
			GlobalStateHandler.SendRedState(fmt.Sprintf("Read card error: %s", err.Error()))
			time.Sleep(5 * time.Second)
			continue
		}

		GlobalStateHandler.SendBlueState("Parsing Card Data")
		ef, err := readesm.ParseData(cardData)
		if err != nil {
			log.Printf("Failed to parse DDD: %s", err)
			time.Sleep(5 * time.Second)
			continue
		}

		fileName := CreateFileName(ef, time.Now().UTC())

		if appConfig.FtpServer != "" {
			SaveDdd(cardData, fileName)

			err = UploadFtp(cardData, fileName)
		} else {
			err = SaveDdd(cardData, fileName)
		}

		if err == nil {
			tachocard_reader.WaitCardEjected(indexReader)
		}
		time.Sleep(2 * time.Second)
	}
}

func WaitAndReadCard(reader string) (indexReader int, dddFile []byte, err error) {
	GlobalStateHandler.SendGreyState("Checking card readers")
	time.Sleep(time.Second)
	if err = tachocard_reader.CheckEnableReaders(); err != nil {
		return -1, []byte{}, err
	}

	GlobalStateHandler.SendGreyState("Waiting a card")
	time.Sleep(time.Second)
	indexReader, err = tachocard_reader.WaitCard(reader)
	if err != nil {
		return -1, []byte{}, err
	}

	GlobalStateHandler.SendBlueState("Reading the card")
	data, err := tachocard_reader.ReadСard("", indexReader)
	return indexReader, data, err
}
