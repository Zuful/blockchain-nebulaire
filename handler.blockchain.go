package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

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
	transaction.Receiver = "Minor A"
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

func replaceChainHandler(c *gin.Context) {

	var message string
	var isChainReplaced bool = blockchain.replaceChain()

	if isChainReplaced == true {
		message = "The node has different chains, so the chains was replaced by the longest one."
	} else {
		message = "All good, the chain is the largest one."
	}

	c.JSON(200, gin.H{
		"message": message,
		"chain":   blockchain.Chain,
	})
}

func connectNodesHandler(c *gin.Context) {

	var code int
	var message string
	var err error
	var postedNodes struct {
		Nodes []string `json:"nodes"`
	}

	err = c.Bind(&postedNodes)
	if err != nil {
		log.Println(err)
	}

	var nodes []string = postedNodes.Nodes

	if len(nodes) == 0 {
		code = 400
		message = "No nodes."
	} else {

		for _, oneNodeAddress := range nodes {
			blockchain.addNode(oneNodeAddress)
		}

		code = 201
		message = "All the nodes are now connected, the Phonecoin now contains the following nodes: "

	}

	c.JSON(code, gin.H{
		"message":    message,
		"totalNodes": blockchain.Nodes,
	})
}
