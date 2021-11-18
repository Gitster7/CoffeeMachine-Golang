package commands

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/abiosoft/ishell"
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
func (imCommand ingredientsCommand) addCommands(shell ishell.Shell) {
	ingredient := &ishell.Cmd{
		Name: "ingredient",
		Func: func(c *ishell.Context) {
			c.ShowPrompt(false)
			defer c.ShowPrompt(true)
			initializeCoffeeMachine()
		},
		Help: "ingredient subcommand, available commands are: getAllIngredients, getIngredientsByName, putIngredients, putIngredientsByName, patchIngredients",
	}

	ingredient.AddCmd(&ishell.Cmd{
		Name: "getAllIngredients",
		Aliases: []string{"getIngredients", "ingredients"},
		Func: func(c *ishell.Context) {
			c.ShowPrompt(false)
			defer c.ShowPrompt(true)
			getAllIngredients()
		},
		Help: "Get current status of all ingredients",
	})
	ingredient.AddCmd(&ishell.Cmd{
		Name: "getIngredientsByName",
		Aliases: []string{"getIngredient"},
		Func: func(c *ishell.Context) {
			c.ShowPrompt(false)
			defer c.ShowPrompt(true)
			getIngredientsByName()
		},
		Help: "Get Ingredients by name",
	})
	ingredient.AddCmd(&ishell.Cmd{
		Name: "putIngredients",
		Func: func(c *ishell.Context) {
			c.ShowPrompt(false)
			defer c.ShowPrompt(true)
			putIngredients()
		},
		Help: "Add new set of Ingredients",
	})
	ingredient.AddCmd(&ishell.Cmd{
		Name: "putIngredientsByName",
		Func: func(c *ishell.Context) {
			c.ShowPrompt(false)
			defer c.ShowPrompt(true)
			putIngredientsByName()
		},
		Help: "Update only one Ingredient",
	})
	ingredient.AddCmd(&ishell.Cmd{
		Name: "patchIngredients",
		Func: func(c *ishell.Context) {
			c.ShowPrompt(false)
			defer c.ShowPrompt(true)
			patchIngredients()
		},
		Help: "Update some Ingredients, missing ones will be left the same",
	})
	shell.AddCmd(ingredient)
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