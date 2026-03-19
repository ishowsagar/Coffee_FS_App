package controllers

import (
	"encoding/json"
	"net/http"
	"server/helpers"
	"server/services"
	"strconv"

	"github.com/go-chi/chi/v5"
)

var coffee services.Coffee
var models services.Models //@ Serves access to type that stores Coffee & JsonResponse types that implements methods Belongs to them
// Get all coffees
func GetAllCoffees(w http.ResponseWriter,r *http.Request) {
	all,err := models.Coffee.GetAllCoffees()
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJson(w,http.StatusOK,helpers.Envelop{"coffees":all})
}

// Get Coffee by id slug ({id}👇👇)
func GetCoffeeByID(w http.ResponseWriter,r *http.Request) {
	id := chi.URLParam(r,"id")
	// coffee,err := coffee.GetCoffeeByID(id)
	coffee,err := models.Coffee.GetCoffeeByID(id)

	// if caught error getting coffee
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	// otherwise successfully got coffee by calling that func query --> send res back to client✅✅
	helpers.WriteJson(w,http.StatusOK,helpers.Envelop{"coffee":coffee})
}

// Get Coffee by  slug ({name}👇👇) 
func GetCoffeeByName(w http.ResponseWriter,r *http.Request) {
	name_slug := chi.URLParam(r,"name")
	retrieved_coffee,err := models.Coffee.GetCoffeeByName(name_slug)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	// sending response back to client
	helpers.WriteJson(w,http.StatusOK,helpers.Envelop{"coffee":retrieved_coffee})

}

// Get Coffee by query Params ({?price=__}👇👇) 
func GetCoffeeByPriceQP(w http.ResponseWriter, r *http.Request) {
	
	// fetch client url params
	price_param := r.URL.Query().Get("price")
	if price_param == "" {
		http.Error(w, "missing price query param", http.StatusBadRequest)
		return
	}

	// @ converts string query param into a decimal price value
	priceValue, err := strconv.ParseFloat(price_param, 32)
	if err != nil {
		http.Error(w,"pass a valid numeric price only",http.StatusBadRequest)
		return
	}

	// successfully fetched numeric value of price q.p
	retrieved_coffee,err := models.Coffee.GetCoffeeByqparamsPrice(float32(priceValue))
	if err != nil {
		http.Error(w,"failed to get coffee due to unknown price passed to it",http.StatusBadRequest)
		return
	}

	// send response to the client
	helpers.WriteJson(w,http.StatusOK,helpers.Envelop{"coffee":retrieved_coffee})
} 

// Get Coffee by query Params ({?region=__}👇👇) 
func GetCoffeeByQueryParams(w http.ResponseWriter,r *http.Request) {
	
	// # step 1 => fetch query param from URL using r data struct
	region_qparam := r.URL.Query().Get("region")
	if region_qparam == "" {
		region_qparam = chi.URLParam(r,"region")
	}
	if region_qparam == "" {
		http.Error(w,"missing region query param",http.StatusBadRequest)
		return
	}
	
	// # step 2 => Invoke method that belongs to Coffee type to call db to get coffee data by passing query param
	retrieved_coffee,err := models.Coffee.GetCoffeeByQueryParams(region_qparam)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	// # step 3 => Send response back to client using helper function
	helpers.WriteJson(w,http.StatusOK,helpers.Envelop{"coffee":retrieved_coffee})

}
//@PUT ---- Update Coffee by id slug ({id}👇👇)
func UpdateCoffeeByID(w http.ResponseWriter,r *http.Request) {
	id := chi.URLParam(r,"id")
	var coffeeUpdateInputVar services.Coffee
	err:= json.NewDecoder(r.Body).Decode(&coffeeUpdateInputVar)
	// if caught error decoding r.Body into coffeeupdateinputvar
	if err != nil {
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}
	// updatedCoffee,err:= coffee.UpdateCoffee(id,coffeeUpdateInputVar)
	updatedCoffee,err:= models.Coffee.UpdateCoffee(id,coffeeUpdateInputVar)
	// if caught error getting coffee
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	// otherwise successfully got coffee by calling that func query --> send res back to client✅✅
	helpers.WriteJson(w,http.StatusOK,helpers.Envelop{"updated_coffee":updatedCoffee})
}

// creates coffee
func CreateCoffee(w http.ResponseWriter,r *http.Request) {
	var coffeeData services.Coffee
	//! what decoder doing is --> # decoding whats coming from req.Body and injecting into coffeeData
	err := json.NewDecoder(r.Body).Decode(&coffeeData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	// at this point --> ✅✅successfully have decoded data into coffeeData instance which has all the required data
	
	// & Since we got coffee data, now we can pass to func that query it to createCoffee
	coffeeCreated,err := models.Coffee.CreateCoffee(coffeeData) // ✅✅ creates coffee 
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	// sending response back to client
	helpers.WriteJson(w,http.StatusOK,helpers.Envelop{"coffee":coffeeCreated})
}

func DeleteCoffee(w http.ResponseWriter,r *http.Request) {
	id := chi.URLParam(r,"id")
	err := models.Coffee.DeleteCoffeeByID(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	// sending response back to client
	helpers.WriteJson(w,http.StatusOK,helpers.Envelop{"status":"succesfully deleted🛑🛑!"})

}