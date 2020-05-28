package server

import (
	"log"
	"time"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"github.com/kuznetsovin/go_tachograph_card/tachocard_reader"
	"transtur_card_atm/config"
	"transtur_card_atm/readesm"
)


func ServeCardFiles(reader string) {
	log.Printf("Serving cards (reader: %s)", reader)

	for {
		cardData, err := tachocard_reader.WaitAndReadCard(reader)
		if err != nil {
			log.Printf("Failed to read a card data: " + err.Error())
			time.Sleep(15 * time.Second)
			continue
		}

		ef, err := readesm.ParseData(cardData)
		if (err != nil) {
			log.Printf("Failed to parse DDD: %s", err)
			time.Sleep(15 * time.Second)
			continue
		}

		fileName := CreateFileName(ef, time.Now())
		err = SaveDdd(cardData, fileName)
		if (err != nil) {
			log.Printf("Failed to save file %s: %s", fileName, err)
			time.Sleep(15 * time.Second)
			continue
		}

		time.Sleep(15 * time.Second)
	}
}

func CreateFileName(ef readesm.EsmFile, now time.Time) string {
	// Format: C_YYYYMMDD_hhmm_FirstnameInitial_SecondName_CardNumber.DDD
	// Example: C_20200428_0956_A_Buta_000000000B8Z9000.DDD

	b, ok := ef.GetIdentification()
	if !ok {
		return fmt.Sprintf("C_%d%02d%02d_%02d%02d_%s_%s_%s.DDD",
			now.Year(), now.Month(), now.Day(), 
			now.Hour(), now.Minute(),
			"unknown",
			"unknown",
			"unknown",
		)
	}

	return fmt.Sprintf("C_%d%02d%02d_%02d%02d_%s_%s_%s.DDD",
		now.Year(), now.Month(), now.Day(), 
		now.Hour(), now.Minute(),
		b.CardHolderName.FirstName[:1],
		b.CardHolderName.Surname,
		b.CardNumber,
	)
}

func SaveDdd(cardData []byte, fileName string) error {	
	filePath := filepath.Join(config.GetUploadDir(), fileName)
	log.Printf("SaveDdd into %s", filePath)

	return ioutil.WriteFile(filePath, cardData, 0644)
}
