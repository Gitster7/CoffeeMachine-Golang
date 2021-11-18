package commands

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/abiosoft/ishell"
	"github.com/kOstovic/CoffeeMachine/internal/models"
	"os"
)

type drinkCommand struct {
	Name string
	Drink models.Drink
}
type consumeDrinkStruct struct {
	Denomination models.Denomination
}


func newDrinkCommand() *drinkCommand {
	return &drinkCommand{}
}
func (idCommand drinkCommand) addCommands(shell ishell.Shell)() {

	drink := &ishell.Cmd{
		Name: "drink",
		Func: func(c *ishell.Context) {
			c.ShowPrompt(false)
			defer c.ShowPrompt(true)
			initializeCoffeeMachine()
		},
		Help: "drink subcommand, available commands are: getAllAvailableDrinks, getConsumeDrink, postAddDrink, postRemoveDrink",
	}

	drink.AddCmd(&ishell.Cmd{
		Name: "getAllAvailableDrinks",
		Aliases: []string{"getAllDrinks", "drinks"},
		Func: func(c *ishell.Context) {
			c.ShowPrompt(false)
			defer c.ShowPrompt(true)
			getAllAvailableDrinks()
		},
		Help: "Get all currently available drinks",
	})
	drink.AddCmd(&ishell.Cmd{
		Name: "getConsumeDrink",
		Aliases: []string{"consumeDrink", "consume"},
		Func: func(c *ishell.Context) {
			var allDrinksNames = models.GetAvailableDrinksName()
			if len(allDrinksNames) != 0 {
				choice := c.MultiChoice(allDrinksNames, "Choose your drink: ")
				fmt.Printf("Drink choosen is: '%v'\n", allDrinksNames[choice])
				getConsumeDrink(allDrinksNames[choice])
			} else {
				fmt.Printf("There are no drinks to choose, Initialize them first.\n")
			}
		},
		Help: "Consume drink",
	})
	drink.AddCmd(&ishell.Cmd{
		Name: "postAddDrink",
		Aliases: []string{"addDrink"},
		Func: func(c *ishell.Context) {
			c.ShowPrompt(false)
			defer c.ShowPrompt(true)
			postAddDrink()
		},
		Help: "Add new drink",
	})
	drink.AddCmd(&ishell.Cmd{
		Name: "postRemoveDrink",
		Aliases: []string{"removeDrink"},
		Func: func(c *ishell.Context) {
			c.ShowPrompt(false)
			defer c.ShowPrompt(true)
			postRemoveDrink()
		},
		Help: "Remove drink from list",
	})
	shell.AddCmd(drink)
}

func getAllAvailableDrinks(){
	fmt.Printf("Drinks in machine are: '%v'\n", models.GetAvailableDrinks())
}

func getConsumeDrink(name string){
	fmt.Printf("Put your Denominations as json eg: {\"denomination\": {\"Half\":15, \"One\":15}}\n")
	var consumeStruct consumeDrinkStruct

	err := json.NewDecoder(os.Stdin).Decode(&consumeStruct)
	if err != nil {
		fmt.Printf("There was an error: '%v'\n", err.Error())
	}

	prereq, cost, err := models.CheckPrereqForDrink(name)
	if !prereq || err != nil {
		fmt.Printf("There was an error: '%v'\n", err.Error())
		return
	}

	denRet, err := models.UpdateDenominationConsume(consumeStruct.Denomination, cost)
	if err != nil {
		fmt.Printf("Denomination returned: '%v',There was an error: '%v'\n", denRet, err.Error())
		return
	}
	models.ConsumeDrink(name)
	fmt.Printf("Drink served: '%v' and Denomination returned: '%v'\n", name, denRet)

}
func postAddDrink(){
	fmt.Println("Enter json of name and drink eg: {\"name\":\"tea\", \"Drink\": {\"Water\":10,\"Milk\":2,\"Sugar\":4,\"CoffeeBeans\":0,\"TeaBeans\":5,\"Cups\":1,\"Money\":4}}. Ingredients requirements of drink that are not in JSON will be 0.")
	var drinkStruct drinkCommand

	err := json.NewDecoder(os.Stdin).Decode(&drinkStruct)
	if err != nil {
		fmt.Printf("There was an error: '%v'\n", err.Error())
	}
	drink, err := models.AddDrink(drinkStruct.Name, drinkStruct.Drink)
	if err != nil {
		fmt.Printf("There was an error: '%v'\n", err.Error())
		return
	}
	fmt.Printf("Drink Added in machine: '%v'\n", drink)
}
func postRemoveDrink(){
	fmt.Println("Enter Drink name(case sensitive) you wish to remove:")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	_, err := models.RemoveDrink(scanner.Text())
	if err != nil {
		fmt.Printf("There was an error: '%v'\n", err.Error())
		return
	}
	fmt.Printf("Drink '%v' is removed from machine\n", scanner.Text())
}