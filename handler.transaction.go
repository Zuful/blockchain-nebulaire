package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func addTransactionHandler(c *gin.Context) {

	var code int
	var message string
	var index uint32
	var transaction Transaction

	err := c.Bind(&transaction)
	if err != nil {
		log.Println(err)
	}

	if transaction.Amount == 0 || transaction.Receiver == "" || transaction.Sender == "" {
		code = 400
		message = "Some elements of the transaction are missing"
	} else {
		code = 201
		index = addTransactions(transaction)
		message = "This transaction will be added in block " + fmt.Sprint(index)
	}

	c.JSON(code, gin.H{
		"message": message,
	})
}
