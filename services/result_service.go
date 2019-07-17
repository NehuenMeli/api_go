package services

import (
	"../utils"
	"../models"
	"sync"
)

func GetResult(id int64) (*models.Result,*utils.ApiError) {

	var waitGroup sync.WaitGroup
	channel := make(chan *utils.ApiError)
	var result models.Result
	var apiError *utils.ApiError

	defer close(channel)

	user := &models.User{ID:id}
	if apiError := user.Get(); apiError !=nil{
		return nil, apiError
	}

	result.User = user
	result.Site = &models.Site{ID:user.SiteID}
	result.Country = &models.Country{ID:user.CountryID}

	waitGroup.Add(2)

	go func() {
		for i:=0;i<2 ;i++  {
			apiError = <-channel
			if apiError!=nil {
				apiError = apiError
			}
			if result.Site.Name!="" {

			}
			waitGroup.Done()
		}
	}()

	go AsyncGetSiteFromAPI(result.Site,&channel)
	go AsyncGetCountryFromAPI(result.Country,&channel)

	waitGroup.Wait()

	if apiError!=nil {
		return nil,apiError
	}

	return &result,nil
}
