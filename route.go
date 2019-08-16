package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/trisnatera/posttest-individu/controller"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", controller.Welcome)
	router.HandleFunc("/content", controller.AddDataContent).Methods("POST")
	router.HandleFunc("/content", controller.GetAllDataContent).Methods("GET")
	router.HandleFunc("/content/{id}", controller.GetOneDataContent).Methods("GET")
	router.HandleFunc("/content/{id}", controller.UpdateDataContent).Methods("PATCH")
	router.HandleFunc("/content/{id}", controller.DeleteDataContent).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))

}
