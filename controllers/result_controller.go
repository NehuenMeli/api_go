package controllers

import (
	"../services"
	"../utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	BurstLimit = 3
	BurstTime  = 5
)

var (
	Burst = make(chan int, BurstLimit)
	queue = 0
)

func GetResult(c *gin.Context) {

	userID := c.Param(utils.ParamUserID)

	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		apiError := &utils.ApiError{err.Error(), http.StatusBadRequest}
		c.JSON(apiError.Status, apiError)
	}

	result, apiError := services.GetResult(id)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
	}

	c.JSON(http.StatusOK, result)
}

func HelperGetResult(c *gin.Context) {
	fmt.Println("Numero de llamados:", queue)
	if queue > 3 {
		c.JSON(429, "Too Many Request")
		return
	}
	queue++
	<-Burst
	queue--
	GetResult(c)
}
