package models

import (
	"fmt"
)

type Drink struct {
	Water       uint16
	Milk        uint16
	Sugar       uint16
	CoffeeBeans uint16
	TeaBeans    uint16
	Cups        uint16
	Money       float64
}

var (
	drinks = make(map[string]Drink)
)

func GetAvailableDrinks() map[string]Drink {
	return drinks
}

func GetAvailableDrinksName() []string {
	drinksNameArr := make([]string, 0, len(drinks))
	for key := range drinks {
		drinksNameArr = append(drinksNameArr, key)
	}
	return drinksNameArr
}

func GetDrinkByName(name string) Drink {
	return drinks[name]
}

func AddDrink(name string, drink Drink) (Drink, error) {
	if drink.Water < 0 || drink.Milk < 0 || drink.Sugar < 0 ||
		drink.CoffeeBeans < 0 || drink.TeaBeans < 0 || drink.Cups < 0 || drink.Money < 0 {
		return Drink{}, fmt.Errorf("Drink must have non negative values for ingredients and money %v", drink)
	}
	if drink.Water <= 0 && drink.Milk <= 0 && drink.Sugar <= 0 &&
		drink.CoffeeBeans <= 0 && drink.TeaBeans <= 0 && drink.Cups <= 0 && drink.Money <= 0 {
		return Drink{}, fmt.Errorf("Drink must have at least one 0 or positive value %v", drink)
	}
	_, drinkExist := drinks[name]
	if drinkExist == true {
		return drink, fmt.Errorf("Drink already exists '%v'", drinks[name])
	}
	drinks[name] = drink
	return drink, nil
}

func RemoveDrink(name string) (bool, error) {
	_, drinkExist := drinks[name]
	if drinkExist != true {
		return false, fmt.Errorf("Drink doesn't exists '%v'", drinks[name])
	} else {
		delete(drinks, name)
		return true, nil
	}
}

func ConsumeDrink(name string) (bool, error) {
	prereq, _, err := CheckPrereqForDrink(name)
	if prereq {
		machineIngredients.Water -= drinks[name].Water
		machineIngredients.Milk -= drinks[name].Milk
		machineIngredients.Sugar -= drinks[name].Sugar
		machineIngredients.CoffeeBeans -= drinks[name].CoffeeBeans
		machineIngredients.TeaBeans -= drinks[name].TeaBeans
		machineIngredients.Cups -= drinks[name].Cups
		return true, nil
	} else {
		return false, err
	}
}

func CheckPrereqForDrink(name string) (bool, float64, error) {
	_, drinkExist := drinks[name]
	if drinkExist == true {
		if (int(machineIngredients.Water)-int(drinks[name].Water)) < 0 ||
			(int(machineIngredients.Milk)-int(drinks[name].Milk)) < 0 ||
			(int(machineIngredients.Sugar)-int(drinks[name].Sugar)) < 0 ||
			(int(machineIngredients.CoffeeBeans)-int(drinks[name].CoffeeBeans)) < 0 ||
			(int(machineIngredients.TeaBeans)-int(drinks[name].TeaBeans)) < 0 ||
			(int(machineIngredients.Cups)-int(drinks[name].Cups)) < 0 {
			return false, 0, fmt.Errorf("There are not enough ingredients for drink with name '%s'", name)
		}
		return true, drinks[name].Money, nil
	} else {
		return false, 0, fmt.Errorf("Drink with name '%s' not found", name)
	}
}
