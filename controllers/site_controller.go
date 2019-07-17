package controllers

import (
	"../services"
	"../utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSite( c *gin.Context){

	siteID := c.Param(utils.ParamSiteID)

	site,apiError := services.GetSiteFromAPI(siteID)
	if apiError!=nil {
		c.JSON(apiError.Status,apiError)
	}

	c.JSON(http.StatusOK,site)
}

func GetSites( c *gin.Context){

	sites,apiError := services.GetSitesFromAPI()
	if apiError!=nil {
		c.JSON(apiError.Status,apiError)
	}

	c.JSON(http.StatusOK,sites)
}
