package models

import (
"encoding/json"
"fmt"
"io/ioutil"
"net/http"
"../utils"
)

type Country struct {
	ID               	string `json:"id"`
	Name         	 	string `json:"name"`
	Locale 			 	string `json:"locale"`
	CurrencyID       	string `json:"currency_id"`
	DecimalSeparator 	string `json:"decimal_separator"`
	ThousandsSeparator 	string `json:"thousands_separator"`
	TimeZone 			string `json:"time_zone"`
	GeoInformation 		struct{
		Location 		struct{
			Latitude 	float64 `json:"latitude"`
			Longitude 	float64 `json:"longitude"`
		} `json:"location"`
	} `json:"geo_information"`
	States				[]struct{
		ID				string `json:"id"`
		Name 			string `json:"name"`
	}
}

func (country *Country) Get() *utils.ApiError{

	if country.ID == "" {
		fmt.Print("El id del país se encuentra vacío")
		return &utils.ApiError{
			Message: "El id del país se encuentra vacío",
			Status: http.StatusBadRequest,
		}
	}

	url := fmt.Sprintf("%s%s", utils.UrlCountries, country.ID)

	fmt.Print(url)

	response, err := http.Get(url)
	if err != nil{
		return &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	if err := json.Unmarshal(data, &country); err != nil{
		return &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	return nil
}

func GetAllCountries() (*[]Country,*utils.ApiError) {
	var countries []Country

	response, e := http.Get(utils.UrlCountries)
	if e !=nil {
		return nil,&utils.ApiError{Message: e.Error(),Status:http.StatusInternalServerError}
	}

	data, e := ioutil.ReadAll(response.Body)
	if e !=nil {
		return nil,&utils.ApiError{Message: e.Error(),Status:http.StatusInternalServerError}
	}

	if e = json.Unmarshal(data,&countries); e !=nil{
		return nil,&utils.ApiError{Message: e.Error(),Status:http.StatusInternalServerError}
	}

	return &countries,nil;
}


