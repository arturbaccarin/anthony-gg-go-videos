package core

import (
	"fmt"
	"projectx/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddBlock(t *testing.T) {
	bc := newBlockchainWithGenesis(t)

	lenBlocks := 1000
	for i := 0; i < lenBlocks; i++ {
		block := randomBlockWithSignature(t, uint32(i+1), getPrevBlockHash(t, bc, uint32(i+1)))
		assert.Nil(t, bc.AddBlock(block))
	}

	assert.Equal(t, bc.Height(), uint32(lenBlocks))
	assert.Equal(t, len(bc.headers), lenBlocks+1)

	assert.NotNil(t, bc.AddBlock(randomBlock(89)))
}

func TestNewBlockchain(t *testing.T) {
	bc, err := NewBlockchain(randomBlock(0))
	assert.Nil(t, err)
	assert.NotNil(t, bc)
	assert.Equal(t, bc.Height(), uint32(0))

	fmt.Println(bc.Height())
}

func TestHasBlock(t *testing.T) {
	bc := newBlockchainWithGenesis(t)

	assert.True(t, bc.HasBlock(0))
}

func TestGetHeader(t *testing.T) {
	bc := newBlockchainWithGenesis(t)

	lenBlocks := 1000
	for i := 0; i < lenBlocks; i++ {
		block := randomBlockWithSignature(t, uint32(i+1))
		assert.Nil(t, bc.AddBlock(block))

		header, err := bc.GetHeader(uint32(i + 1))
		assert.Nil(t, err)
		assert.Equal(t, header, block.Header)
	}

}

func TestAddBlockToHigh(t *testing.T) {
	bc := newBlockchainWithGenesis(t)

	assert.NotNil(t, bc.AddBlock(randomBlockWithSignature(t, 3)))
}

func newBlockchainWithGenesis(t *testing.T) *Blockchain {
	bc, err := NewBlockchain(randomBlock(0))
	assert.Nil(t, err)

	return bc
}

func getPrevBlockHash(t *testing.T, bc *Blockchain, height uint32) types.Hash {
	prevHeader, err := bc.GetHeader(height - 1)
	assert.Nil(t, err)

	return BlockHasher{}.Hash(prevHeader)
}
