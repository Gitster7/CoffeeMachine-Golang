package main

import (
	"net/http"

	"github.com/kOstovic/CoffeeMachine/cmd/CoffeeMachineRestApiV1/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
