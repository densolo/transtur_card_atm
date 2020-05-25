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
	var pin string

	if err := checkEnableReaders(); err != nil {
		return err
	}

	indexReader, err := waitCard(reader)
	if err != nil {
		return err
	}

	// fmt.Print("PIN: ")
	// if err := hideInput(&pin); err != nil {
	// 	return err
	// }
	
	dddFile, err := ReadСard(pin, indexReader)
	if err != nil {
		return err
	}

	if err := SaveDdd(dddFile); err != nil {
		return err
	}
	return err
}
