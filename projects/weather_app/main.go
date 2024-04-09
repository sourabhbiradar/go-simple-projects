package main

import (
	
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type apiConfigData struct{
	OpenWeatherMapApiKey string `json:"OpenWeatherMapApiKey"`

}

type weatherData struct{
	Name string `json:"name"`
	
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

func loadApiconfig(filename string) (apiConfigData,error){
	bytes,err:=ioutil.ReadFile(filename)

	if err!=nil{
		return apiConfigData{},err
	}
	var c apiConfigData
	err=json.Unmarshal(bytes,&c)

	if err!=nil{
		return apiConfigData{},err
    }
	return c,nil
}

func query(city string) (weatherData,error){
	apiConfig,err:=loadApiconfig(".apiConfig")

	if err!=nil{
		return weatherData{},err
	}

	resp,err:=http.Get("http://openweathermap.org/data/2.5/weather?APPID="+ apiConfig.OpenWeatherMapApiKey + "&q=" + city)

	if err!=nil{
		return weatherData{},err
	}
	defer resp.Body.Close()

	var d weatherData
	if err!=nil{
		return weatherData{},err
	}
	return d,nil
}

func main(){
http.HandleFunc("/weather/",
func (w http.ResponseWriter , r * http.Request){
	city:=strings.SplitN(r.URL.Path,"/",3)[2]

	data,err:=query(city)

	if err!=nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(data)
})
   http.ListenAndServe(":8080",nil)
}