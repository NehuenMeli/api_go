package services

import (
	"../models"
	"../utils"
)

func GetUserFromAPI(id int64) (*models.User,*utils.ApiError) {
	user := &models.User{ID:id}

	if apiError := user.Get(); apiError !=nil{
		return nil, apiError
	}
	return user,nil
}


