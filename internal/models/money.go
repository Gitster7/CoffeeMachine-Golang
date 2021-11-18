package models

import (
	"fmt"
	"reflect"
)

type Denomination struct {
	Half  int
	One   int
	Two   int
	Five  int
	Ten   int
	Total float64
}

var (
	money *Denomination = new(Denomination)
)

func GetCurrentMoney() *Denomination {
	return money
}

func InitializeDenominations(d Denomination) (Denomination, error) {
	if d.Half < 0 || d.One < 0 || d.Two < 0 ||
		d.Five < 0 || d.Ten < 0 {
		return Denomination{}, fmt.Errorf("Initializing CoffeMachine must have money to work %v", d)
	}

	d.Total = float64(d.Half)*0.5 + float64(d.One) + float64(d.Two)*2 + float64(d.Five)*5 + float64(d.Ten)*10
	money = &d
	return *money, nil
}

func GetDenominationValueByName(denomination string) (string, error) {
	r := reflect.ValueOf(*money)
	for i := 0; i < r.NumField(); i++ {
		if denomination == r.Type().Field(i).Name {
			return fmt.Sprintf("Field: %s Value: %v", r.Type().Field(i).Name, r.Field(i).Interface()), nil
		}
	}
	return "", fmt.Errorf("Denomination with name '%s' not found", denomination)
}

func UpdateDenominationPut(d Denomination) (Denomination, error) {
	if d.Half < 0 || d.One < 0 || d.Two < 0 ||
		d.Five < 0 || d.Ten < 0 {
		return Denomination{}, fmt.Errorf("Values in Denomination cannot be negative'%v'", d)
	}
	money.Half = d.Half
	money.One = d.One
	money.Two = d.Two
	money.Five = d.Five
	money.Ten = d.Ten
	money.Total = float64(d.Half)*0.5 + float64(d.One) + float64(d.Two)*2 + float64(d.Five)*5 + float64(d.Ten)*10

	return *money, nil
}

func UpdateDenominationValueByName(denomination string, value int) (Denomination, error) {
	if value < 0 {
		return Denomination{}, fmt.Errorf("Value cannot be negative'%v'", value)
	}
	switch denomination {
	case "Half":
		money.Half = value
	case "One":
		money.One = value
	case "Two":
		money.Two = value
	case "Five":
		money.Five = value
	case "Ten":
		money.Ten = value
	default:
		return Denomination{}, fmt.Errorf("Denomination with name '%s' not found", denomination)
	}
	money.Total = float64(money.Half)*0.5 + float64(money.One) + float64(money.Two)*2 + float64(money.Five)*5 + float64(money.Ten)*10
	return *money, nil
}

func UpdateDenominationPatch(d Denomination) (Denomination, error) {
	if d.Half > 0 {
		money.Half = d.Half
	}
	if d.One > 0 {
		money.One = d.One
	}
	if d.Two > 0 {
		money.Two = d.Two
	}
	if d.Five > 0 {
		money.Five = d.Five
	}
	if d.Ten > 0 {
		money.Ten = d.Ten
	}
	money.Total = float64(money.Half)*0.5 + float64(money.One) + float64(money.Two)*2 + float64(money.Five)*5 + float64(money.Ten)*10
	return *money, nil
}

func UpdateDenominationConsume(d Denomination, cost float64) (Denomination, error) {
	prereq, denCal, denRet, err := CheckPrereqForMoney(d, cost)
	if prereq == false {
		return d, err
	}
	if denCal.Half != 0 {
		money.Half += denCal.Half
	}
	if denCal.One != 0 {
		money.One += denCal.One
	}
	if denCal.Two != 0 {
		money.Two += denCal.Two
	}
	if denCal.Five != 0 {
		money.Five += denCal.Five
	}
	if denCal.Ten != 0 {
		money.Ten += denCal.Ten
	}
	money.Total = float64(money.Half)*0.5 + float64(money.One) + float64(money.Two)*2 + float64(money.Five)*5 + float64(money.Ten)*10
	return denRet, nil
}

func CheckPrereqForMoney(d Denomination, cost float64) (bool, Denomination, Denomination, error) {
	d.Total = float64(d.Half)*0.5 + float64(d.One) + float64(d.Two)*2 + float64(d.Five)*5 + float64(d.Ten)*10
	if d.Total == cost {
		return true, d, Denomination{}, nil
	} else if d.Total < cost {
		return false, Denomination{}, d, fmt.Errorf("Drink cost %v. Not enough money", cost)
	} else {
		//fmt.Printf("%#v", d.Total-cost)
		prereq, aDen := cashierAlgorithmCheck(d.Total-cost, d)
		if prereq {
			return true, Denomination{d.Half - aDen.Half, d.One - aDen.One, d.Two - aDen.Two, d.Five - aDen.Five, d.Ten - aDen.Ten, 0}, aDen, nil
		} else {
			return false, Denomination{}, d, fmt.Errorf("Not enough coins to return")
		}
	}
}

func cashierAlgorithmCheck(returnMoney float64, newDen Denomination) (bool, Denomination) {
	deno := [5]float64{0.5, 1, 2, 5, 10}
	ans := [5]int{0, 0, 0, 0, 0}
	val := [5]int{money.Half + newDen.Half, money.One + newDen.One, money.Two + newDen.Two, money.Five + newDen.Five, money.Ten + newDen.Ten}

	numOfDen := (len(deno) - 1)
	retMoney := returnMoney
	for numOfDen >= 0 {
		for retMoney >= deno[numOfDen] {
			if val[numOfDen] > 0 {
				retMoney -= deno[numOfDen]
				val[numOfDen]--
				ans[numOfDen]++
			} else {
				break
			}
		}
		numOfDen--
	}
	if retMoney > 0 {
		//fmt.Printf("%#v", retMoney)
		return false, Denomination{}
	} else {
		return true, Denomination{ans[0], ans[1], ans[2], ans[3], ans[4], 0}
	}
}
