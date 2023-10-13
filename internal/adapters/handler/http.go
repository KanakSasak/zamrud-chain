package handler

import (
	"github.com/gin-gonic/gin"
	"zamrud/internal/core/services"
)

type HttpHandler struct {
	BlockchainService *services.BlockchainService
}

func (h *HttpHandler) SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/blockchain", h.GetBlockchain)
	r.POST("/transaction", h.AddTransaction)
	r.POST("/mine", h.Mine)

	return r
}

func (h *HttpHandler) GetBlockchain(c *gin.Context) {
	chain := h.BlockchainService.GetChain()
	c.JSON(200, chain)
}

func (h *HttpHandler) AddTransaction(c *gin.Context) {
	// Logic to add transaction
}

func (h *HttpHandler) Mine(c *gin.Context) {
	// Logic to mine a block
}
