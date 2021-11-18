package models

import (
	"fmt"
	"reflect"
)

type Ingredient struct {
	Water       uint16
	Milk        uint16
	Sugar       uint16
	CoffeeBeans uint16
	TeaBeans    uint16
	Cups        uint16
}

var (
	machineIngredients *Ingredient = new(Ingredient)
)

func GetMachineIngredients() *Ingredient {
	return machineIngredients
}

func InitializeIngredients(ing Ingredient) (Ingredient, error) {
	if ing.Water <= 0 && ing.Milk <= 0 && ing.Sugar <= 0 &&
		ing.CoffeeBeans <= 0 && ing.TeaBeans <= 0 && ing.Cups <= 0 {
		return Ingredient{}, fmt.Errorf("Initializing CoffeMachine must have some Ingredients to work %v", ing)
	}
	machineIngredients = &ing
	return *machineIngredients, nil
}

func GetIngredienteValueByName(ingredient string) (string, error) {
	r := reflect.ValueOf(*machineIngredients)
	for i := 0; i < r.NumField(); i++ {
		if ingredient == r.Type().Field(i).Name {
			return fmt.Sprintf("Field: %s Value: %v", r.Type().Field(i).Name, r.Field(i).Interface()), nil
		}
	}
	return "", fmt.Errorf("Ingredient with name '%s' not found", ingredient)
}

func UpdateIngredientPatch(ing Ingredient) (Ingredient, error) {
	if ing.Water > 0 {
		machineIngredients.Water = ing.Water
	}
	if ing.Milk > 0 {
		machineIngredients.Milk = ing.Milk
	}
	if ing.Sugar > 0 {
		machineIngredients.Sugar = ing.Sugar
	}
	if ing.CoffeeBeans > 0 {
		machineIngredients.CoffeeBeans = ing.CoffeeBeans
	}
	if ing.TeaBeans > 0 {
		machineIngredients.TeaBeans = ing.TeaBeans
	}
	if ing.Cups > 0 {
		machineIngredients.Cups = ing.Cups
	}
	return *machineIngredients, nil
}

func UpdateIngredientPut(ing Ingredient) (Ingredient, error) {
	machineIngredients.Water = ing.Water
	machineIngredients.Milk = ing.Milk
	machineIngredients.Sugar = ing.Sugar
	machineIngredients.CoffeeBeans = ing.CoffeeBeans
	machineIngredients.TeaBeans = ing.TeaBeans
	machineIngredients.Cups = ing.Cups

	return *machineIngredients, nil
}

func UpdateIngredientValueByName(ingredient string, value uint16) (Ingredient, error) {
	switch ingredient {
	case "Water":
		machineIngredients.Water = value
	case "Milk":
		machineIngredients.Milk = value
	case "Sugar":
		machineIngredients.Sugar = value
	case "CoffeeBeans":
		machineIngredients.CoffeeBeans = value
	case "TeaBeans":
		machineIngredients.TeaBeans = value
	case "Cups":
		machineIngredients.Cups = value
	default:
		return Ingredient{}, fmt.Errorf("Ingredient with name '%s' not found", ingredient)
	}
	return *machineIngredients, nil
}
