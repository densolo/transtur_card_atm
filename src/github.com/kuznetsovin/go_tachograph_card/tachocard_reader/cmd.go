package tachocard_reader

import (
	"github.com/howeyc/gopass"
)

func hideInput(param *string) error {
	field, err := gopass.GetPasswd()
	if err != nil {
		return err
	}

	*param = string(field)
	return err
}

func SaveLocal(reader string) error {
	dddFile, err := WaitAndReadCard(reader)
	if err != nil {
		return err
	}

	if err := SaveDdd(dddFile); err != nil {
		return err
	}
	return err
}

func WaitAndReadCard(reader string) (dddFile []byte, err error) {
	var pin string

	if err = checkEnableReaders(); err != nil {
		return []byte{}, err
	}

	indexReader, err := waitCard(reader)
	if err != nil {
		return []byte{}, err
	}

	return ReadСard(pin, indexReader)
}
