package main

import (
	"os"
	"io"
	"log"
	"time"
	"fmt"
	"path/filepath"
	"transtur_card_atm/config"
	"transtur_card_atm/server"
)


func main() {
	log.Printf("Transtur Card ATM")
	initLogger()
	server.ServeCardFiles()
}


func initLogger() {
	appRoot := config.GetAppRoot()

	os.MkdirAll(filepath.Join(appRoot, "logs"), 0755)
	f, err := os.OpenFile(filepath.Join(appRoot, "logs/transtur_card_atm.log." + getTimeSuffix()), os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}	
	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)
}

func getTimeSuffix()(string){
	t := time.Now()
	return fmt.Sprintf("%d-%02d-%02d.%02d%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute())
}
