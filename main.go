/*

Descreption : This project is to handile all  apis
*/
package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"

	"log"
	"net/http"
	cg "prod/src/customlibrary/configuration"
	controller "prod/src/customlibrary/controller"
	sr "prod/src/customlibrary/services"
)

func getPostData(r *http.Request) map[string]string {

	decoder := json.NewDecoder(r.Body)

	t := map[string]interface{}{}
	err := decoder.Decode(&t)
	if err != nil {

		fmt.Println(r)
		log.Println(err)
	}
	postdata := map[string]string{}
	for key, val := range t {
		postdata[fmt.Sprint(key)] = fmt.Sprint(val)
	}
	return postdata
}

func SaveUserDetails(w http.ResponseWriter, r *http.Request) {

	postdata := getPostData(r)
	status := controller.SaveUserDetails(postdata)
	w.Header().Set("Content-Type", "application/json")
	if status {
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(400)
	}
	json.NewEncoder(w).Encode(postdata)
}

func UpdateUserDetails(w http.ResponseWriter, r *http.Request) {

	postdata := getPostData(r)
	status := controller.UpdateUserDetails(postdata)
	w.Header().Set("Content-Type", "application/json")
	if status {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(400)
	}
	json.NewEncoder(w).Encode(postdata)
}

func DeleteUserDetails(w http.ResponseWriter, r *http.Request) {

	postdata := getPostData(r)
	status := controller.DeleteUserDetails(postdata)
	w.Header().Set("Content-Type", "application/json")
	if status {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(400)
	}
	json.NewEncoder(w).Encode(postdata)
}

func GetAllUserDetails(w http.ResponseWriter, r *http.Request) {
	postdata := getPostData(r)
	status, response := controller.GetAllUserDetails(postdata)
	w.Header().Set("Content-Type", "application/json")
	if status {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(400)
	}
	json.NewEncoder(w).Encode(response)

}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte(`<!DOCTYPE html>
<html>
<head>
</head>
<body>
<p>Hello world</p>
</body>
</html>`))
}

func main() {

	cg.SetConfigParams()
	sr.SetDbConnection("ClientDatabase", cg.Config.ClientDatabase.MaxOpenConnections, cg.Config.ClientDatabase.MaxIdleConnections)

	fmt.Printf("Starting server at port " + cg.Config.Port1 + "\n")

	router := mux.NewRouter().StrictSlash(true)
	initaliseApiHandlers(router)

	if err := http.ListenAndServe(":"+cg.Config.Port1, router); err != nil {
		log.Fatal(err)
	}

}

func initaliseApiHandlers(router *mux.Router) {

	router.HandleFunc("/", homeHandler)

	router.HandleFunc("/user/save", SaveUserDetails).Methods("POST")
	router.HandleFunc("/user/update", UpdateUserDetails).Methods("PUT")
	router.HandleFunc("/user/delete", DeleteUserDetails).Methods("DELETE")
	router.HandleFunc("/user/list", GetAllUserDetails).Methods("GET")

}
