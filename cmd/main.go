package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"tendermint-client/internal/model"
	"tendermint-client/internal/service"

	"github.com/gorilla/mux"
	"github.com/tkanos/gonfig"
)

type params struct {
	coin   string
	text   string
	txhash string
}

func sendTx(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := req.ParseForm()
	if err != nil {
		panic(err)
	}
	key := req.Form.Get("key")
	value := req.Form.Get("value")

	var result model.TransactionCommit

	result = service.SendTx(key, value)

	json.NewEncoder(w).Encode(result)
}

func queryKey(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := req.ParseForm()
	if err != nil {
		panic(err)
	}
	key := req.Form.Get("key")

	var result model.TransactionData

	result = service.QueryKey(key)

	json.NewEncoder(w).Encode(result)
}

func main() {
	configuration := model.Configuration{}
	err := gonfig.GetConf("./config.json", &configuration)
	if err != nil {
		fmt.Printf("Error Occurred while loading env %s\n", err)
	}
	fmt.Println("Tendermint Client Active, Running @ ", configuration.Port)

	router := mux.NewRouter()
	router.HandleFunc("/sendTx", sendTx).Methods("POST")
	router.HandleFunc("/queryKey", queryKey).Methods("POST")
	log.Fatal(http.ListenAndServe(":"+configuration.Port, router))
}
