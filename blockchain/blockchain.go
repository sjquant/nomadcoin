package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type Block struct {
	Data     string
	Hash     string
	PrevHash string
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
	block := Block{data, "", bc.getLastHash()}
	block.calcHash()
	bc.blocks = append(bc.blocks, &block)
}

func (bc *blockChain) AllBlocks() []*Block {
	return bc.blocks
}

func GetBlockChain() *blockChain {
	if bc == nil {
		once.Do(func() {
			bc = &blockChain{}
		})
	}
	return bc
}
