package main

import (
	"fmt"
    "strconv"
)

//Inventory structure
type Inventory struct{
	Material string
	Quantity int
}

//Functions to create Beverages
func CreateBeverageGingerTea(array []Inventory) []Inventory{

	fmt.Println("\n\n\t Serving Ginger Tea...........")

	//Check the Quantity of ingridients in inventory
	//If empty Refill the inventory
	for i := 0; i < len(array); i++{
		if array[i].Quantity <= 0{
			fmt.Println("\n OOPS!!!! Ingridient Missing..........")
			fmt.Println("\n Refilling.........! Please try again.")
			array[i].Quantity += 500
		}
	}

	//Use ingridients from the inventory 
	for i := 0; i < len(array); i++{
		if array[i].Material == "Hot Water"{
			array[i].Quantity -= 50
		}
	}
	for i := 0; i < len(array); i++{
		if array[i].Material == "Hot Milk"{
			array[i].Quantity -= 10
		}
	}
	for i := 0; i < len(array); i++{
		if array[i].Material == "Tea Leaf Syrup"{
			array[i].Quantity -= 10
		}
	}
	for i := 0; i < len(array); i++{
		if array[i].Material == "Ginger Syrup"{
			array[i].Quantity -= 5
		}
	}
	for i := 0; i < len(array); i++{
		if array[i].Material == "Sugar Syrup"{
			array[i].Quantity -= 10
		}
	}
	return array
}

func CreateBeverageElaichiTea(array []Inventory) []Inventory{
	fmt.Println("\n\n\t Serving Elaichi Tea...........")

	//Check the Quantity of ingridients in inventory
	//If empty Refill the inventory
	for i := 0; i < len(array); i++{
		if array[i].Quantity <= 0{
			fmt.Println("\n OOPS!!!! Ingridient Missing..........")
			fmt.Println("\n Refilling.........! Please try again.")
			array[i].Quantity += 500
		}
	}

	//Use ingridients from the inventory 
	for i := 0; i < len(array); i++{
		if array[i].Material == "Hot Water"{
			array[i].Quantity -= 50
		}
	}
	for i := 0; i < len(array); i++{
		if array[i].Material == "Hot Milk"{
			array[i].Quantity -= 10
		}
	}
	for i := 0; i < len(array); i++{
		if array[i].Material == "Tea Leaf Syrup"{
			array[i].Quantity -= 10
		}
	}
	for i := 0; i < len(array); i++{
		if array[i].Material == "Elaichi Syrup"{
			array[i].Quantity -= 5
		}
	}
	for i := 0; i < len(array); i++{
		if array[i].Material == "Sugar Syrup"{
			array[i].Quantity -= 10
		}
	}
	return array
}

func CreateBeverageCoffee(array []Inventory) []Inventory{
	fmt.Println("\n\n\t Serving Coffee...........")

	//Check the Quantity of ingridients in inventory
	//If empty Refill the inventory
	for i := 0; i < len(array); i++{
		if array[i].Quantity <= 0{
			fmt.Println("\n OOPS!!!! Ingridient Missing..........")
			fmt.Println("\n Refilling.........! Please try again.")
			array[i].Quantity += 500
		}
	}
	//Use ingridients from the inventory 
	for i := 0; i < len(array); i++{
		if array[i].Material == "Hot Water"{
			array[i].Quantity -= 50
		}
	}
	for i := 0; i < len(array); i++{
		if array[i].Material == "Hot Milk"{
			array[i].Quantity -= 10
		}
	}

	for i := 0; i < len(array); i++{
		if array[i].Material == "Coffee Syrup"{
			array[i].Quantity -= 10
		}
	}
	for i := 0; i < len(array); i++{
		if array[i].Material == "Sugar Syrup"{
			array[i].Quantity -= 10
		}
	}
	return array
}

func CreateBeverageHotMilk(array []Inventory) []Inventory{
	fmt.Println("\n\n\t Serving Hot Milk...........")

	//Check the Quantity of ingridients in inventory
	//If empty Refill the inventory
	for i := 0; i < len(array); i++{
		if array[i].Quantity <= 0{
			fmt.Println("\n OOPS!!!! Ingridient Missing..........")
			fmt.Println("\n Refilling.........! Please try again.")
			array[i].Quantity += 500
		}
	}
	//Use ingridients from the inventory 
	for i := 0; i < len(array); i++{
		if array[i].Material == "Milk"{
			array[i].Quantity -= 50
		}
	}
	return array
}

func CreateBeverageHotWater(array []Inventory) []Inventory{
	fmt.Println("\n\n\t Serving Hot Water...........")

	//Check the Quantity of ingridients in inventory
	//If empty Refill the inventory
	for i := 0; i < len(array); i++{
		if array[i].Quantity <= 0{
			fmt.Println("\n OOPS!!!! Ingridient Missing..........")
			fmt.Println("\n Refilling.........! Please try again.")
			array[i].Quantity += 500
		}
	}

	//Use ingridients from the inventory 
	for i := 0; i < len(array); i++{
		if array[i].Material == "Water"{
			array[i].Quantity -= 50
		}
	}
	return array
}

//Function to create a new inventory to store ingredients
func CreateInventory() []Inventory{

	//Declare ingridents
	var HotWater Inventory
	HotWater.Material ="Hot Water"
	HotWater.Quantity = 50

	var HotMilk Inventory
	HotMilk.Material ="Hot Milk"
	HotMilk.Quantity = 500

	var TeaLeafSyrup Inventory
	TeaLeafSyrup.Material ="Tea Leaf Syrup"
	TeaLeafSyrup.Quantity = 100

	var GingerSyrup Inventory
	GingerSyrup.Material ="Ginger Syrup"
	GingerSyrup.Quantity = 50

	var SugarSyrup Inventory
	SugarSyrup.Material ="Sugar Syrup"
	SugarSyrup.Quantity = 100

	var ElaichiSyrup Inventory
	ElaichiSyrup.Material ="Elaichi Syrup"
	ElaichiSyrup.Quantity = 50

	var CoffeeSyrup Inventory
	CoffeeSyrup.Material ="Coffee Syrup"
	CoffeeSyrup.Quantity = 100

	var Milk Inventory
	Milk.Material ="Milk      "
	Milk.Quantity = 500

	var Water Inventory
	Water.Material ="Water     "
	Water.Quantity = 500

	//Declare inventory
	InventoryItems := []Inventory{}

	//add ingridents to inventory
	InventoryItems = append(InventoryItems, HotWater)
	InventoryItems = append(InventoryItems, HotMilk)
	InventoryItems = append(InventoryItems, TeaLeafSyrup)
	InventoryItems = append(InventoryItems, GingerSyrup)
	InventoryItems = append(InventoryItems, SugarSyrup)
	InventoryItems = append(InventoryItems, ElaichiSyrup)
	InventoryItems = append(InventoryItems, CoffeeSyrup)
	InventoryItems = append(InventoryItems, Water)
	InventoryItems = append(InventoryItems, Milk)

	return InventoryItems

}


// Main method
func main() {

	//List of Beverages
	Beverages := []string{"Ginger Tea", "Elaichi Tea", "Coffee", "Hot Milk", "Hot Water"}

	//create a new inventory
	var NewInventory []Inventory
	NewInventory = CreateInventory()
	fmt.Println("\n\t  Chai Point")
	fmt.Println("\n\tBeverages Available")
	for i := 0; i < len(Beverages); i++{
		fmt.Println("\n", Beverages[i])
	}

	fmt.Println("\n\tNew Inventory")
	for i := 0; i < len(NewInventory); i++{
		statement := NewInventory[i].Material + "\t\t\t" + strconv.Itoa(NewInventory[i].Quantity)
		fmt.Println("\n", statement)
	}
//----------------------------------------------CASE 1 : CHAI POINT MACHINE SERVING GINGER TEA-----------------------------------------------

	//Create Ginger tea and then updated the inventory
	NewInventory = CreateBeverageGingerTea(NewInventory)
	fmt.Println("\n\n\tCurrent Inventory")
	for i := 0; i < len(NewInventory); i++{
		statement := NewInventory[i].Material + "\t\t\t" + strconv.Itoa(NewInventory[i].Quantity)
		fmt.Println("\n", statement)
	}

//--------------------CASE 2 : SERVING ELAICHI TEA (NOTE: HOT WATER EMPTY IN INVENTORY THEREFORE IT IS REFILLED)-------------------------------

	//Create Ginger tea and then updated the inventory
	NewInventory = CreateBeverageElaichiTea(NewInventory)
	fmt.Println("\n\n\tCurrent Inventory")
	for i := 0; i < len(NewInventory); i++{
		statement := NewInventory[i].Material + "\t\t\t" + strconv.Itoa(NewInventory[i].Quantity)
		fmt.Println("\n", statement)
	}

//-----------------------------------EXTRA CASES : HERE ALL THE OTHER BEVERAGES CAN BE TEXTED(UNCOMMENT SECTIONS TO TEST)----------------------------

// //Create Ginger tea and then updated the inventory
// NewInventory = CreateBeverageCoffee(NewInventory)
// fmt.Println("\n\n\tCurrent Inventory")
// for i := 0; i < len(NewInventory); i++{
// 	statement := NewInventory[i].Material + "\t\t\t" + strconv.Itoa(NewInventory[i].Quantity)
// 	fmt.Println("\n", statement)
// }

// //Create Ginger tea and then updated the inventory
// 	NewInventory = CreateBeverageHotMilk(NewInventory)
// 	fmt.Println("\n\n\tCurrent Inventory")
// 	for i := 0; i < len(NewInventory); i++{
// 		statement := NewInventory[i].Material + "\t\t\t" + strconv.Itoa(NewInventory[i].Quantity)
// 		fmt.Println("\n", statement)
// 	}

// //Create Ginger tea and then updated the inventory
// NewInventory = CreateBeverageHotWater(NewInventory)
// fmt.Println("\n\n\tCurrent Inventory")
// 	for i := 0; i < len(NewInventory); i++{
// 	statement := NewInventory[i].Material + "\t\t\t" + strconv.Itoa(NewInventory[i].Quantity)
// 	fmt.Println("\n", statement)
// }

}