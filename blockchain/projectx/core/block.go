package core

import (
	"bytes"
	"encoding/gob"
	"io"
	"projectx/crypto"
	"projectx/types"
)

type Header struct {
	Version       uint32
	DataHash      types.Hash
	PrevBlockHash types.Hash
	Height        uint32 // height of a block refers to its position in the blockchain
	Timestamp     int64
}

type Block struct {
	*Header
	Transactions []Transaction
	Validator    crypto.PublicKey
	Signature    *crypto.Signature

	// Cached version of the header hash
	hash types.Hash
}

func NewBlock(h *Header, tx []Transaction) *Block {
	return &Block{
		Header:       h,
		Transactions: tx,
	}
}

func (b *Block) Sign(privKey crypto.PrivateKey) *crypto.Signature {
	// sign, err := privKey.Sign()
}

func (b *Block) Decode(r io.Reader, dec Decoder[*Block]) error {
	return dec.Decode(r, b)
}

func (b *Block) Encode(w io.Writer, enc Encoder[*Block]) error {
	return enc.Encode(w, b)
}

func (b *Block) Hash(hasher Hasher[*Block]) types.Hash {
	if b.hash.IsZero() {
		b.hash = hasher.Hash(b)
	}

	return b.hash
}

func (b *Block) HeaderBytes() []byte {
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	enc.Encode(b.Header)

	return buf.Bytes()
}

// func (b *Block) Hash() types.Hash {
// 	buf := &bytes.Buffer{}
// 	b.Header.EncodeBinary(buf)

// 	if b.hash.IsZero() {
// 		b.hash = types.Hash(sha256.Sum256(buf.Bytes()))
// 	}

// 	return b.hash
// }
