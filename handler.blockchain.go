package main

import "github.com/gin-gonic/gin"

func mineBlockHandler(c *gin.Context) {
	var previousBlock Block = blockchain.getPreviousBlock()
	var previousProof int64 = previousBlock.Proof
	var proof int64 = blockchain.proofOfWork(previousProof)
	var previousHash string = blockchain.hash(previousBlock)
	var block Block = blockchain.createBlock(proof, previousHash)

	c.JSON(200, gin.H{
		"message":      "Congratulations, you just mined a block!",
		"index":        block.Index,
		"timestamp":    block.Timestamp,
		"proof":        block.Proof,
		"previousHash": block.PreviousHash,
	})
}

func getChainHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"chain":  blockchain.Chain,
		"length": len(blockchain.Chain),
	})
}

func isBlockchainValidHandler(c *gin.Context) {

	c.JSON(200, gin.H{
		"isValid": blockchain.isChainValid(blockchain.Chain),
	})
}
