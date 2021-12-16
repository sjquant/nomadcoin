package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"sync"
)

type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevHash,omitempty"`
	Height   int    `json:"height"`
}

type blockChain struct {
	blocks []*Block
}

var bc *blockChain
var once sync.Once

func (bc *blockChain) getLastHash() string {
	if len(bc.blocks) == 0 {
		return ""
	}
	return bc.blocks[len(bc.blocks)-1].Hash
}

func (b *Block) calcHash() {
	b.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte(b.Data+b.PrevHash)))
}

func (bc *blockChain) AddBlock(data string) {
	block := Block{data, "", bc.getLastHash(), len(bc.blocks) + 1}
	block.calcHash()
	bc.blocks = append(bc.blocks, &block)
}

func (bc *blockChain) AllBlocks() []*Block {
	return bc.blocks
}

var ErrNotFound = errors.New("block not found")

func (bc *blockChain) GetBlock(height int) (*Block, error) {
	if height > len(bc.blocks) {
		return nil, ErrNotFound
	}
	return bc.blocks[height-1], nil
}

func GetBlockChain() *blockChain {
	if bc == nil {
		once.Do(func() {
			bc = &blockChain{}
		})
	}
	return bc
}
