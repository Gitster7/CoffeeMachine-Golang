package commands

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/kOstovic/CoffeeMachine/internal/models"
	"os"
)

type ingredientsCommand struct {
	Ingredients models.Ingredient
}
type ingredientByName struct {
	Name string
	Amount int
}

func newIngredientsCommand() *ingredientsCommand {
	return &ingredientsCommand{}
}
func (imCommand ingredientsCommand) ServeCommand() {
	fmt.Println("ingredients command. available commands are: getAllIngredients, getIngredientsByName, putIngredients, putIngredientsByName, patchIngredients, end")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		switch scanner.Text() {
		case "getAllIngredients":
			fmt.Println("Entering getAllIngredients subroutine.")
			getAllIngredients()
			break
		case "getIngredientsByName":
			fmt.Println("Entering getIngredientsByName subroutine.")
			getIngredientsByName()
			break
		case "putIngredients":
			fmt.Println("Entering putIngredients subroutine.")
			putIngredients()
			break
		case "putIngredientsByName":
			fmt.Println("Entering putIngredientsByName subroutine.")
			putIngredientsByName()
			break
		case "patchIngredients":
			fmt.Println("Entering patchIngredients subroutine.")
			patchIngredients()
			break
		case "end":
			fmt.Println("end command. Exiting subroutine.")
			return
			break
		default:
			println("Invalid command. Commands are: getAllIngredients, getIngredientsByName, putIngredients, putIngredientsByName, patchIngredients, end")
		}
		fmt.Println("drink command. available commands are: getAllIngredients, getIngredientsByName, putIngredients, putIngredientsByName, patchIngredients, end")
	}
}

func getAllIngredients(){
	fmt.Printf("Current Ingredients in machine are: '%v'\n", *models.GetMachineIngredients())
}
func getIngredientsByName(){
	fmt.Println("Enter Ingredient name: Water, Milk, Sugar, CoffeeBeans, TeaBeans, Cups")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	cm, err := models.GetIngredienteValueByName(scanner.Text())
	if err != nil {
		fmt.Printf("There was an error: '%v'\n", err.Error())
		return
	}
	fmt.Printf("Ingredient requested in machine is: '%v'\n", cm)
}
func putIngredients(){
	fmt.Println("Enter json of ingredient eg: { \"ingredients\":{\"Water\":1000,...}}. All Ingredients that are not in JSON will be 0.")
	var ingStruct ingredientsCommand

	err := json.NewDecoder(os.Stdin).Decode(&ingStruct)
	if err != nil {
		fmt.Printf("There was an error: '%v'\n", err.Error())
	}
	cm, err := models.UpdateIngredientPut(ingStruct.Ingredients)
	if err != nil {
		fmt.Printf("There was an error: '%v'\n", err.Error())
		return
	}
	fmt.Printf("Ingredients in machine now are: '%v'\n", cm)
}
func putIngredientsByName(){
	fmt.Println("Enter json of name(case sensitive), amount to update Ingredient by name eg: {\"name\":\"Water\", \"amount\":244}")
	var ingStruct ingredientByName

	err := json.NewDecoder(os.Stdin).Decode(&ingStruct)
	if err != nil {
		fmt.Printf("There was an error: '%v'\n", err.Error())
	}
	valueuint16 := uint16(ingStruct.Amount)
	if err != nil {
		fmt.Printf("There was an error: '%v'\n", err.Error())
		return
	}
	cm, err := models.UpdateIngredientValueByName(ingStruct.Name, valueuint16)
	if err != nil {
		fmt.Printf("There was an error: '%v'\n", err.Error())
		return
	}
	fmt.Printf("Ingredients in machine now are: '%v'\n", cm)
}
func patchIngredients(){
	fmt.Println("Enter json of ingredient eg: { \"ingredients\":{\"Water\":1000,...}}. Ingredients that are not in JSON will not be updated.")
	var ingStruct ingredientsCommand

	err := json.NewDecoder(os.Stdin).Decode(&ingStruct)
	if err != nil {
		fmt.Printf("There was an error: '%v'\n", err.Error())
	}
	cm, err := models.UpdateIngredientPatch(ingStruct.Ingredients)
	if err != nil {
		fmt.Printf("There was an error: '%v'\n", err.Error())
		return
	}
	fmt.Printf("Ingredients in machine now are: '%v'\n", cm)
}