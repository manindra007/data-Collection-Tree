package main

import (
	data "datacollectiontree/src/data"
	"fmt"
)

// type KeyValStruct struct {
// 	Key string `json:"key"`
// 	Val string `json:"val"`
// }

// type Insertstruct struct {
// 	Dim     []KeyValStruct `json:"dim"`
// 	Metrics []KeyValStruct `json:"metrics"`
// }

// type response struct {
// 	status string
// }

// func CollectionTree(w http.ResponseWriter, r *http.Request) {
// 	reqbody, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Println(w, "kindly provide data")
// 	}
// 	var newEvent Insertstruct
// 	json.Unmarshal(reqbody, &newEvent)
// 	w.WriteHeader(http.StatusCreated)

// 	var resp response
// 	cr, err := crwl.Webcrawler(newEvent)
// 	if err != nil {
// 		resp = response{
// 			Value: cr,
// 			Err:   err.Error(),
// 		}
// 	} else {
// 		resp = response{
// 			Value: cr,
// 			Err:   "",
// 		}
// 	}

// 	json.NewEncoder(w).Encode(resp)
// }

// func HandleRequest() {
// 	myRouter := mux.NewRouter().StrictSlash(true)
// 	myRouter.HandleFunc("v1/insert", CollectionTree).Methods("POST")
// 	log.Fatal(http.ListenAndServe(":8081", myRouter))
// }

func main() {
	fmt.Println("Data Collectiton  Tree")
	// HandleRequest()

	val := `{
		"dim": [{
			"key": "device",
			"val": "mobile"
			},
			{
				"key": "country",
				"val": "IN"
			}],
			"metrics": [{
				"key": "webreq",
				"val": 70
			},
			{
				"key": "timespent",
				"val": 30
			}]
		}`
	data.SetData([]byte(val))

	val = `{
		"dim": [{
			"key": "device",
			"val": "desktop"
			},
			{
				"key": "country",
				"val": "IN"
			}],
			"metrics": [{
				"key": "webreq",
				"val": 50
			},
			{
				"key": "timespent",
				"val": 40
			}]
		}`
	data.SetData([]byte(val))

	val = `{
		"dim": [{
			"key": "device",
			"val": "mobile"
			},
			{
				"key": "country",
				"val": "USA"
			}],
			"metrics": [{
				"key": "webreq",
				"val": 50
			},
			{
				"key": "timespent",
				"val": 40
			}]
		}`
	data.SetData([]byte(val))

	val = `{
		"dim": [{
		"key": "country"
		"val": "US"
		}]
	}`
	data.FetchData([]byte(val))
}
