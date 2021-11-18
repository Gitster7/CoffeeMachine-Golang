package main

import (
	"fmt"
	"github.com/kOstovic/CoffeeMachine/cmd/CoffeeMachineConsoleV1/commands"
)

func main() {
	fmt.Println("Simple Coffee Machine Shell")
	fmt.Println("-------------------------------------------------------------------")
	commands.RegisterCommands()
}