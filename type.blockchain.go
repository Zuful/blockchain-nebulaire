package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Blockchain struct {
	Chain []Block  `json:"chain"`
	Nodes []string `json:"nodes"`
}

type Block struct {
	Index        uint32        `json:"index"`
	Timestamp    string        `json:"timestamp"`
	Proof        uint32        `json:"proof"`
	PreviousHash string        `json:"previousHash"`
	Transactions []Transaction `json:"transactions"`
}

func (blockchain *Blockchain) createBlock(proof uint32, previousHash string) Block {

	var index uint32 = uint32(len(blockchain.Chain) + 1)
	var block Block = Block{
		Index:        index,
		Timestamp:    time.Now().String(),
		Proof:        proof,
		PreviousHash: previousHash,
		Transactions: transactions,
	}

	transactions = make([]Transaction, 0) // now that the block has been mined, the transactions must be reset
	blockchain.Chain = append(blockchain.Chain, block)

	return block
}

func (blockchain *Blockchain) getPreviousBlock() Block {
	var currentIndex uint32 = uint32(len(blockchain.Chain))
	var previousIndex uint32 = currentIndex - 1
	fmt.Println(previousIndex)
	return blockchain.Chain[previousIndex]
}

func (blockchain *Blockchain) proofOfWork(previousProof uint32) uint32 {
	var newProof uint32 = 1
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
		var previousProof uint32 = previousBlock.Proof
		var proof uint32 = block.Proof
		var hash string = getHashFromProblemToSolve(proof, previousProof)

		if hash[:4] != "0000" {
			return false
		}

		previousBlock = block
		blockIndex++
	}

	return true
}

func (blockchain *Blockchain) addNode(address string) {
	parsedUrl, err := url.Parse(address)
	if err != nil {
		log.Print(err)
	}

	blockchain.Nodes = append(blockchain.Nodes, parsedUrl.Host)
}

func (blockchain *Blockchain) replaceChain() bool {
	var network []string = blockchain.Nodes
	var longestChain []Block
	var maxLength uint32 = uint32(len(blockchain.Chain))

	for _, oneNode := range network {
		statusCode, response := getChainRequest(oneNode)

		if statusCode == 200 {
			var chainResponse ChainResponse
			var err error = json.Unmarshal([]byte(response), &chainResponse)
			if err != nil {
				log.Print(err)
			}

			if chainResponse.Length > maxLength && blockchain.isChainValid(chainResponse.Chain) { //the chain of this iteration is the longest so far
				maxLength = chainResponse.Length
				longestChain = chainResponse.Chain
			}

			if len(longestChain) > 0 {
				blockchain.Chain = chainResponse.Chain //the current chain (chain of the type) must be updated to the longest one
				return true
			}
		}
	}

	return false
}

func getChainRequest(node string) (int, string) {
	resp, err := http.Get("http://" + node + "/get-chain")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return resp.StatusCode, string(body)
}

func getHashFromProblemToSolve(proof uint32, previousProof uint32) string {
	sha := sha256.New()

	var operationRes float64 = math.Pow(float64(proof), 2) - math.Pow(float64(previousProof), 2)
	var strOperationRes string = strconv.FormatInt(int64(operationRes), 10)

	sha.Write([]byte(strOperationRes))

	return hex.EncodeToString(sha.Sum(nil))
}
