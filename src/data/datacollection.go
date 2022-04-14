package data

import (
	"encoding/json"
	"fmt"
)

type WebCountry struct {
	Country    string
	WebReqeust int
	TimeSpent  int
	DeviceType map[string]*WebDevice
}

type WebDevice struct {
	Device     string
	WebReqeust int
	TimeSpent  int
}

type WebTime struct {
	WebReqeust  int
	TimeSpent   int
	CountryType map[string]WebCountry
}

var WebUsage WebTime

type InputKeyString struct {
	Key string `json:"key"`
	Val string `json:"val"`
}
type InputKeyInt struct {
	Key string `json:"key"`
	Val int    `json:"val"`
}

type InputData struct {
	// Data int `json:"data"`
	Dim     []InputKeyString `json:"dim"`
	Metrics []InputKeyInt    `json:"metrics"`
}

func SetData(v []byte) {
	// fmt.Println(string(v))
	var ID InputData
	var err error
	if err = json.Unmarshal(v, &ID); err != nil {
		fmt.Println("error occured", err.Error())
		return
	}

	var country, devtype string
	var webr, times int
	if ID.Dim[0].Key == "country" && ID.Dim[1].Key == "device" {
		country = ID.Dim[0].Val
		devtype = ID.Dim[1].Val
	} else if ID.Dim[1].Key == "country" && ID.Dim[0].Key == "device" {
		country = ID.Dim[1].Val
		devtype = ID.Dim[0].Val
	} else {
		fmt.Println("error")
	}
	if ID.Metrics[0].Key == "webreq" && ID.Metrics[1].Key == "timespent" {
		webr = ID.Metrics[0].Val
		times = ID.Metrics[1].Val
	} else if ID.Metrics[1].Key == "webreq" && ID.Metrics[0].Key == "timespent" {
		webr = ID.Metrics[1].Val
		times = ID.Metrics[0].Val
	} else {
		fmt.Println("error")
	}

	WebUsage.WebReqeust = WebUsage.WebReqeust + webr
	WebUsage.TimeSpent = WebUsage.TimeSpent + times

	if WebUsage.CountryType == nil {
		WebUsage.CountryType = make(map[string]WebCountry)
	}
	var val WebCountry
	var ok bool
	if val, ok = WebUsage.CountryType[country]; !ok {
		WebUsage.CountryType[country] = WebCountry{}
	}
	val.WebReqeust = webr + val.WebReqeust
	val.TimeSpent = times + val.TimeSpent
	if val.DeviceType == nil {
		val.DeviceType = make(map[string]*WebDevice)
		// WebUsage.CountryType[country] = val
	}
	// } else {
	// }

	WebUsage.CountryType[country] = val
	if dev, ok := WebUsage.CountryType[country].DeviceType[devtype]; !ok {
		WebUsage.CountryType[country].DeviceType[devtype] = &WebDevice{
			Device:     devtype,
			WebReqeust: webr,
			TimeSpent:  times,
		}
	} else {
		WebUsage.CountryType[country].DeviceType[devtype] = &WebDevice{
			WebReqeust: dev.WebReqeust + webr,
			TimeSpent:  dev.TimeSpent + times,
		}
	}
	// fmt.Println("h3")
	// fmt.Println(WebUsage)
	// fmt.Println(WebUsage.CountryType[country])
	// fmt.Println(WebUsage.CountryType[country].DeviceType[devtype])
}

func FetchData(v []byte) {
	fmt.Println(WebUsage)

}
