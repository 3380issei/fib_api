package main

import (
	"os"

	"github.com/3380issei/fib_api/controller"
	"github.com/3380issei/fib_api/router"
	"github.com/3380issei/fib_api/usecase"
)

func main() {
	fibUsecase := usecase.NewFibUsecase()
	fibController := controller.NewFibController(fibUsecase)
	r := router.NewRouter(fibController)

	r.Run(os.Getenv("API_ADDRESS"))
}
