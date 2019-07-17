package controllers

import (
	"../services"
	"../utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetResult(c *gin.Context)  {

	userID := c.Param(utils.ParamUserID)

	id,err := strconv.ParseInt(userID,10,64)
	if err != nil {
		apiError := &utils.ApiError{err.Error(),http.StatusBadRequest}
		c.JSON(apiError.Status,apiError)
	}

	result,apiError := services.GetResult(id)
	if apiError!=nil {
		c.JSON(apiError.Status,apiError)
	}

	c.JSON(http.StatusOK,result)
}
