package commands

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/kOstovic/CoffeeMachine/internal/models"
	"os"
)

type drinkCommand struct {
	Name string
	Drink models.Drink
}
type consumeDrinkStruct struct {
	Name string
	Denomination models.Denomination
}


func newDrinkCommand() *drinkCommand {
return &drinkCommand{}
}
func (idCommand drinkCommand) ServeCommand() {
	fmt.Println("drink command. available commands are: getAllAvailableDrinks, getConsumeDrink, postAddDrink, postRemoveDrink, end")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		switch scanner.Text() {
		case "getAllAvailableDrinks":
			fmt.Println("Entering getAllAvailableDrinks subroutine.")
			getAllAvailableDrinks()
			break
		case "getConsumeDrink":
			fmt.Println("Entering getConsumeDrink subroutine.")
			getConsumeDrink()
			break
		case "postAddDrink":
			fmt.Println("Entering postAddDrink subroutine.")
			postAddDrink()
			break
		case "postRemoveDrink":
			fmt.Println("Entering postRemoveDrink subroutine.")
			postRemoveDrink()
			break
		case "end":
			fmt.Println("end command. Exiting subroutine.")
			return
			break
		default:
			println("Invalid command. Commands are: getAllAvailableDrinksEntering, getConsumeDrink, postAddDrink, postRemoveDrink, end")
		}
		fmt.Println("drink command. available commands are: getAllAvailableDrinks, getConsumeDrink, postAddDrink, postRemoveDrink, end")
	}
}

func getAllAvailableDrinks(){
	fmt.Printf("Drinks in machine are: '%v'\n", models.GetAvailableDrinks())
}

func getConsumeDrink(){
	fmt.Println("Enter Drink in json name(case sensitive), Denomination(json) eg: {\"name\":\"tea\", \"denomination\": {\"Half\":15, \"One\":15}}")

	var consumeStruct consumeDrinkStruct

	err := json.NewDecoder(os.Stdin).Decode(&consumeStruct)
	if err != nil {
		fmt.Printf("There was an error: '%v'\n", err.Error())
	}

	prereq, cost, err := models.CheckPrereqForDrink(consumeStruct.Name)
	if !prereq || err != nil {
		fmt.Printf("There was an error: '%v'\n", err.Error())
		return
	}

	denRet, err := models.UpdateDenominationConsume(consumeStruct.Denomination, cost)
	if err != nil {
		fmt.Printf("Denomination returned: '%v',There was an error: '%v'\n", denRet, err.Error())
		return
	}
	models.ConsumeDrink(consumeStruct.Name)
	fmt.Printf("Drink served: '%v' and Denomination returned: '%v'\n", consumeStruct.Name, denRet)

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