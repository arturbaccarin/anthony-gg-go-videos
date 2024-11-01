package core

type Blockchain struct {
	store     Storage
	headers   []*Header
	validator Validator
}

func (bc *Blockchain) AddBlock(b *Block) error {
	// validate
	return nil
}

// [0, 1, 2, 3] => 3 height
func (bc *Blockchain) Height() uint32 {
	return uint32(len(bc.headers) - 1)
}
