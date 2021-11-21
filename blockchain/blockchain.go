package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type block struct {
	data     string
	hash     string
	prevHash string
}

type blockChain struct {
	blocks []*block
}

var bc *blockChain
var once sync.Once

func (bc *blockChain) getLastHash() string {
	if len(bc.blocks) == 0 {
		return ""
	}
	return bc.blocks[len(bc.blocks)-1].hash
}

func (b *block) calcHash() {
	b.hash = fmt.Sprintf("%x", sha256.Sum256([]byte(b.data+b.prevHash)))
}

func (bc *blockChain) AddBlock(data string) {
	block := block{data, "", bc.getLastHash()}
	block.calcHash()
	bc.blocks = append(bc.blocks, &block)
}

func (bc *blockChain) PrintBlocks() {
	for _, b := range bc.blocks {
		fmt.Printf("Data: %s\n", b.data)
		fmt.Printf("Hash: %s\n", b.hash)
		fmt.Printf("Prev Hash: %s\n\n", b.prevHash)
	}
}

func GetBlockChain() *blockChain {
	if bc == nil {
		once.Do(func() {
			bc = &blockChain{}
		})
	}
	return bc
}
