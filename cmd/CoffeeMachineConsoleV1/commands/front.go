package commands

import (
	"bufio"
	"fmt"
	"os"
)

type command interface {
	ServeCommand()
}

var (
	cmCommand, iCommand, dCommand, mCommand command
)

func RegisterCommands() {
	cmCommand = newCoffeeMachineCommand()
	iCommand = newIngredientsCommand()
	dCommand = newDrinkCommand()
	mCommand = newMoneyCommand()

	scanner()
}

func scanner() {
	fmt.Println("Input keyword of command to use that command. Input end to go back.")
	fmt.Println("Commands are: drink, ingredient, money, coffeeMachine, end.")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		switch scanner.Text() {
		case "drink":
			fmt.Println("drink command. Entering drink subroutine.")
			drink()
			break
		case "ingredient":
			fmt.Println("ingredient command. Entering ingredient subroutine.")
			ingredient()
			break
		case "money":
			fmt.Println("money command. Entering money subroutine.")
			money()
			break
		case "coffeeMachine":
			fmt.Println("coffeeMachine command. Entering coffeeMachine subroutine.")
			coffeeMachine()
			break
		case "end":
			fmt.Println("end command. Exiting program.")
			end(0)
			break
		default:
			println("Invalid command. Commands are: drink, ingredient, money, coffeeMachine, end")
		}
		fmt.Println("Commands are: drink, ingredient, money, coffeeMachine, end.")

	}
}

func coffeeMachine() {
	cmCommand.ServeCommand()
}
func drink() {
	dCommand.ServeCommand()
}
func ingredient() {
	iCommand.ServeCommand()
}
func money() {
	mCommand.ServeCommand()
}

func end(exitCode int) {
	os.Exit(exitCode)
}