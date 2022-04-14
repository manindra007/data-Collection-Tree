package data

import (
	"encoding/json"
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

type InputSetData struct {
	Dim     []InputKeyString `json:"dim"`
	Metrics []InputKeyInt    `json:"metrics"`
}

type InputGetData struct {
	Dim []InputKeyString `json:"dim"`
}

type OutPutResponse struct {
	Status string     `json:"Res"`
	Output OutputData `json:"Output"`
}

type OutputData struct {
	Dim     []OutputKeyString `json:"dim"`
	Metrics []OutputkeyInt    `json:"metrics"`
}

type OutputkeyInt struct {
	Key string `json:"key"`
	Val int    `json:"val"`
}

type OutputKeyString struct {
	Key string `json:"key"`
	Val string `json:"val"`
}

func SetData(v []byte) *OutPutResponse {
	var ID InputSetData
	var err error
	if err = json.Unmarshal(v, &ID); err != nil {
		return &OutPutResponse{
			Status: "404 data not found",
		}
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
		return &OutPutResponse{
			Status: "404 data not found",
		}
	}
	if ID.Metrics[0].Key == "webreq" && ID.Metrics[1].Key == "timespent" {
		webr = ID.Metrics[0].Val
		times = ID.Metrics[1].Val
	} else if ID.Metrics[1].Key == "webreq" && ID.Metrics[0].Key == "timespent" {
		webr = ID.Metrics[1].Val
		times = ID.Metrics[0].Val
	} else {
		return &OutPutResponse{
			Status: "404 data not found",
		}
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
	}
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
	return &OutPutResponse{
		Status: "200 OK",
	}
}

func FetchData(v []byte) *OutPutResponse {
	var ID InputGetData
	var err error
	if err = json.Unmarshal(v, &ID); err != nil {
		return &OutPutResponse{
			Status: "404 data not found",
		}
	}

	var res *OutPutResponse
	for _, val := range ID.Dim {
		var country string
		if val.Key == "country" {
			country = ID.Dim[0].Val
		} else {
			return &OutPutResponse{
				Status: "404 data not found",
			}
		}
		var r OutputData
		r.Dim = make([]OutputKeyString, 0)
		d := OutputKeyString{Key: "country", Val: country}
		r.Dim = append(r.Dim, d)
		r.Metrics = make([]OutputkeyInt, 0)
		m1 := OutputkeyInt{Key: "webreq", Val: WebUsage.CountryType[country].WebReqeust}
		m2 := OutputkeyInt{Key: "timespent", Val: WebUsage.CountryType[country].TimeSpent}
		r.Metrics = append(r.Metrics, m1)
		r.Metrics = append(r.Metrics, m2)

		res = &OutPutResponse{
			Status: "200 OK",
			Output: OutputData{Dim: r.Dim, Metrics: r.Metrics},
		}
	}
	return res

}
