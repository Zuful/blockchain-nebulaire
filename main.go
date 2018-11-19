package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

var blockchain Blockchain

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:8100",
			"http://localhost:4200",
		},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		/*AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},*/
		MaxAge: 12 * time.Hour,
	}))

	r.Use(cors.Default()) //allows all origins (must but changed when put in production)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/mine-block", mineBlockHandler)
	r.GET("/get-chain", getChainHandler)
	r.GET("/is-blockchain-valid", isBlockchainValidHandler)

	r.Run()
}

func init() {
	blockchain.createBlock(1, "0")
}
