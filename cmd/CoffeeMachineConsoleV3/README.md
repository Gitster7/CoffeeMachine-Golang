# CoffeeMachine
CoffeeMachine application implementation in GoLang used for learning GoLang using abiosoft/ishell for interactive shell.

## Running CoffeeMachineConsoleV3

Run command in cmd/CoffeeMachineConsoleV3
```
go build -o coffeeMachine.exe main.go
```
and run coffeeMachine.exe

APP is running in local a console, commands are shown in console
and example jsons for commands can be taken from postman collection

## Commands
CoffeeMachineConsoleV3 have:
- subcommands instead of new command
- help for each command and subcommand.
- autocomplete function
- choosing of drinks by multiple choice 

CoffeeMachineConsoleV3 has commands and subcommands:
>
> - for initializing:
> > - coffeemachine
	  > >  - initializeCoffeeMachine

> - for checking status of ingredient and updating Ingredient model:
> > - ingredient
	  > >  - getAllIngredients
> >  - getIngredientsByName
> >  - putIngredients
> >  - putIngredientsByName
> >  - patchIngredients

> - for checking status of money and updating money based on Denomination model:
> > - money
	  > >  - getAllAvailableDenomination
> >  - getDenominationByName
> >  - putDenomination
> >  - putDenominationByName
> >  - patchDenomination

> - for checking all available drinks, adding them and consuming them:
> > - drinks
	  > >  - getAllAvailableDrinks
> >  - getConsumeDrink
> >  - postAddDrink
> >  - postRemoveDrink
> 
> > - other commands:
> >  - help
> >  - clear
> >  - exit
