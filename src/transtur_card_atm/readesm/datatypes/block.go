package datatypes

type Block interface {
	GetBlockType() int
	Size() int
}

type BlockHeader struct {
	BlockType int
	DataSize int
	HasSignature bool
}

func (block BlockHeader) GetBlockType() int {
	return block.BlockType
}

func (block BlockHeader) Size() int {
	sz := block.DataSize + 5
	if block.HasSignature {
		sz += 5 + 128
	}
	return sz
}
