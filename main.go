package main 

import (
	"fmt"
	"log"

	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
	"./connect"

  "./structures"
)

func main(){
	connect.InitializeDataBase()
	defer connect.CloseConnection()

	r := mux.NewRouter()
	r.HandleFunc("/users/{id}",GetUser).Methods("GET")
	r.HandleFunc("/users/new", NewUser).Methods("POST")
	r.HandleFunc("/users/update/{id}", UpdateUser).Methods("PATCH")
	r.HandleFunc("/users/delete/{id}", DeleteUser).Methods("DELETE")

	fmt.Println("Servidor en el puerto 8000")
	http.ListenAndServe(":8000", r)

}

func GetUser(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	user_id := vars["id"]

	status := "success"
	var message string
	user := connect.GetUser(user_id)

	if(user.Id <= 0){
		status = "error"
		message = "User not found"
	}

	response := structures.Reponse{ status, user, message}
	json.NewEncoder(w).Encode(response)
}

func NewUser(w http.ResponseWriter, r *http.Request){
	user := GetUserRequest(r)
	response := structures.Reponse{ "success", connect.CreateUser(user), ""}
	json.NewEncoder(w).Encode(response)
}

func UpdateUser(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	user_id := vars["id"]

	user := GetUserRequest(r)
	response := structures.Reponse{ "success", connect.UpdateUser(user_id, user), ""}
	json.NewEncoder(w).Encode(response)
}


func DeleteUser(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	user_id := vars["id"]
	connect.DeleteUser(user_id)

	user := structures.User{}

	response := structures.Reponse{ "success",user , ""}
	json.NewEncoder(w).Encode(response)

}


func GetUserRequest(r *http.Request)structures.User {
	var user structures.User

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)

	if err != nil {
		log.Fatal(err)
	}
	return user
}



