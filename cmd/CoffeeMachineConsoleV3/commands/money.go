package commands

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/abiosoft/ishell"
	"github.com/kOstovic/CoffeeMachine/internal/models"
	"os"
)

type moneyCommand struct {
	Denomination models.Denomination
}
type moneyByName struct {
	Name string
	Amount int
}

func newMoneyCommand() *moneyCommand {
	return &moneyCommand{}
}

func (mmCommand moneyCommand) addCommands(shell ishell.Shell) {
	money := &ishell.Cmd{
		Name: "money",
		Func: func(c *ishell.Context) {
			c.ShowPrompt(false)
			defer c.ShowPrompt(true)
			initializeCoffeeMachine()
		},
		Help: "money subcommand, available commands are: getAllAvailableDenomination, getDenominationByName, putDenomination, putDenominationByName, patchDenomination",
	}

	money.AddCmd(&ishell.Cmd{
		Name: "getAllAvailableDenomination",
		Aliases: []string{"getDenominations", "denominations"},
		Func: func(c *ishell.Context) {
			c.ShowPrompt(false)
			defer c.ShowPrompt(true)
			getAllAvailableDenomination()
		},
		Help: "Get current status of all denominations",
	})
	money.AddCmd(&ishell.Cmd{
		Name: "getDenominationByName",
		Aliases: []string{"getDenomination"},
		Func: func(c *ishell.Context) {
			c.ShowPrompt(false)
			defer c.ShowPrompt(true)
			getDenominationByName()
		},
		Help: "Get Denomination by name",
	})
	money.AddCmd(&ishell.Cmd{
		Name: "putDenomination",
		Func: func(c *ishell.Context) {
			c.ShowPrompt(false)
			defer c.ShowPrompt(true)
			putDenomination()
		},
		Help: "Add new set of Denominations",
	})
	money.AddCmd(&ishell.Cmd{
		Name: "putDenominationByName",
		Func: func(c *ishell.Context) {
			c.ShowPrompt(false)
			defer c.ShowPrompt(true)
			putDenominationByName()
		},
		Help: "Update only one Denomination",
	})
	money.AddCmd(&ishell.Cmd{
		Name: "patchDenomination",
		Func: func(c *ishell.Context) {
			c.ShowPrompt(false)
			defer c.ShowPrompt(true)
			patchDenomination()
		},
		Help: "Update some Denominations, missing ones will be left the same",
	})
	shell.AddCmd(money)
}

func getAllAvailableDenomination(){
	fmt.Printf("Current Denominations in machine are: '%v'\n", *models.GetCurrentMoney())
}

func getDenominationByName(){
	fmt.Println("Enter Ingredient name: Half, One, Two, Five, Ten, Total")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	dm, err := models.GetDenominationValueByName(scanner.Text())
	if err != nil {
		fmt.Printf("There was an error: '%v'\n", err.Error())
		return
	}
	fmt.Printf("Denomination requested in machine is: '%v'\n", dm)
}
func putDenomination(){
	fmt.Println("Enter json of denomination eg: { \"denomination\":{\"Half\":10,...}}. Denomination that are not in JSON will be 0.")
	var denStruct moneyCommand

	err := json.NewDecoder(os.Stdin).Decode(&denStruct)
	if err != nil {
		fmt.Printf("There was an error: '%v'\n", err.Error())
	}
	dm, err := models.UpdateDenominationPut(denStruct.Denomination)
	if err != nil {
		fmt.Printf("There was an error: '%v'\n", err.Error())
		return
	}
	fmt.Printf("Denomination in machine now are: '%v'\n", dm)
}
func putDenominationByName(){
	fmt.Println("Enter name(case sensitive), amount to update Denomination by name eg: {\"name\":\"Half\", \"amount\":244}")
	var denStruct moneyByName

	err := json.NewDecoder(os.Stdin).Decode(&denStruct)
	if err != nil {
		fmt.Printf("There was an error: '%v'\n", err.Error())
	}
	dm, err := models.UpdateDenominationValueByName(denStruct.Name, denStruct.Amount)
	if err != nil {
		fmt.Printf("There was an error: '%v'\n", err.Error())
		return
	}
	fmt.Printf("Denomination in machine now are: '%v'\n", dm)
}
func patchDenomination(){
	fmt.Println("Enter json of denomination eg: { \"denomination\":{\"Half\":10,...}}. Denomination that are not in JSON will not be updated.")
	var denStruct moneyCommand

	err := json.NewDecoder(os.Stdin).Decode(&denStruct)
	if err != nil {
		fmt.Printf("There was an error: '%v'\n", err.Error())
	}
	dm, err := models.UpdateDenominationPatch(denStruct.Denomination)
	if err != nil {
		fmt.Printf("There was an error: '%v'\n", err.Error())
		return
	}
	fmt.Printf("Denomination in machine now are: '%v'\n", dm)
}