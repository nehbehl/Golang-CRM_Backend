package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

//Customer data structure
type Customer struct {
	ID        uuid.UUID
	Name      string
	Role      string
	Email     string
	Phone     uint64
	Contacted bool
}

//Mock database to store customer data
var mockDB = []Customer{}

//Handler function to fetch all customer data
func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mockDB)
}

//Hanlder function to fetch a specific customer data
func getCustomer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var getCustomerInfo = Customer{}
	var flag bool

	w.Header().Set("Content-Type", "application/json")

	//Iterating through customer mock database to 
	//get customer specific customer data using ID passed in request
	for _, a := range mockDB {
		if a.ID.String() == vars["id"] {
			getCustomerInfo = a
			flag = true
			break
		}
	}

	//Write response
	if flag {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(getCustomerInfo)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Record not found")
	}
}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newEntry = Customer{}
	
	//Read request body and add to mock customer database
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &newEntry)
	newEntry.ID = generateUUID()
	mockDB = append(mockDB, newEntry) //Appending record
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(mockDB)
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	var index = 0
	var flag bool

	//Iterating through mock database to find the record
	//using id passed in request parameters
	for _, a := range mockDB {
		if a.ID.String() == vars["id"] {
			flag = true
			break
		}
		index++
	}
	
	//Writing response
	if flag {
		if(index==len(mockDB)-1) {
			mockDB = mockDB[:index-1] //Handling case when last record needs to be deleted
		}else {
			mockDB = append(mockDB[:index], mockDB[index+1])	//Deleting specific record	
		}		
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Record deleted successfully")
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Record not found")
	}
	
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	var getCustomerInfo = Customer{}
	var flag bool
	var index = 0
	w.Header().Set("Content-Type", "application/json")

	for _, a := range mockDB {
		if a.ID.String() == vars["id"] {
			reqBody, _ := ioutil.ReadAll(r.Body)
			json.Unmarshal(reqBody, &a)
			getCustomerInfo = a
			flag = true
			break
		}
		index++
	}

	if flag {
		w.WriteHeader(http.StatusOK)
		mockDB=append(mockDB[:index],getCustomerInfo)
		json.NewEncoder(w).Encode(getCustomerInfo)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Record not found")
	}
}

func index(w http.ResponseWriter, r * http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

func generateUUID() uuid.UUID {
	id := uuid.New()
	return id
}

// Installed go get github.com/google/uuid for generating unique customer id

func main() {
	router := mux.NewRouter()
	mockDB = append(mockDB, Customer{ID: generateUUID(), Name: "TestUser1", Role: "tester", Email: "tester1@test.com", Phone: 110001, Contacted: true})
	mockDB = append(mockDB, Customer{ID: generateUUID(), Name: "TestUser2", Role: "tester", Email: "tester2@test.com", Phone: 110401, Contacted: false})

	//Route calls
	router.HandleFunc("/", index).Methods("GET")	
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")

	fmt.Println("Server is starting on port 3000...")
	http.ListenAndServe(":3000", router)
}
