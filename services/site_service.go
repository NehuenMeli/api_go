package services

import (
	"../utils"
	"../models"
)

func GetSiteFromAPI(id string) (*models.Site,*utils.ApiError) {
	site := &models.Site{ID:id}

	if apiError := site.Get(); apiError !=nil{
		return nil, apiError
	}
	return site,nil
}

func GetSitesFromAPI() (*[]models.Site,*utils.ApiError) {
	sites,err := models.GetAllSites()
	if  err !=nil{
		return nil, err
	}
	return sites,nil
}

func AsyncGetSiteFromAPI(site *models.Site,channel *chan *utils.ApiError)  {
	*channel<-site.Get()
}
