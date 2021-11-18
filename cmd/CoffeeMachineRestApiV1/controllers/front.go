package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterControllers() {
	cmContoller := newCoffeeMachineController()
	iContoller := newIngredientsController()
	dContoller := newDrinkController()
	mContoller := newMoneyController()

	http.Handle("/coffeemachine", *cmContoller)
	http.Handle("/coffeemachine/", *cmContoller)
	http.Handle("/coffeemachine/ingredients", *iContoller)
	http.Handle("/coffeemachine/ingredients/", *iContoller)
	http.Handle("/coffeemachine/drinks", *dContoller)
	http.Handle("/coffeemachine/drinks/", *dContoller)
	http.Handle("/coffeemachine/money", *mContoller)
	http.Handle("/coffeemachine/money/", *mContoller)
	http.Handle("/Coffeemachine", *cmContoller)
	http.Handle("/Coffeemachine/", *cmContoller)
	http.Handle("/coffeemachine/Ingredients", *iContoller)
	http.Handle("/coffeemachine/Ingredients/", *iContoller)
	http.Handle("/coffeemachine/Drinks", *dContoller)
	http.Handle("/coffeemachine/Drinks/", *dContoller)
	http.Handle("/coffeemachine/Money", *mContoller)
	http.Handle("/coffeemachine/Money/", *mContoller)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
