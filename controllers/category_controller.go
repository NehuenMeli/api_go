package controllers

import (
	"../services"
	"../utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCategory( c *gin.Context){

	categoryID := c.Param(utils.ParamCategoryID)

	category,apiError := services.GetCategoryFromAPI(categoryID)
	if apiError!=nil {
		c.JSON(apiError.Status,apiError)
	}

	c.JSON(http.StatusOK,category)
}

func GetCategories( c *gin.Context){

	siteId := c.Param(utils.ParamSiteID)

	categories,apiError := services.GetCategoriesFromAPI(siteId)
	if apiError!=nil {
		c.JSON(apiError.Status,apiError)
	}

	c.JSON(http.StatusOK,categories)
}
