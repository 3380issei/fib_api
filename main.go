package main

import (
	"github.com/3380issei/fib_api/controller"
	"github.com/3380issei/fib_api/router"
	"github.com/3380issei/fib_api/usecase"
)

func main() {
	fu := usecase.NewFibUsecase()
	fc := controller.NewFibController(fu)
	r := router.NewRouter(fc)

	r.Run(":8080")
}
