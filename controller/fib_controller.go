package controller

import (
	"strconv"

	"github.com/3380issei/fib_api/usecase"
	"github.com/gin-gonic/gin"
)

type FibController interface {
	GetFib(c *gin.Context)
}

type fibController struct {
	fu usecase.FibUsecase
}

func NewFibController(fu usecase.FibUsecase) FibController {
	return &fibController{fu}
}

func (fc *fibController) GetFib(c *gin.Context) {
	nString := c.Query("n")
	n, err := strconv.Atoi(nString)
	if err != nil || n <= 0 {
		c.JSON(400, gin.H{"error": "Invalid parameter"})
		return
	}

	nUint := uint(n)
	fibBigInt := fc.fu.GetFib(nUint)
	fib := fibBigInt.String()

	c.JSON(200, gin.H{"result": fib})
}
