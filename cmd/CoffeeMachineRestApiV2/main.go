package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kOstovic/CoffeeMachine/cmd/CoffeeMachineRestApiV2/controllers"
	_ "github.com/kOstovic/CoffeeMachine/cmd/CoffeeMachineRestApiV2/docs"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
)

// @title CoffeeMachine Swagger API
// @version 2.0
// @description Swagger API for Golang Project CoffeeMachine.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @query.collection.format multi
// @license.name Apache 2.0
// @license.url https://github.com/kOstovic/coffeemachine/blob/master/LICENSE

// @BasePath /coffeemachine
// Package classification of Product API.
func main() {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	coffeemachine := router.Group("/coffeemachine")
	controllers.RegisterRoutesCoffeeMachine(coffeemachine.Group("/"))
	controllers.RegisterRoutesDrink(coffeemachine.Group("/drinks"))
	controllers.RegisterRoutesIngredients(coffeemachine.Group("/ingredients"))
	controllers.RegisterRoutesDenomination(coffeemachine.Group("/money"))
	router.GET("/coffeemachine/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Ping test
	coffeemachine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	//router := controllers.SetupRouter()
	// Listen and Server in 0.0.0.0:3000
	router.Run(":3000")
}