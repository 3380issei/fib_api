package router

import (
	"github.com/3380issei/fib_api/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter(fc controller.FibController) *gin.Engine {
	r := gin.Default()
	r.GET("/fib", fc.GetFib)

	return r
}
