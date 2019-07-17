package controllers

import (
	"../services"
	"../utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCountry( c *gin.Context){

	countryID := c.Param(utils.ParamCountryID)

	country,apiError := services.GetCountryFromAPI(countryID)
	if apiError!=nil {
		c.JSON(apiError.Status,apiError)
	}

	c.JSON(http.StatusOK,country)
}

func GetCountries( c *gin.Context){
	countries,apiError := services.GetCountriesFromAPI()
	if apiError!=nil {
		c.JSON(apiError.Status,apiError)
	}

	c.JSON(http.StatusOK,countries)
}
