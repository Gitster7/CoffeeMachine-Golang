package commands

import (
	"encoding/json"
	"fmt"
	"github.com/abiosoft/ishell"
	"github.com/kOstovic/CoffeeMachine/internal/models"
	"os"
)

type coffeeMachineCommand struct {
	Ingredients models.Ingredient
	Money       models.Denomination
}

//machineInitialized is private variable used for checking whether machine has been initialized
var (
	machineInitialized bool = false
)
func newCoffeeMachineCommand() *coffeeMachineCommand {
	return &coffeeMachineCommand{}
}
func (cmCommand coffeeMachineCommand) addCommands(shell ishell.Shell){
	shell.AddCmd(&ishell.Cmd{
		Name: "coffeeMachineInit",
		Aliases: []string{"initialize", "initializeCoffeeMachine"},
		Func: func(c *ishell.Context) {
			c.ShowPrompt(false)
			defer c.ShowPrompt(true)
			initializeCoffeeMachine()
		},
		Help: "Initialize coffeeMachine",
	})
}

func initializeCoffeeMachine() {
	if machineInitialized == true {
		fmt.Printf("Ingredients in machine are: '%v', Denominations in machine are: '%v', error if existing '%v'\n", *models.GetMachineIngredients(), *models.GetCurrentMoney(), fmt.Errorf("Coffee Machine object already Initialized\n"))
		return
	}
	fmt.Println("Enter json eg: {\"ingredients\": {...}, \"money\": {...}}")
	var initStruct coffeeMachineCommand

	err := json.NewDecoder(os.Stdin).Decode(&initStruct)
	if err != nil {
		fmt.Printf("There was an error: '%v'\n", err.Error())
	}
	if initStruct.Ingredients.Water < 0 || initStruct.Ingredients.Milk < 0 || initStruct.Ingredients.Sugar < 0 ||
		initStruct.Ingredients.CoffeeBeans < 0 || initStruct.Ingredients.TeaBeans < 0 || initStruct.Ingredients.Cups < 0 ||
		initStruct.Money.Half < 0 || initStruct.Money.One < 0 || initStruct.Money.Two < 0 || initStruct.Money.Five < 0 || initStruct.Money.Ten < 0 {
		fmt.Printf("Ingredients in machine are: '%v', Denominations in machine are: '%v', error if existing '%v'\n", models.Ingredient{}, models.Denomination{}, fmt.Errorf("Values in ingredient and money cannot be negative'%v'", initStruct))
		return
	}

	cm, err := models.InitializeIngredients(initStruct.Ingredients)
	if err != nil {
		fmt.Printf("Ingredients in machine are: '%v', Denominations in machine are: '%v', error if existing '%v'\n", models.Ingredient{}, models.Denomination{}, err.Error())
		return
	}
	mm, err := models.InitializeDenominations(initStruct.Money)
	if err != nil {
		fmt.Printf("Ingredients in machine are: '%v', Denominations in machine are: '%v', error if existing '%v'\n", models.Ingredient{}, models.Denomination{}, err.Error())
		return
	}
	machineInitialized = true
	fmt.Printf("Ingredients in machine are: '%v', Denominations in machine are: '%v'\n", cm, mm)
}