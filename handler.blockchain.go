package main

import "github.com/gin-gonic/gin"

type ChainResponse struct {
	Chain  []Block `json:"chain"`
	Length uint32  `json:"length"`
}

func mineBlockHandler(c *gin.Context) {
	var previousBlock Block = blockchain.getPreviousBlock()
	var previousProof uint32 = previousBlock.Proof
	var proof uint32 = blockchain.proofOfWork(previousProof)
	var previousHash string = blockchain.hash(previousBlock)

	var transaction Transaction
	transaction.Sender = nodeAddress
	transaction.Receiver = "Yamani"
	transaction.Amount = 1
	addTransactions(transaction)

	var block Block = blockchain.createBlock(proof, previousHash)

	c.JSON(200, gin.H{
		"message":      "Congratulations, you just mined a block!",
		"index":        block.Index,
		"timestamp":    block.Timestamp,
		"proof":        block.Proof,
		"previousHash": block.PreviousHash,
		"transactions": block.Transactions,
	})
}

func getChainHandler(c *gin.Context) {
	var response ChainResponse
	response.Chain = blockchain.Chain
	response.Length = uint32(len(blockchain.Chain))

	c.JSON(200, response)
}

func isBlockchainValidHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"isValid": blockchain.isChainValid(blockchain.Chain),
	})
}
