package main

import "github.com/sjquant/nomadcoin/blockchain"

func main() {
	chain := blockchain.GetBlockChain()
	chain.AddBlock("First Block")
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")
	chain.PrintBlocks()
}
