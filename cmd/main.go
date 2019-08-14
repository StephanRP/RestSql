package main

import (
	"RestSQL/pkg/api"
	"RestSQL/pkg/config"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Rest API v2.0 Electric Boogaloo")
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/Health", api.HealthCheck).Methods("GET")
	myRouter.HandleFunc("/all", api.Handler(config.MemberAll)).Methods("GET")
	myRouter.HandleFunc("/deleteMember/{id}", api.Handler(config.DeleteMember)).Methods("DELETE")
	myRouter.HandleFunc("/memberSearch/{name}", api.Handler(config.MemberSearch)).Methods("POST")
	myRouter.HandleFunc("/addMember/{name}/{lob}/{pcp}", api.Handler(config.AddMember)).Methods("PUT")
	myRouter.HandleFunc("/updateMember/{id}/{update}/{new}", api.Handler(config.UpdateMember)).Methods("PUT")
	log.Panic(http.ListenAndServe(":8881", myRouter))
}
