package config

import (
	"os"
	"log"
	"path/filepath"
	"io/ioutil"
	"encoding/json"
)

type Config struct {
	Debug bool            `json:"debug"`
	ReaderName string     `json:"card_reader_name"`
	LocalSavePath string  `json:"local_save_path"`
	FtpServer string      `json:"ftp_server"`
	FtpUploadPath string  `json:"ftp_upload_path"`
	FtpUser string        `json:"ftp_user"`
	FtpPass string        `json:"ftp_pass"`
}

var (
	appConfig = Config{}
)

func GetAppRoot() string {
	log.Printf("Path: %s", os.Args[0])
	executable, err := filepath.Abs(os.Args[0])
	if err != nil {
		log.Fatalf("Cannot determine path root: " + err.Error())
		return ""
	}
	binPath := filepath.Dir(executable)
	return binPath  //strings.TrimSuffix(binPath, "bin")
}

func GetUploadDir() string {
	return filepath.Join(GetAppRoot(), "TCReaderUpload")
}

func GetAppConfig() Config {
	return appConfig
}

func LoadAppConfig() Config {
	path := filepath.Join(GetAppRoot(), "transtur_card_atm.json")
	appConfig = LoadConfig(path)
	configJson, _ := json.MarshalIndent(appConfig, "", "    ")
	log.Printf("App configuration: %s", string(configJson))
	return appConfig
}

func LoadConfig(path string) Config {
	config := Config{
		Debug: true,
		LocalSavePath: "TCReaderUpload",
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("Failed to read configuration file '%s': %s", path, err.Error())
		return config
	}

	err = json.Unmarshal([]byte(data), &config)
	if err != nil {
		log.Printf("Failed to parse configuration file '%s': %s", path, err.Error())
	}

	return config
}
