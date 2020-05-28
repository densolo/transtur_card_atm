package cardblocks

import (
	"transtur_card_atm/readesm/datatypes"
	"transtur_card_atm/readesm/utils"
)

const (
	Identification_BlockType = 0x0520
)

type Identification struct {
	datatypes.BlockHeader

	CardNumber string
    CardHolderName datatypes.Name
}

func (block Identification) GetBlockType() int {
	return block.BlockHeader.GetBlockType()
}

func (block Identification) Size() int {
	return block.BlockHeader.Size()
}

func (block *Identification) ParseData(data []byte) {
	block.CardNumber = utils.ReadString(data, 6, 16)

	block.CardHolderName = datatypes.Name{
		Surname: utils.ReadCodeString(data, 70, 36),
		FirstName: utils.ReadCodeString(data, 70+36, 36),
	}
}
