# CoffeeMachine
CoffeeMachine application implementation in GoLang used for learning GoLang.


## Running CoffeeMachineConsoleV1

Run command in cmd/CoffeeMachineConsoleV1
```
go build -o coffeeMachine.exe main.go
```
and run coffeeMachine.exe

APP is running in local a console, commands are shown in console 
and example jsons for commands can be taken from postman collection

## Commands
CoffeeMachineConsoleV1 has commands and commands inside of commands:
>
> - for initializing:
> > - coffeemachine
> >  - initializeCoffeeMachine
> >  - end - to exit to main part

> - for checking status of ingredient and updating Ingredient model:
> > - ingredient
> >  - getAllIngredients
> >  - getIngredientsByName
> >  - putIngredients
> >  - putIngredientsByName
> >  - patchIngredients
> >  - end - to exit to main part

> - for checking status of money and updating money based on Denomination model:  
> > - money
> >  - getAllAvailableDenomination
> >  - getDenominationByName
> >  - putDenomination
> >  - putDenominationByName
> >  - patchDenomination
> >  - end - to exit to main part

> - for checking all available drinks, adding them and consuming them:
> > - drinks
> >  - getAllAvailableDrinks
> >  - getConsumeDrink
> >  - postAddDrink
> >  - postRemoveDrink
> >  - end - to exit to main part
