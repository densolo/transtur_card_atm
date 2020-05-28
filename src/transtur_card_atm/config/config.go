package config

import (
	"os"
	"log"
	"path/filepath"
	//"strings"
)

type Config struct {
	
}

func GetAppRoot() string {
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
