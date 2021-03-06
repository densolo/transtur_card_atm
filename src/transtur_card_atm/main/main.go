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
	"github.com/kuznetsovin/go_tachograph_card/tachocard_reader"
)


func main() {
	log.Printf("Transtur Card ATM")
	initLogger()
	initConfig()
	
	go func(){
		server.WaitGuiWindow()
		server.ServeCardFiles()
	}()
	server.RunGui()
}

func initLogger() {
	appRoot := config.GetAppRoot()

	os.MkdirAll(filepath.Join(appRoot, "logs"), 0755)
	f, err := os.OpenFile(filepath.Join(appRoot, "logs/transtur_card_atm.log." + getTimeSuffix() + ".txt"), os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}	
	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)
}

func initConfig() {
	config.LoadAppConfig()
	uploadDir := config.GetUploadDir()
	os.MkdirAll(uploadDir, 0755)
	tachocard_reader.SetUploadDir(uploadDir)
}

func getTimeSuffix()(string){
	t := time.Now()
	return fmt.Sprintf("%d-%02d-%02d.%02d-%02d-%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}
