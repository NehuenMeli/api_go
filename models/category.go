package models

import(
	"../utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Category struct {
	ID 			string `json:"id"`
	Name 		string `json:"name"`
	Imagen 		string `json:"picture"`
	TotalItems  string `json:"total_items_in_this_category"`
}

func (category *Category) Get() *utils.ApiError  {
	if category.ID == "" {
		return &utils.ApiError{
			Message:"El id de la categoría se encuentra vacío.",
			Status: http.StatusBadRequest,
		}
	}

	url := fmt.Sprintf("%s%s", utils.UrlCategories,category.ID)

	response, e := http.Get(url)
	if e !=nil {
		return &utils.ApiError{Message: e.Error(),Status:http.StatusInternalServerError}
	}

	data, e := ioutil.ReadAll(response.Body)
	if e !=nil {
		return &utils.ApiError{Message: e.Error(),Status:http.StatusInternalServerError}
	}

	if e = json.Unmarshal(data,&category); e !=nil{
		return &utils.ApiError{Message: e.Error(),Status:http.StatusInternalServerError}
	}

	return nil;
}

func GetAllCategories(siteId string) (*[]Category,*utils.ApiError) {

	var categories []Category

	if siteId == "" {
		return nil,&utils.ApiError{
			Message:"El id del site se encuentra vacío.",
			Status: http.StatusBadRequest,
		}
	}

	response, e := http.Get(fmt.Sprintf(utils.UrlCategoriesBySite,siteId))
	if e != nil {
		return nil, &utils.ApiError{Message: e.Error(), Status: http.StatusInternalServerError}
	}

	data, e := ioutil.ReadAll(response.Body)
	if e != nil {
		return nil, &utils.ApiError{Message: e.Error(), Status: http.StatusInternalServerError}
	}

	if e = json.Unmarshal(data, &categories); e != nil {
		return nil, &utils.ApiError{Message: e.Error(), Status: http.StatusInternalServerError}
	}

	return &categories,nil
}
