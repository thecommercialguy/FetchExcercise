package main

import (
	"errors"
	"math"
	"net/http"
	"strconv"
	"strings"
	"unicode"
)

// Processes and stores receipts
func (cfg *apiConfig) handlerGetPointsByID(w http.ResponseWriter, r *http.Request) {
	receiptID := r.PathValue("id")

	value, ok := cfg.DB.Load(receiptID)
	err := errors.New("receipt not found for id")
	if !ok {
		respondWithError(w, http.StatusNotFound, "No receipt found for that ID.", err)
		return
	}

	type ResponseBody struct {
		Points int64 `json:"points"`
	}

	receipt := value.(Receipt)

	// Obtaining points awarded by field
	retailerPoints := retailerPoints(receipt.Retailer)
	totalPoints := totalPoints(receipt.Total)
	itemPoints := itemPoints(receipt.Items)
	shortDescriptionPoints := shortDescriptionPoints(receipt.Items)
	purchaseDatePoints := purchaseDatePoints(receipt.PurchaseDate)
	purchaseTimePoints := purchaseTimePoints(receipt.PurchaseTime)

	// Summation of points awarded to Receipt
	receiptPoints := retailerPoints + totalPoints + itemPoints + shortDescriptionPoints + purchaseDatePoints + purchaseTimePoints
	int64Points := int64(receiptPoints)

	respondWithJSON(w, http.StatusOK, ResponseBody{
		Points: int64Points,
	})

}

// Calculates and returns points awarded based off "Retailer" field
func retailerPoints(retailer string) int {
	points := 0

	for _, cha := range retailer {
		if unicode.IsLetter(cha) {
			points++
		}

		if unicode.IsDigit(cha) {
			points++
		}
	}

	return points

}

// Calculates and returns points awarded based off "Total" field
func totalPoints(total string) int {
	points := 0

	decimalIndex := strings.Index(total, ".")
	value := total[decimalIndex+1:]

	if value == "00" {
		points += 50
	}

	valueInt, _ := strconv.Atoi(value)

	if valueInt%25 == 0 {
		points += 25
	}

	return points

}

// Calculates and returns points awarded based off "Items" array field
func itemPoints(items []Item) int {
	points := 0
	numItems := len(items)
	if numItems < 1 {
		return points
	}

	secondItems := math.Floor(float64(numItems) / 2)

	points = int(secondItems * 5)

	return points
}

// Calculates and returns points awarded based off "ShortDescription" field
func shortDescriptionPoints(items []Item) int {
	points := 0

	for _, item := range items {
		description := item.ShortDescription

		descriptionTrimmed := strings.TrimSpace(description)

		if len(descriptionTrimmed)%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)

			toRound := price * .2
			roundedAmount := int(math.Ceil(toRound))

			points += roundedAmount
		}

	}

	return points
}

// Calculates and returns points awarded based off "PurchaseDate" field
func purchaseDatePoints(purchaseDate string) int {
	points := 0
	dateString := purchaseDate[8:10]

	dateInt, _ := strconv.Atoi(dateString)

	if dateInt%2 != 0 {
		points += 6
	}

	return points

}

// Calculates and returns points awarded based off "PurchaseTime" field
func purchaseTimePoints(purchaseTime string) int {
	points := 0
	purchaseTimeSplit := strings.Split(purchaseTime, ":")
	purchaseTimeJoined := strings.Join(purchaseTimeSplit, "")
	purchaseTimeValue, _ := strconv.Atoi(purchaseTimeJoined)

	if 1400 < purchaseTimeValue && purchaseTimeValue < 1600 {
		points += 10
	}

	return points
}
