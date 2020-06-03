
package server

import (
	"log"
	"time"
	"bytes"
	"errors"
	"strings"
	"github.com/jlaffaye/ftp"
	"transtur_card_atm/config"
)

func UploadFtp(cardData []byte, fileName string) error {
	appConfig := config.GetAppConfig()
	filePath := strings.TrimRight(appConfig.FtpUploadPath, "/") + "/" + fileName
	log.Printf("UploadFtp on %s into %s", appConfig.FtpServer, filePath)

	GlobalStateHandler.SendBlueState("Uploading a file on FTP")
	c, err := ConnectFtp()
	if err != nil {
		GlobalStateHandler.SendRedState("FTP upload failure")
		return err
	}
	defer CloseFtp(c)

	data := bytes.NewBufferString(string(cardData))
	err = c.Stor(filePath, data)
	if err != nil {
		GlobalStateHandler.SendRedState("FTP upload failure")
		log.Printf("Failed to upload a file on ftp: %s", err.Error())
		return err
	}

	log.Printf("File upload completed: %s", filePath)
	GlobalStateHandler.SendBlueState("Completed. Take your card.")
	return nil
}

func ConnectFtp() (*ftp.ServerConn, error) {
	appConfig := config.GetAppConfig()

	log.Printf("Connecting to FTP %s", appConfig.FtpServer)
	c, err := ftp.Dial(appConfig.FtpServer, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Printf("Failed to connect to FTP server '%s': %s", appConfig.FtpServer, err.Error())
		return nil, err
	}

	log.Printf("Login to FTP server '%s' with user %s", appConfig.FtpServer, appConfig.FtpUser)
	if appConfig.FtpUser == "" || appConfig.FtpPass == "" {
		err = errors.New("Ftp user and password cannot be empty")		
	} else {
		err = c.Login(appConfig.FtpUser, appConfig.FtpPass)
	}
	
	if err != nil {
		log.Printf("Failed to login to FTP server '%s' with user '%s': %s", appConfig.FtpServer, appConfig.FtpUser, err.Error())
		return nil, err
	}

	return c, err
}

func CloseFtp(c *ftp.ServerConn) {
	log.Printf("Closing FTP connection")
	if err := c.Quit(); err != nil {
		log.Printf("Failed to close ftp connection: %s", err.Error())
	}
}
