package main

import (
	"./controllers"
	"./utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

const (
	port = ":8080"
)

var (
	router = gin.Default()
)

func main() {

	go func() {
		for range time.Tick(time.Second * controllers.BurstTime) {
			for i := 0; i < controllers.BurstLimit; i++ {
				controllers.Burst <- i
			}
		}
	}()

	router.GET(fmt.Sprintf("/user/:%s", utils.ParamUserID), controllers.GetUser)
	router.GET(fmt.Sprintf("/site/:%s", utils.ParamSiteID), controllers.GetSite)
	router.GET("/sites", controllers.GetSites)
	router.GET(fmt.Sprintf("/country/:%s", utils.ParamCountryID), controllers.GetCountry)
	router.GET("/countries", controllers.GetCountries)
	router.GET(fmt.Sprintf("/categories/:%s", utils.ParamSiteID), controllers.GetCategories)
	router.GET(fmt.Sprintf("/category/:%s", utils.ParamCategoryID), controllers.GetCategory)
	router.GET(fmt.Sprintf("/result/:%s", utils.ParamUserID), controllers.HelperGetResult)
	err := router.Run(port)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
