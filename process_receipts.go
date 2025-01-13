package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"

	"github.com/google/uuid"
)

type Receipt struct {
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

// Determines and returns points awarded to a receipt
func (cfg *apiConfig) handlerProcessReceipts(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Retailer     string `json:"retailer"`
		PurchaseDate string `json:"purchaseDate"`
		PurchaseTime string `json:"purchaseTime"`
		Items        []Item `json:"items"`
		Total        string `json:"total"`
	}

	// Decode JSON request body into Go readable struct
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "The receipt is invalid.", err)
		return
	}

	// Generate new UUID using "github.com/google/uuid"
	newUUID := uuid.New()
	uuidString := newUUID.String()

	// Store params data in a new Receipt object.
	newReceipt := Receipt{
		ID:           uuidString,
		Retailer:     params.Retailer,
		PurchaseDate: params.PurchaseDate,
		PurchaseTime: params.PurchaseTime,
		Items:        params.Items,
		Total:        params.Total,
	}

	// Validates "Receipt" fields
	valid := validateReceipt(newReceipt)
	if !valid {
		respondWithError(w, http.StatusBadRequest, "The receipt is invalid.", err)
		return
	}

	// Structure of JSON response body
	type ResponseBody struct {
		Id string `json:"id"`
	}

	// Store newly validated Receipt in DB (sync.Map), using UUID generated as the key
	cfg.DB.Store(uuidString, newReceipt)

	respondWithJSON(w, http.StatusOK, ResponseBody{
		Id: uuidString,
	})

}

// Ensures each Receipt field conforms to expected patterns
// -> true If valid
// -> false If invalid
func validateReceipt(receipt Receipt) bool {
	textPattern := regexp.MustCompile(`^[\w\s&-]+$`)
	valid := textPattern.MatchString(receipt.Retailer)
	if !valid {
		log.Printf("Malformed Retailer field: %v", receipt.Retailer)
		return valid
	}

	purchaseDatePattern := regexp.MustCompile(`\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01])`)
	valid = purchaseDatePattern.MatchString(receipt.PurchaseDate)
	if !valid {
		log.Printf("Malformed PurchaseDate field: %v", receipt.PurchaseDate)
		return valid
	}

	purchaseTimePattern := regexp.MustCompile(`(0[0-9]|1[0-9]|2[0-4]):[0-5][0-9]`)
	valid = purchaseTimePattern.MatchString(receipt.PurchaseTime)
	if !valid {
		log.Printf("Malformed PurchaseTime field: %v", receipt.PurchaseTime)
		return valid
	}

	expensePattern := regexp.MustCompile(`^\d+\.\d{2}$`)
	valid = expensePattern.MatchString(receipt.Total)
	if !valid {
		log.Printf("Malformed Total field: %v", receipt.Total)
		return valid
	}

	items := receipt.Items
	for _, item := range items {
		valid = textPattern.MatchString(item.ShortDescription)
		if !valid {
			log.Printf("Malformed ShortDescription field: %v", item.ShortDescription)
			return valid
		}

		valid = expensePattern.MatchString(item.Price)
		if !valid {
			log.Printf("Malformed Price field: %v", item.Price)
			return valid
		}

	}

	return valid
}
