package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"strconv"
	"time"
)

type Blockchain struct {
	Chain []Block
}

type Block struct {
	Index        int64  `json:"index"`
	Timestamp    string `json:"timestamp"`
	Proof        int64  `json:"proof"`
	PreviousHash string `json:"previousHash"`
}

func (blockchain *Blockchain) createBlock(proof int64, previousHash string) Block {

	var index int64 = int64(len(blockchain.Chain) + 1)
	var block Block = Block{
		Index:        index,
		Timestamp:    time.Now().String(),
		Proof:        proof,
		PreviousHash: previousHash,
	}

	blockchain.Chain = append(blockchain.Chain, block)

	return block
}

func (blockchain *Blockchain) getPreviousBlock() Block {
	var currentIndex int = len(blockchain.Chain)
	var previousIndex int = currentIndex - 1
	fmt.Println(previousIndex)
	return blockchain.Chain[previousIndex]
}

func (blockchain *Blockchain) proofOfWork(previousProof int64) int64 {
	var newProof int64 = 1
	var checkProof bool = false

	for checkProof == false {

		var hash string = getHashFromProblemToSolve(newProof, previousProof)

		if hash[:4] == "0000" {
			checkProof = true
		} else {
			newProof++
		}
	}

	return newProof
}

func (blockchain *Blockchain) hash(block Block) string {
	jsonBlock, err := json.Marshal(block)
	if err != nil {
		log.Println(err)
	}

	sha := sha256.New()
	sha.Write([]byte(jsonBlock))

	return hex.EncodeToString(sha.Sum(nil))
}

func (blockchain *Blockchain) isChainValid(chain []Block) bool {
	var previousBlock Block = chain[0]
	var blockIndex int64 = 1

	for blockIndex < int64(len(chain)) {
		//Check the block hash
		var block Block = chain[blockIndex]

		if block.PreviousHash != blockchain.hash(previousBlock) {
			return false
		}
		//Check the proof of work
		var previousProof int64 = previousBlock.Proof
		var proof int64 = block.Proof
		var hash string = getHashFromProblemToSolve(proof, previousProof)

		if hash[:4] != "0000" {
			return false
		}

		previousBlock = block
		blockIndex++
	}

	return true
}

func getHashFromProblemToSolve(proof int64, previousProof int64) string {
	sha := sha256.New()

	var operationRes float64 = math.Pow(float64(proof), 2) - math.Pow(float64(previousProof), 2)
	var strOperationRes string = strconv.FormatInt(int64(operationRes), 10)

	sha.Write([]byte(strOperationRes))

	return hex.EncodeToString(sha.Sum(nil))
}
