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

type GetFibResponse struct {
	// GoのuInt64の最大値を超える可能性があるため、stringで返す
	Result  string `json:"result"`
	Message string `json:"message"`
}

const (
	InvalidParameterMessageNotInt     = "Invalid parameter (not integer)"
	InvalidParameterMessageNotNatural = "Invalid parameter (not natural number)"
)

func (fc *fibController) GetFib(c *gin.Context) {
	nString := c.Query("n")
	n, err := strconv.Atoi(nString)
	if err != nil {
		c.JSON(400, GetFibResponse{Message: InvalidParameterMessageNotInt})
		return
	}
	if n <= 0 {
		c.JSON(400, GetFibResponse{Message: InvalidParameterMessageNotNatural})
		return
	}
	nUint := uint(n)
	fibBigInt := fc.fu.GetFib(nUint)
	fibString := fibBigInt.String()

	c.JSON(200, GetFibResponse{Result: fibString})
}
