package models

import (
	"fmt"
	"strings"
	"testing"
)

var (
	testnameMoney          string
	moneyTestUpdateCleanup Denomination = Denomination{
		Half: 5,
		One:  5,
		Two:  5,
		Five: 5,
		Ten:  5,
	}
	moneyTestOK Denomination = Denomination{
		Half: 1,
		One:  1,
		Two:  1,
		Five: 1,
		Ten:  1,
	}
	moneyTestFail Denomination = Denomination{
		Half: -1,
		One:  1,
		Two:  1,
		Five: 1,
		Ten:  1,
	}
	moneyTestUpdate Denomination = Denomination{
		Half: 2,
		One:  2,
		Two:  2,
		Five: 2,
		Ten:  2,
	}
)

func TestInitializeDenominations(t *testing.T) {
	testnameMoney = fmt.Sprintf("%s", "InitializemoneyTestFail")
	t.Run(testnameMoney, func(t *testing.T) {
		result, err := InitializeDenominations(moneyTestFail)
		if err == nil {
			t.Errorf("Wrong value in denomination sent, test should fail, result = %v", result)
		}
	})

	testnameMoney = fmt.Sprintf("%s", "InitializemoneyTestOK")
	t.Run(testnameMoney, func(t *testing.T) {
		result, _ := InitializeDenominations(moneyTestOK)
		if result.Total != 18.5 {
			t.Errorf("Result returned from total should be 18.5, got:  result = %v", result)
		}
	})
}

func TestGetDenominationValueByName(t *testing.T) {

	testnameMoney = fmt.Sprintf("%s", "Half")
	t.Run(testnameMoney, func(t *testing.T) {
		result, _ := GetDenominationValueByName("Half")
		if !strings.Contains(result, "Half Value: 1") {
			t.Errorf("result = %v", result)
		}
	})

	testnameMoney = fmt.Sprintf("%s", "One")
	t.Run(testnameMoney, func(t *testing.T) {
		result, _ := GetDenominationValueByName("One")
		if !strings.Contains(result, "One Value: 1") {
			t.Errorf("result = %v", result)
		}
	})

	testnameMoney = fmt.Sprintf("%s", "Two")
	t.Run(testnameMoney, func(t *testing.T) {
		result, _ := GetDenominationValueByName("Two")
		if !strings.Contains(result, "Two Value: 1") {
			t.Errorf("result = %v", result)
		}
	})

	testnameMoney = fmt.Sprintf("%s", "Five")
	t.Run(testnameMoney, func(t *testing.T) {
		result, _ := GetDenominationValueByName("Five")
		if !strings.Contains(result, "Five Value: 1") {
			t.Errorf("result = %v", result)
		}
	})

	testnameMoney = fmt.Sprintf("%s", "Ten")
	t.Run(testnameMoney, func(t *testing.T) {
		result, _ := GetDenominationValueByName("Ten")
		if !strings.Contains(result, "Ten Value: 1") {
			t.Errorf("result = %v", result)
		}
	})
}

func TestUpdateDenominationPut(t *testing.T) {
	testnameMoney = fmt.Sprintf("%s", "UpdateMoneyTestFail")
	t.Run(testnameMoney, func(t *testing.T) {
		result, err := UpdateDenominationPut(moneyTestFail)
		if err == nil {
			t.Errorf("Wrong value in denomination sent, test should fail, result = %v", result)
		}
	})

	testnameMoney = fmt.Sprintf("%s", "UpdateMoneyTestUpdateOK")
	t.Run(testnameMoney, func(t *testing.T) {
		result, _ := UpdateDenominationPut(moneyTestUpdate)
		if result.Total != 37 {
			t.Errorf("Result returned from total should be 37, got: result = %v", result)
		}
	})
}

func TestUpdateDenominationValueByName(t *testing.T) {

	var tests = []struct {
		denName string
		value   int
		want    float64
		err     error
	}{
		{"Half", 3, 37.5, nil},
		{"One", 3, 38.5, nil},
		{"Two", 3, 40.5, nil},
		{"Five", 3, 45.5, nil},
		{"Ten", 3, 55.5, nil},
	}
	testnameMoney = fmt.Sprintf("%s,%d", "UnknownDenomination", 3)
	t.Run(testnameMoney, func(t *testing.T) {
		resultUnknown, errUnknown := UpdateDenominationValueByName("UnknownDenomination", 3)
		if errUnknown == nil {
			t.Errorf("There should be an error because UnknownDenomination doesn't exist, got: result = %v", resultUnknown)
		}
	})
	testnameMoney = fmt.Sprintf("%s,%d", "Half", -1)
	t.Run(testnameMoney, func(t *testing.T) {
		resultNegative, errNegative := UpdateDenominationValueByName("Half", -1)
		if errNegative == nil {
			t.Errorf("There should be an error because denomination value cannot be negative, got: result = %v", resultNegative)
		}
	})
	for _, tt := range tests {
		testnameMoney = fmt.Sprintf("UpdateDenominationValueByName %s,%d", tt.denName, tt.value)
		t.Run(testnameMoney, func(t *testing.T) {
			ans, _ := UpdateDenominationValueByName(tt.denName, tt.value)
			if ans.Total != tt.want {
				t.Errorf("UpdateDenominationValueByName falied, got %f, want %f", ans.Total, tt.want)
			}
		})
	}
	t.Cleanup(func() {
		_, _ = UpdateDenominationPut(moneyTestUpdateCleanup)
	})
}

func TestUpdateDenominationPatch(t *testing.T) {
	moneyTestPatch := Denomination{
		Half: 6,
		One:  0,
		Two:  0,
		Five: 0,
		Ten:  0,
	}
	testnameMoney = fmt.Sprintf("%s", "PatchHalf")
	t.Run(testnameMoney, func(t *testing.T) {
		moneyTestPatch.Half = 6
		result, _ := UpdateDenominationPatch(moneyTestPatch)
		if result.Half != 6 && result.One != 5 {
			t.Errorf("There should be 6 denominations of Half and old 5 denominations of One, got: result = %v", result)
		}
	})

	testnameMoney = fmt.Sprintf("%s", "PatchOneAndTwo")
	t.Run(testnameMoney, func(t *testing.T) {
		moneyTestPatch.Half = -1
		moneyTestPatch.One = 6
		moneyTestPatch.Two = 6
		result, _ := UpdateDenominationPatch(moneyTestPatch)
		if result.One != 6 && result.Two != 6 && result.Half != 5 {
			t.Errorf("There should be 6 denominations of One and 6 denominations of Two and old 5 denominations of Half, got: result = %v", result)
		}
	})

	testnameMoney = fmt.Sprintf("%s", "PatchFiveAndTen")
	t.Run(testnameMoney, func(t *testing.T) {
		moneyTestPatch.Half = -1
		moneyTestPatch.One = 0
		moneyTestPatch.Two = 0
		moneyTestPatch.Five = 6
		moneyTestPatch.Ten = 6
		result, _ := UpdateDenominationPatch(moneyTestPatch)
		if result.Five != 6 && result.Ten != 6 && result.Half != 5 {
			t.Errorf("There should be 6 denominations of Five and 6 denominations of Ten and old 5 denominations of Half, got: result = %v", result)
		}
	})
	t.Cleanup(func() {
		_, _ = UpdateDenominationPut(moneyTestUpdateCleanup)
	})
}

func TestUpdateDenominationConsume(t *testing.T) {
	var tests = []struct {
		denom Denomination
		cost  float64
		want  Denomination
		err   error
	}{
		{Denomination{5, 0, 0, 0, 0, 0}, 2.5, Denomination{Half: 0, One: 0, Two: 0, Five: 0, Ten: 0}, nil},
		{Denomination{0, 5, 0, 0, 0, 0}, 5, Denomination{Half: 0, One: 0, Two: 0, Five: 0, Ten: 0}, nil},
		{Denomination{0, 0, 5, 0, 0, 0}, 10, Denomination{Half: 0, One: 0, Two: 0, Five: 0, Ten: 0}, nil},
		{Denomination{0, 0, 0, 5, 0, 0}, 25, Denomination{Half: 0, One: 0, Two: 0, Five: 0, Ten: 0}, nil},
		{Denomination{0, 0, 0, 0, 5, 0}, 50, Denomination{Half: 0, One: 0, Two: 0, Five: 0, Ten: 0}, nil},
		{Denomination{6, 0, 0, 0, 0, 0}, 2.5, Denomination{Half: 1, One: 0, Two: 0, Five: 0, Ten: 0}, nil},
		{Denomination{6, 1, 0, 0, 0, 0}, 2.5, Denomination{Half: 1, One: 1, Two: 0, Five: 0, Ten: 0}, nil},
		{Denomination{6, 1, 2, 0, 0, 0}, 2.5, Denomination{Half: 1, One: 0, Two: 0, Five: 1, Ten: 0}, nil},
		{Denomination{0, 1, 1, 0, 0, 0}, 3, Denomination{Half: 0, One: 0, Two: 0, Five: 0, Ten: 0}, nil},
		{Denomination{0, 1, 6, 0, 0, 0}, 3, Denomination{Half: 0, One: 0, Two: 0, Five: 0, Ten: 1}, nil},
		{Denomination{0, 0, 0, 0, 1, 0}, 5, Denomination{Half: 0, One: 0, Two: 0, Five: 1, Ten: 0}, nil},
		{Denomination{1, 0, 0, 0, 500, 0}, 500, Denomination{Half: 1, One: 0, Two: 0, Five: 0, Ten: 450}, nil},
	}

	for _, tt := range tests {
		testnameMoney = fmt.Sprintf("denominations %v,cost %f", tt.denom, tt.cost)
		t.Run(testnameMoney, func(t *testing.T) {
			ans, _ := UpdateDenominationConsume(tt.denom, tt.cost)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}

	var testsError = []struct {
		denom Denomination
		cost  float64
		want  Denomination
		err   error
	}{
		{Denomination{1, 0, 0, 0, 0, 0}, 5, Denomination{Half: 1, One: 0, Two: 0, Five: 0, Ten: 0}, fmt.Errorf("Drink cost %v. Not enough money", 5)},
		{Denomination{0, 0, 0, 0, 1, 0}, 0.5, Denomination{Half: 0, One: 0, Two: 0, Five: 0, Ten: 1}, fmt.Errorf("Not enough coins to return")},
	}
	for _, tt := range testsError {
		testnameMoney = fmt.Sprintf("denominations %v,cost %f", tt.denom, tt.cost)
		t.Run(testnameMoney, func(t *testing.T) {
			_, _ = UpdateDenominationValueByName("Half", 0)
			ans, err := UpdateDenominationConsume(tt.denom, tt.cost)
			if ans != tt.want || err.Error() != tt.err.Error() {
				t.Errorf("got %v, want %v, got error %v, want error %v", ans, tt.want, err, tt.err)
			}
		})
	}
	t.Cleanup(func() {
		_, _ = UpdateDenominationPut(moneyTestUpdateCleanup)
	})
}
