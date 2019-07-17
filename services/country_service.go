package services

import (
	"../utils"
	"../models"
)

func GetCountryFromAPI(id string) (*models.Country,*utils.ApiError) {
	country := &models.Country{ID:id}

	if apiError := country.Get(); apiError !=nil{
		return nil, apiError
	}
	return country,nil
}

func GetCountriesFromAPI() (*[]models.Country,*utils.ApiError) {
	countries,err := models.GetAllCountries()
	if  err !=nil{
		return nil, err
	}
	return countries,nil
}

func AsyncGetCountryFromAPI(country *models.Country,channel *chan *utils.ApiError)  {
	*channel<-country.Get()
}