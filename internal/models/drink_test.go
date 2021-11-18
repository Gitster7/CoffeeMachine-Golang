package models

import (
	"fmt"
	"testing"
)

var (
	testnameDrink string
	drinkTestOK   Drink = Drink{
		Water:       2,
		Milk:        2,
		Sugar:       2,
		CoffeeBeans: 2,
		TeaBeans:    2,
		Cups:        2,
		Money:       2,
	}
	drinkTestOKZeroMoney Drink = Drink{
		Water:       1,
		Milk:        1,
		Sugar:       1,
		CoffeeBeans: 1,
		TeaBeans:    1,
		Cups:        1,
		Money:       0,
	}
	drinkTestFailMoney Drink = Drink{
		Water:       0,
		Milk:        1,
		Sugar:       1,
		CoffeeBeans: 1,
		TeaBeans:    1,
		Cups:        1,
		Money:       -1,
	}
	drinkTestFailZero Drink = Drink{
		Water:       0,
		Milk:        0,
		Sugar:       0,
		CoffeeBeans: 0,
		TeaBeans:    0,
		Cups:        0,
		Money:       0,
	}
	ingredientTestCleanUpInDrink Ingredient = Ingredient{
		Water:       0,
		Milk:        0,
		Sugar:       0,
		CoffeeBeans: 0,
		TeaBeans:    0,
		Cups:        0,
	}
)

func TestAddDrink(t *testing.T) {

	var tests = []struct {
		name  string
		drink Drink
		err   error
	}{
		{"drinkTestOK", drinkTestOK, nil},
		{"drinkTestOKZeroMoney", drinkTestOKZeroMoney, nil},
		{"Water", Drink{Water: 1}, nil},
		{"Milk", Drink{Milk: 1}, nil},
		{"Sugar", Drink{Sugar: 1}, nil},
		{"CoffeeBeans", Drink{CoffeeBeans: 1}, nil},
		{"TeaBeans", Drink{TeaBeans: 1}, nil},
		{"Cups", Drink{Cups: 1}, nil},
	}

	for _, tt := range tests {
		testnameDrink = fmt.Sprintf("Add Drinkname %v", tt.name)
		t.Run(testnameDrink, func(t *testing.T) {
			result, err := AddDrink(tt.name, tt.drink)
			if result != tt.drink {
				t.Errorf("AddDrink got %v, want %v, err %v", result, tt.drink, err)
			}
		})
	}
	var testsError = []struct {
		name  string
		drink Drink
		want  Drink
		err   error
	}{
		{"drinkTestOK", drinkTestOK, drinkTestOK, fmt.Errorf("Drink already exists '%v'", drinkTestOK)},
		{"drinkTestFailMoney", drinkTestFailMoney, Drink{}, fmt.Errorf("Drink must have non negative values for ingredients and money %v", drinkTestFailMoney)},
		{"drinkTestFailZero", drinkTestFailZero, Drink{}, fmt.Errorf("Drink must have at least one 0 or positive value %v", Drink{})},
	}

	for _, tt := range testsError {
		testnameDrink = fmt.Sprintf("Consume Fail Drinkname %v", tt.name)
		t.Run(testnameDrink, func(t *testing.T) {
			result, err := AddDrink(tt.name, tt.drink)
			if result != tt.want || err.Error() != tt.err.Error() {
				t.Errorf("AddDrink got %v, want %v, err %v", result, tt.drink, err)
			}
		})
	}
}

func TestGetDrinkByName(t *testing.T) {
	testnameDrink = fmt.Sprintf("%s", "GetdrinkOK")
	t.Run(testnameDrink, func(t *testing.T) {
		result := GetDrinkByName("drinkTestOK")
		if result != drinkTestOK {
			t.Errorf("Drink should exist. result = %v", drinkTestOK)
		}
	})

	testnameDrink = fmt.Sprintf("%s", "GetdrinkFail")
	t.Run(testnameDrink, func(t *testing.T) {
		result := GetDrinkByName("drinkTestFailMoney")
		if result == drinkTestFailMoney {
			t.Errorf("Drink should not exist. result = %v", result)
		}
	})
}

func TestGetAvailableDrinks(t *testing.T) {
	testnameDrink = fmt.Sprintf("%s", "GetAllDrinks")
	t.Run(testnameDrink, func(t *testing.T) {
		result := GetAvailableDrinks()
		if len(result) != 8 || result["drinkTestOK"] != drinkTestOK {
			t.Errorf("Drink should exist. result = %v", result)
		}
	})
}

func TestGetAvailableDrinksName(t *testing.T) {
	testnameDrink = fmt.Sprintf("%s", "GetAllDrinkNames")
	t.Run(testnameDrink, func(t *testing.T) {
		result := GetAvailableDrinksName()
		if len(result) != 8 {
			t.Errorf("DrinkNames should exist and there should be 8 of them, result = %v", result)
		}
	})
}

func TestCheckPrereqForDrink(t *testing.T) {
	var testsError = []struct {
		name   string
		result bool
		err    error
	}{
		{"drinkTestOK", false, fmt.Errorf("There are not enough ingredients for drink with name '%v'", "drinkTestOK")},
		{"drinkFail", false, fmt.Errorf("Drink with name '%v' not found", "drinkFail")},
	}

	for _, tt := range testsError {
		testnameDrink = fmt.Sprintf("Check to consume Drinkname %v", tt.name)
		t.Run(testnameDrink, func(t *testing.T) {
			ans, _, err := CheckPrereqForDrink(tt.name)
			machineIngredients := GetMachineIngredients()
			if ans != tt.result || err.Error() != tt.err.Error() {
				t.Errorf("TestCheckPrereqForDrink got %v, want %v, err %v, status of all ingredients %v", ans, tt.result, err, machineIngredients)
			}
		})
	}

	var tests = []struct {
		name   string
		result bool
		err    error
	}{
		{"drinkTestOK", true, nil},
		{"Water", true, nil},
		{"Milk", true, nil},
		{"Sugar", true, nil},
		{"CoffeeBeans", true, nil},
		{"TeaBeans", true, nil},
		{"Cups", true, nil},
	}

	_, _ = InitializeIngredients(Ingredient{Water: 10, Milk: 10, Sugar: 10, CoffeeBeans: 10, TeaBeans: 10, Cups: 10})

	for _, tt := range tests {
		testnameDrink = fmt.Sprintf("Check to consume Drinkname %v", tt.name)
		t.Run(testnameDrink, func(t *testing.T) {
			ans, _, err := CheckPrereqForDrink(tt.name)
			machineIngredients := GetMachineIngredients()
			if ans != tt.result {
				t.Errorf("TestCheckPrereqForDrink got %v, want %v, err %v, status of all ingredients %v", ans, tt.result, err, machineIngredients)
			}
		})
	}

	t.Cleanup(func() {
		_, _ = UpdateIngredientPut(ingredientTestCleanUpInDrink)
	})
}

func TestDrinkConsume(t *testing.T) {
	var testsError = []struct {
		name   string
		result bool
		err    error
	}{
		{"drinkTestOK", false, fmt.Errorf("There are not enough ingredients for drink with name '%v'", "drinkTestOK")},
		{"drinkFail", false, fmt.Errorf("Drink with name '%v' not found", "drinkFail")},
	}

	for _, tt := range testsError {
		testnameDrink = fmt.Sprintf("Consume Drinkname %v", tt.name)
		t.Run(testnameDrink, func(t *testing.T) {
			result, err := ConsumeDrink(tt.name)
			machineIngredients := GetMachineIngredients()
			if result != tt.result || err.Error() != tt.err.Error() {
				t.Errorf("TestDrinkConsume got %v, want %v, err %v, status of all ingredients %v", result, tt.result, err, machineIngredients)
			}
		})
	}

	var tests = []struct {
		name   string
		result bool
		err    error
	}{
		{"drinkTestOK", true, nil},
		{"Water", true, nil},
		{"Milk", true, nil},
		{"Sugar", true, nil},
		{"CoffeeBeans", true, nil},
		{"TeaBeans", true, nil},
		{"Cups", true, nil},
	}

	_, _ = InitializeIngredients(Ingredient{Water: 10, Milk: 10, Sugar: 10, CoffeeBeans: 10, TeaBeans: 10, Cups: 10})

	for _, tt := range tests {
		testnameDrink = fmt.Sprintf("Consume Drinkname %v", tt.name)
		t.Run(testnameDrink, func(t *testing.T) {
			result, err := ConsumeDrink(tt.name)
			machineIngredients := GetMachineIngredients()
			if result != tt.result {
				t.Errorf("TestDrinkConsume got %v, want %v, err %v, status of all ingredients %v", result, tt.result, err, machineIngredients)
			}
		})
	}

	t.Cleanup(func() {
		_, _ = UpdateIngredientPut(ingredientTestCleanUpInDrink)
	})
}

func TestRemoveDrink(t *testing.T) {
	testnameDrink = fmt.Sprintf("%s", "RemoveDrink")
	t.Run(testnameDrink, func(t *testing.T) {
		result, err := RemoveDrink("drinkTestOK")
		if result == false {
			t.Errorf("Drink should exist and be deleted. result = %v", err)
		}
	})
}
