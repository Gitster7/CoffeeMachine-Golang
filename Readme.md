# CoffeeMachine
CoffeeMachine Rest API and Console APP implementation in GoLang used for learning GoLang.

## Endpoints and Commands
#### Endpoints REST API version CoffeeMachine has endpoint:
>
> - for initializing
>  - /coffeemachine
> - for checking status of ingredient and updating Ingredient model
>  - /coffeemachine/ingredient
> - for checking status of money and updating money based on Denomination model    
>  - /coffeemachine/money
> - for checking all available drinks, adding them and consuming them   
>  - /coffeemachine/drinks
> -  /coffeemachine/swagger/index.html swagger endpoint in restAPIv2

#### Console APP version CoffeeMachine has command structure:

>
> - for initializing 
>  - coffeemachine subcommand
> - for checking status of ingredient and updating Ingredient model
>  - ingredient subcommand
> - for checking status of money and updating money based on Denomination model    
>  - money subcommand
> - for checking all available drinks, adding them and consuming them   
>  - drinks subcommand

## Models
```
//model used for initializing machine
cofeeMachineController struct {
	Ingredients models.Ingredient
	Money       models.Denomination
}

//model used for adding Ingredient to the machine and internally consuming it
type Ingredient struct {
	Water       int
	Milk        int
	Sugar       int
	CoffeeBeans int
	TeaBeans    int
	Cups        int
}

//model used for adding drinks
Drink struct {
	Water       int
	Milk        int
	Sugar       int
	CoffeeBeans int
	TeaBeans    int
	Cups        int
	Money       float64
}

//model used for updating money and consuming drink
Denomination struct {
	Half  int
	One   int
	Two   int
	Five  int
	Ten   int
	Total float64
}
```
## Running CoffeeMachineRestApi

Same is for CoffeeMachineRestApiV1 and CoffeeMachineRestApiV2
Run command in cmd/CoffeeMachineRestApi
```
go build -o coffeeMachine.exe main.go
```
and run coffeeMachine.exe

API is running in localhost:3000
(port will be customizable in future refactor)

Postman collection can be used for testing or just go to swagger endpoint in CoffeeMachineRestApiV2.

## Running CoffeeMachineConsole

Run command in cmd/CoffeeMachineConsoleV1 or cmd/CoffeeMachineConsoleV3
```
go build -o coffeeMachine.exe main.go
```
and run coffeeMachine.exe

APP is running in local a console, commands are shown in console (or from help command in v3)
and example jsons for commands can be taken from postman collection


## Docker

Currently, building Dockerfile from deployments/apiVersion/Dockerfile

Run BuildDocker.ps1 in api folder to automatically build docker image - if you previously set: Set-ExecutionPolicy unrestricted

Or just run BuildDocker.bat

After that you can run this docker image for example on some other port like this:
```
docker run -p 3002:3000 github.com/kostovic/coffeemachine:restapiv2.0
```
restapiv2.0 is running in production mode by default

## Docker Compose

Docker Compose file is set in deployment/DockerCompose
Edit .env file to change variables from docker-compose file and then to run in detached mode run first command and to stop second command:
```
docker-compose up -d
docker-compose down
```
More in README.md in DockerCompose folder


## author

Omkar Vatsa
