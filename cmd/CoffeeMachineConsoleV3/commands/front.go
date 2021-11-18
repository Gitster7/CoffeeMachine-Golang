package commands

import (
	"github.com/abiosoft/ishell"
	"os"
)

type command interface {
	addCommands(ishell.Shell)
}
var (
	cmCommand, iCommand, dCommand, mCommand command
)
var shell = ishell.New()

func RegisterShell(){
	cmCommand = newCoffeeMachineCommand()
	iCommand = newIngredientsCommand()
	dCommand = newDrinkCommand()
	mCommand = newMoneyCommand()

	coffeeMachine()
	drink()
	ingredient()
	money()

	// display info.
	shell.Println("Sample Interactive Shell for Coffee Machine using abiosoft/ishell. Type help to check all the commands and help <command> for subcommands. You can autocomplete commands too.")
	if len(os.Args) > 1 && os.Args[1] == "exit" {
		shell.Process(os.Args[2:]...)
	} else {
		// start shell
		shell.Run()
		// teardown
		shell.Close()
	}
}
func coffeeMachine() {
	cmCommand.addCommands(*shell)
}
func drink() {
	dCommand.addCommands(*shell)
}
func ingredient() {
	iCommand.addCommands(*shell)
}
func money() {
	mCommand.addCommands(*shell)
}