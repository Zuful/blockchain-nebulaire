package main

type Transaction struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Amount   uint32 `json:"amount"`
}

/*
func (blockchain *Blockchain) getPreviousBlock() Block {
	var currentIndex int = len(blockchain.Chain)
	var previousIndex int = currentIndex - 1
	fmt.Println(previousIndex)
	return blockchain.Chain[previousIndex]
}
*/

func addTransactions(transaction Transaction) uint32 {
	transactions = append(transactions, transaction)
	var previousBlock Block = blockchain.getPreviousBlock()

	return previousBlock.Index + 1
}
