package readesm

import (
	"log"
	"errors"
	"transtur_card_atm/readesm/cardblocks"
	"transtur_card_atm/readesm/datatypes"
	"transtur_card_atm/readesm/utils"
)


type EsmFile struct {
	Blocks []datatypes.Block
}

func (ef EsmFile) GetIdentification() (block cardblocks.Identification, ok bool) {
	b, ok := ef.FindFirstBlock(cardblocks.Identification_BlockType)
	if !ok {
		return cardblocks.Identification{}, false
	}

	return b.(cardblocks.Identification), ok
}

func (ef EsmFile) FindFirstBlock(blockType int) (block datatypes.Block, ok bool) {
	for _, b := range ef.Blocks {
		if b.GetBlockType() == blockType {
			log.Printf("Block FOUND")
			return b, true
		}
	}
	return datatypes.BlockHeader{}, false
}


func ParseData(data []byte)(EsmFile, error) {
	ef := EsmFile{}
	pos := 0

	for pos < len(data) {
		block, err := blockFactory(data, pos)
		if err != nil {
			return ef, err
		}
		ef.Blocks = append(ef.Blocks, block)
		pos += block.Size()
	}

	return ef, nil
}


func blockFactory(data []byte, pos int)(block datatypes.Block, err error) {
	if (data[pos] == 0x76) {
		return nil, errors.New("Unsupported block type")
	} else {
		return cardBlockFactory(data, pos)
	}
}


func cardBlockFactory(data []byte, pos int)(block datatypes.Block, err error) {
	data = data[pos:]
	blockType := utils.ReadInt16(data, 0)
	dataSize := utils.ReadInt16(data, 3)
	hasSignature := false

	bytesLeft := len(data[5 + dataSize:])
	if (bytesLeft >= 5 + 128) {
		nextType := utils.ReadInt16(data, 5 + dataSize)
		sz1 := nextType == blockType
		if (sz1 && utils.ReadInt8(data, 5 + dataSize + 2) == 1) {
			hasSignature = true
		}
	}

	blockHeader := datatypes.BlockHeader{
		BlockType: blockType,
		DataSize: dataSize,
		HasSignature: hasSignature,
	}

	log.Printf("Block type: 0x%x size %d signature %t", blockType, dataSize+5, hasSignature)


	if (blockType == cardblocks.Identification_BlockType) {
		cb := cardblocks.Identification{
			BlockHeader: blockHeader,
		}
		cb.ParseData(data)
		block = cb
	} else {
		block = blockHeader
	}
	return block, nil
}