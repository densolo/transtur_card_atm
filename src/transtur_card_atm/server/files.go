package server

import (
	"log"
	"time"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"transtur_card_atm/config"
	"transtur_card_atm/readesm"
)


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
	GlobalStateHandler.SendBlueState("Saving card data into a file")
	filePath := filepath.Join(config.GetUploadDir(), fileName)
	log.Printf("SaveDdd into %s", filePath)

	err := ioutil.WriteFile(filePath, cardData, 0644)
	if err != nil {
		log.Printf("Failed to save file %s: %s", filePath, err.Error())
		GlobalStateHandler.SendRedState("Save file failure")
		return nil
	}
		
	log.Printf("File save completed: %s", filePath)
	GlobalStateHandler.SendBlueState("Completed. Take your card.")
	return nil
}

