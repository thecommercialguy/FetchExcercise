package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"

	"github.com/google/uuid"
)

type Reciept struct {
	ID           string `json:"id"`
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

func (cfg *apiConfig) handlerProcessReciepts(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Retailer     string `json:"retailer"`
		PurchaseDate string `json:"purchaseDate"`
		PurchaseTime string `json:"purchaseTime"`
		Items        []Item `json:"items"`
		Total        string `json:"total"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "The receipt is invalid.", err)
		return
	}

	newUUID := uuid.New()
	uuidString := newUUID.String()

	newReciept := Reciept{
		ID:           uuidString,
		Retailer:     params.Retailer,
		PurchaseDate: params.PurchaseDate,
		PurchaseTime: params.PurchaseTime,
		Items:        params.Items,
		Total:        params.Total,
	}

	valid := validateReceipt(newReciept)
	if !valid {
		respondWithError(w, http.StatusBadRequest, "The receipt is invalid.", err)
		return
	}

	cfg.DB.Store(uuidString, newReciept)

	respondWithJSON(w, http.StatusOK, newReciept.ID)

}

func validateReceipt(reciept Reciept) bool {
	textPattern := regexp.MustCompile(`^[\w\s&-]+$`)
	valid := textPattern.MatchString(reciept.Retailer)
	if !valid {
		return valid
	}

	purchaseDatePattern := regexp.MustCompile(`\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01])`)
	valid = purchaseDatePattern.MatchString(reciept.PurchaseDate)
	if !valid {
		log.Printf(reciept.PurchaseDate)
		return valid
	}

	purchaseTimePattern := regexp.MustCompile(`(0[0-9]|1[0-9]|2[0-4]):[0-5][0-9]`)
	valid = purchaseTimePattern.MatchString(reciept.PurchaseTime)
	if !valid {
		log.Printf(reciept.PurchaseTime)
		return valid
	}

	expensePattern := regexp.MustCompile(`^\d+\.\d{2}$`)
	valid = expensePattern.MatchString(reciept.Total)
	if !valid {
		log.Printf(reciept.Total)

		return valid
	}

	items := reciept.Items
	for _, item := range items {
		valid = textPattern.MatchString(item.ShortDescription)
		if !valid {
			log.Printf(item.ShortDescription)
			return valid
		}

		valid = expensePattern.MatchString(item.Price)
		if !valid {
			log.Printf(item.Price)
			return valid
		}

	}

	return valid
}
