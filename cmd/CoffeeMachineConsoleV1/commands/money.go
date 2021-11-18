package commands

import (
	"bufio"
	"encoding/json"
	"fmt"
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

func (mmCommand moneyCommand) ServeCommand() {
	fmt.Println("money command. available commands are: getAllAvailableDenomination, getDenominationByName, putDenomination, putDenominationByName, patchDenomination, end")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		switch scanner.Text() {
		case "getAllAvailableDenomination":
			fmt.Println("Entering getAllAvailableDenomination subroutine.")
			getAllAvailableDenomination()
			break
		case "getDenominationByName":
			fmt.Println("Entering getDenominationByName subroutine.")
			getDenominationByName()
			break
		case "putDenomination":
			fmt.Println("Entering putDenomination subroutine.")
			putDenomination()
			break
		case "putDenominationByName":
			fmt.Println("Entering putDenominationByName subroutine.")
			putDenominationByName()
			break
		case "patchDenomination":
			fmt.Println("Entering patchDenomination subroutine.")
			patchDenomination()
			break
		case "end":
			fmt.Println("end command. Exiting subroutine.")
			return
			break
		default:
			println("Invalid command. Commands are: getAllAvailableDenomination, getDenominationByName, putDenomination, putDenominationByName, patchDenomination, end")
		}
		fmt.Println("drink command. available commands are: getAllAvailableDenomination, getDenominationByName, putDenomination, putDenominationByName, patchDenomination, end")
	}
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