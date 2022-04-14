package main

import (
	"datacollectiontree/src/data"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func PushTree(w http.ResponseWriter, r *http.Request) {

	reqbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(w, "kindly provide data")
	}
	var newEvent data.InputSetData
	json.Unmarshal(reqbody, &newEvent)
	w.WriteHeader(http.StatusCreated)

	var resp data.OutPutResponse
	cr := data.SetData(reqbody)
	if cr.Status == "404 data not found" {
		resp = data.OutPutResponse{
			Status: cr.Status,
		}
	} else {
		resp = data.OutPutResponse{
			Status: cr.Status,
			Output: cr.Output,
		}
	}

	json.NewEncoder(w).Encode(resp)
}

func GetTree(w http.ResponseWriter, r *http.Request) {
	reqbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(w, "kindly provide data")
	}
	var newEvent data.InputSetData
	json.Unmarshal(reqbody, &newEvent)
	w.WriteHeader(http.StatusCreated)

	var resp data.OutPutResponse
	cr := data.FetchData(reqbody)
	if cr.Status == "404 data not found" {
		resp = data.OutPutResponse{
			Status: cr.Status,
		}
	} else {
		resp = data.OutPutResponse{
			Status: cr.Status,
			Output: cr.Output,
		}
	}

	json.NewEncoder(w).Encode(resp)
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/v1/insert", PushTree).Methods("POST")
	myRouter.HandleFunc("/v1/query", GetTree).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}
func main() {
	fmt.Println("started running")
	handleRequest()
}
