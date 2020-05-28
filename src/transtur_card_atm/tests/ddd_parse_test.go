package tests

import (
	"time"
	"testing"
	"path/filepath"
	"io/ioutil"
	"log"
	"github.com/stretchr/testify/assert"
	"transtur_card_atm/readesm"
	"transtur_card_atm/server"
)

func TestParse(t *testing.T) {
	infile := filepath.Join("test-data", "CardPeek-butoi_ilie_2020-05-25.ddd")
	data, err := ioutil.ReadFile(infile)
	if (err != nil) {
		t.Fatalf("Test file error %s: %s", infile, err)
	}
	log.Printf("Read %d bytes", len(data))

	ef, err := readesm.ParseData(data)
	if (err != nil) {
		t.Fatalf("Test file error %s: %s", infile, err)
	}

	assert.Equal(t, len(ef.Blocks), 16, "Blocks count mismtach")

	b, ok := ef.GetIdentification()
	
	assert.Equal(t, ok, true, "Identification block not found")
	
	assert.Equal(t, b.CardNumber, "0000000005UYM001")
	assert.Equal(t, b.CardHolderName.FirstName, "Ilie")
	assert.Equal(t, b.CardHolderName.Surname, "Butoi")

	now := time.Date(2020, 05, 28, 10, 12, 15, 651387237, time.UTC)
	name := server.CreateFileName(ef, now)
	assert.Equal(t, name, "C_20200528_1012_I_Butoi_0000000005UYM001.DDD")
}
