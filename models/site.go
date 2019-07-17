package models

import (
	"../utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Site struct {
	ID 						string `json:"id"`
	Name					string `json:"name"`
	CountryId				string `json:"country_id"`
	SaleFeesMode			string `json:"sale_fees_mode"`
	MercadoPagoVersion 		int    `json:"mercado_pago_version"`
	DefaultCurrencyID		string `json:"default_currency_id"`
	ImmediatePayment		string `json:"immediate_payment"`
	PaymentMethodIDs	  	[]string `json:"payment_method_ids"`
	Settings				struct{
		IdentificationTypes		[]string `json:"identification_types"`
		TaxPayerTypes			[]string `json:"taxpayer_types"`
		IdentificationTypeRules	struct{
			IdentificationType	string `json:"identification_type"`
			Rules				[]struct{
				EnabledTaxPayerTypes		[]string `json:"enabled_taxpayer_types"`
				BeginsWith					string `json:"begins_with"`
				Type 						string `json:"type"`
				MinLength					int `json:"min_length"`
				MaxLength					int `json:"max_length"`
			} `json:"rules"`
		}
	} `json:"settings"`
	Currencies				[]struct{
		ID 		string `json:"id"`
		Symbol 	string `json:"symbol"`
	} `json:"currencies"`
	Categories					[]struct{
		ID 		string `json:"id"`
		Name 	string `json:"name"`
	}

}

func (site *Site) Get() *utils.ApiError  {
	if site.ID == "" {
		return &utils.ApiError{
			Message:"El id del site se encuentra vacío.",
			Status: http.StatusBadRequest,
		}
	}

	url := fmt.Sprintf("%s%s",utils.UrlSites,site.ID)

	fmt.Println(url)

	response, e := http.Get(url)
	if e !=nil {
		return &utils.ApiError{Message: e.Error(),Status:http.StatusInternalServerError}
	}

	data, e := ioutil.ReadAll(response.Body)
	if e !=nil {
		return &utils.ApiError{Message: e.Error(),Status:http.StatusInternalServerError}
	}

	if e = json.Unmarshal(data,&site); e !=nil{
		return &utils.ApiError{Message: e.Error(),Status:http.StatusInternalServerError}
	}

	return nil;
}

func GetAllSites() (*[]Site,*utils.ApiError) {
	var sites []Site

	response, e := http.Get(utils.UrlSites)
	if e !=nil {
		return nil,&utils.ApiError{Message: e.Error(),Status:http.StatusInternalServerError}
	}

	data, e := ioutil.ReadAll(response.Body)
	if e !=nil {
		return nil,&utils.ApiError{Message: e.Error(),Status:http.StatusInternalServerError}
	}

	if e = json.Unmarshal(data,&sites); e !=nil{
		return nil,&utils.ApiError{Message: e.Error(),Status:http.StatusInternalServerError}
	}

	return &sites,nil;
}

