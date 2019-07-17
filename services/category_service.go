package services

import (
	"../utils"
	"../models"
)

func GetCategoryFromAPI(id string) (*models.Category,*utils.ApiError) {
	category := &models.Category{ID:id}

	if apiError := category.Get(); apiError !=nil{
		return nil, apiError
	}
	return category,nil
}

func GetCategoriesFromAPI(id string) (*[]models.Category,*utils.ApiError) {
	categories,err := models.GetAllCategories(id)
	if  err !=nil{
		return nil, err
	}
	return categories,nil
}