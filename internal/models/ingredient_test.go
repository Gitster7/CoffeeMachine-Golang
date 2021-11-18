package models

import (
	"fmt"
	"testing"
)

var (
	testnameIngredient    string
	ingredientTestCleanUp Ingredient = Ingredient{
		Water:       10,
		Milk:        10,
		Sugar:       10,
		CoffeeBeans: 10,
		TeaBeans:    10,
		Cups:        10,
	}
	ingredientTestPut Ingredient = Ingredient{
		Water:       5,
		Milk:        5,
		Sugar:       5,
		CoffeeBeans: 5,
		TeaBeans:    5,
		Cups:        5,
	}
)

func TestInitializeIngredients(t *testing.T) {
	testnameIngredient = fmt.Sprintf("%s", "InitializeIngredientsFail")
	t.Run(testnameIngredient, func(t *testing.T) {
		_, err := InitializeIngredients(Ingredient{})
		if err == nil {
			t.Errorf("Machine should not be initialized. All values of Ingredient are 0.")
		}
	})

	testnameIngredient = fmt.Sprintf("%s", "InitializeIngredientsOK")
	t.Run(testnameIngredient, func(t *testing.T) {
		result, _ := InitializeIngredients(ingredientTestCleanUp)
		if result != ingredientTestCleanUp {
			t.Errorf("Machine should have been initialized. All values of Ingredient are 0, instead result = %v", result)
		}
	})
}

func TestGetMachineIngredients(t *testing.T) {
	testnameIngredient = fmt.Sprintf("%s", "GetMachineIngredients")
	t.Run(testnameIngredient, func(t *testing.T) {
		result := GetMachineIngredients()
		if *result != ingredientTestCleanUp {
			t.Errorf("Ingredients should have initialized with ingredientTestCleanUp = %v, instead result = %v", ingredientTestCleanUp, result)
		}
	})
}
func TestGetIngredientValueByName(t *testing.T) {
	var tests = []struct {
		name         string
		ingValueWant string
		err          error
	}{
		{"Water", "Field: Water Value: 10", nil},
		{"Milk", "Field: Milk Value: 10", nil},
		{"Sugar", "Field: Sugar Value: 10", nil},
		{"CoffeeBeans", "Field: CoffeeBeans Value: 10", nil},
		{"TeaBeans", "Field: TeaBeans Value: 10", nil},
		{"Cups", "Field: Cups Value: 10", nil},
	}

	for _, tt := range tests {
		testnameIngredient = fmt.Sprintf("Ingredient %v", tt.name)
		t.Run(testnameIngredient, func(t *testing.T) {
			ans, _ := GetIngredienteValueByName(tt.name)
			if ans != tt.ingValueWant {
				t.Errorf("got %v, want %v", ans, tt.ingValueWant)
			}
		})
	}

	var testsError = []struct {
		name         string
		ingValueWant string
		err          error
	}{
		{"nonExistingIngredient", "", fmt.Errorf("Ingredient with name '%s' not found", "nonExistingIngredient")},
	}
	for _, tt := range testsError {
		testnameIngredient = fmt.Sprintf("Ingredient %v", tt.name)
		t.Run(testnameIngredient, func(t *testing.T) {
			ans, err := GetIngredienteValueByName(tt.name)
			if ans != tt.ingValueWant || err.Error() != tt.err.Error() {
				t.Errorf("got %v, want %v, got error %v, want error %v", ans, tt.ingValueWant, err, tt.err)
			}
		})
	}
}

func TestPatchIngredient(t *testing.T) {

	var tests = []struct {
		value Ingredient
		want  Ingredient
		err   error
	}{
		{Ingredient{Water: 5, Milk: 5, Sugar: 5, CoffeeBeans: 5, TeaBeans: 5, Cups: 5}, Ingredient{Water: 5, Milk: 5, Sugar: 5, CoffeeBeans: 5, TeaBeans: 5, Cups: 5}, nil},
		{Ingredient{Water: 0}, Ingredient{Water: 5, Milk: 5, Sugar: 5, CoffeeBeans: 5, TeaBeans: 5, Cups: 5}, nil},
		{Ingredient{Water: 10}, Ingredient{Water: 10, Milk: 5, Sugar: 5, CoffeeBeans: 5, TeaBeans: 5, Cups: 5}, nil},
		{Ingredient{Milk: 5}, Ingredient{Water: 10, Milk: 5, Sugar: 5, CoffeeBeans: 5, TeaBeans: 5, Cups: 5}, nil},
		{Ingredient{Milk: 6}, Ingredient{Water: 10, Milk: 6, Sugar: 5, CoffeeBeans: 5, TeaBeans: 5, Cups: 5}, nil},
		{Ingredient{Sugar: 6}, Ingredient{Water: 10, Milk: 6, Sugar: 6, CoffeeBeans: 5, TeaBeans: 5, Cups: 5}, nil},
		{Ingredient{CoffeeBeans: 6}, Ingredient{Water: 10, Milk: 6, Sugar: 6, CoffeeBeans: 6, TeaBeans: 5, Cups: 5}, nil},
		{Ingredient{TeaBeans: 6}, Ingredient{Water: 10, Milk: 6, Sugar: 6, CoffeeBeans: 6, TeaBeans: 6, Cups: 5}, nil},
		{Ingredient{Cups: 6}, Ingredient{Water: 10, Milk: 6, Sugar: 6, CoffeeBeans: 6, TeaBeans: 6, Cups: 6}, nil},
		{Ingredient{Milk: 10, Sugar: 10}, Ingredient{Water: 10, Milk: 10, Sugar: 10, CoffeeBeans: 6, TeaBeans: 6, Cups: 6}, nil},
		{Ingredient{CoffeeBeans: 10, TeaBeans: 10, Cups: 10}, Ingredient{Water: 10, Milk: 10, Sugar: 10, CoffeeBeans: 10, TeaBeans: 10, Cups: 10}, nil},
	}
	for _, tt := range tests {
		testnameIngredient = fmt.Sprintf("PatchIngredient %v", tt.value)
		t.Run(testnameIngredient, func(t *testing.T) {
			ans, _ := UpdateIngredientPatch(tt.value)
			if ans != tt.want {
				t.Errorf("PatchIngredient falied, got %v, want %v", ans, tt.want)
			}
		})
	}
	t.Cleanup(func() {
		_, _ = UpdateIngredientPut(ingredientTestCleanUp)
	})
}

func TestUpdateIngredientPut(t *testing.T) {
	testnameIngredient = fmt.Sprintf("%s", "UpdateIngredientTestOk")
	t.Run(testnameIngredient, func(t *testing.T) {
		result, _ := UpdateIngredientPut(ingredientTestPut)
		if result != ingredientTestPut {
			t.Errorf("Result returned from UpdateIngredient does not equal ingredientTestPut, got: result = %v", result)
		}
	})
}

func TestUpdateIngredientValueByName(t *testing.T) {

	var tests = []struct {
		ingName string
		value   uint16
		want    Ingredient
		err     error
	}{
		{"Water", 10, Ingredient{Water: 10, Milk: 5, Sugar: 5, CoffeeBeans: 5, TeaBeans: 5, Cups: 5}, nil},
		{"Milk", 10, Ingredient{Water: 10, Milk: 10, Sugar: 5, CoffeeBeans: 5, TeaBeans: 5, Cups: 5}, nil},
		{"Sugar", 10, Ingredient{Water: 10, Milk: 10, Sugar: 10, CoffeeBeans: 5, TeaBeans: 5, Cups: 5}, nil},
		{"CoffeeBeans", 10, Ingredient{Water: 10, Milk: 10, Sugar: 10, CoffeeBeans: 10, TeaBeans: 5, Cups: 5}, nil},
		{"TeaBeans", 10, Ingredient{Water: 10, Milk: 10, Sugar: 10, CoffeeBeans: 10, TeaBeans: 10, Cups: 5}, nil},
		{"Cups", 10, Ingredient{Water: 10, Milk: 10, Sugar: 10, CoffeeBeans: 10, TeaBeans: 10, Cups: 10}, nil},
		{"Cups", 0, Ingredient{Water: 10, Milk: 10, Sugar: 10, CoffeeBeans: 10, TeaBeans: 10, Cups: 0}, nil},
	}
	testnameIngredient = fmt.Sprintf("%s,%d", "UnknownIngredient", 10)
	t.Run(testnameIngredient, func(t *testing.T) {
		resultUnknown, errUnknown := UpdateIngredientValueByName("UnknownIngredient", 10)
		if errUnknown == nil {
			t.Errorf("There should be an error because UnknownIngredient doesn't exist, got: result = %v", resultUnknown)
		}
	})

	for _, tt := range tests {
		testnameIngredient = fmt.Sprintf("UpdateIngredientValueByName %s,%d", tt.ingName, tt.value)
		t.Run(testnameIngredient, func(t *testing.T) {
			ans, _ := UpdateIngredientValueByName(tt.ingName, tt.value)
			if ans != tt.want {
				t.Errorf("UpdateIngredientValueByName falied, got %v, want %v", ans, tt.want)
			}
		})
	}
	t.Cleanup(func() {
		_, _ = UpdateIngredientPut(ingredientTestCleanUp)
	})
}
