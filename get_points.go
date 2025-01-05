package main

import (
	"errors"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"unicode"
)

func (cfg *apiConfig) handlerGetPointsByID(w http.ResponseWriter, r *http.Request) {
	recieptID := r.PathValue("id")

	value, ok := cfg.DB.Load(recieptID)
	err := errors.New("receipt not found for id")
	if !ok {
		respondWithError(w, http.StatusNotFound, "No receipt found for that ID.", err)
		return
	}

	reciept := value.(Reciept)
	retailerPoints := retailerPoints(reciept.Retailer)
	log.Printf("Retailer: %v", retailerPoints)
	totalPoints := totalPoints(reciept.Total)
	log.Printf("Total: %v", totalPoints)
	itemPoints := itemPoints(reciept.Items)
	log.Printf("Item: %v", itemPoints)
	descriptionPoints := descriptionPoints(reciept.Items)
	log.Printf("Description: %v", descriptionPoints)
	datePoints := datePoints(reciept.PurchaseDate)
	log.Printf("Date: %v", datePoints)
	timePoints := timePoints(reciept.PurchaseTime)
	log.Printf("Time: %v", timePoints)
	recieptPoints := retailerPoints + totalPoints + itemPoints + descriptionPoints + datePoints + timePoints
	pointsString := strconv.Itoa(recieptPoints)

	respondWithJSON(w, http.StatusOK, pointsString)

}

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

func descriptionPoints(items []Item) int {
	points := 0

	for _, item := range items {
		description := item.ShortDescription
		// price := item.Price

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

func datePoints(purchaseDate string) int {
	points := 0
	dateString := purchaseDate[8:10]

	dateInt, _ := strconv.Atoi(dateString)

	if dateInt%2 != 0 {
		points += 6
	}

	return points

}

func timePoints(purchaseTime string) int {
	points := 0
	purchaseTimeSplit := strings.Split(purchaseTime, ":")
	purchaseTimeJoined := strings.Join(purchaseTimeSplit, "")
	purchaseTimeValue, _ := strconv.Atoi(purchaseTimeJoined)

	if 1400 < purchaseTimeValue && purchaseTimeValue < 1600 {
		points += 10
	}

	return points
}
