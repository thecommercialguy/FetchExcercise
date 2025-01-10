package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

// Expecting 200 OK response from handler
func TestHandlerGetPoints_Success(t *testing.T) {
	var database sync.Map

	apiCfg := apiConfig{
		DB: database,
	}

	testReceiptID := "00000000-0000-0000-0000-000000000000"

	testReceipt := Reciept{
		ID:           testReceiptID,
		Retailer:     "Test Retailer",
		PurchaseDate: "2024-12-18",
		PurchaseTime: "12:00",
		Items: []Item{
			{
				ShortDescription: "Test Item", Price: "10.00",
			},
		},
		Total: "10.00",
	}

	apiCfg.DB.Store(testReceiptID, testReceipt)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /reciepts/{id}/points", apiCfg.handlerGetPointsByID)

	path := "/reciepts/" + testReceiptID + "/points"
	req := httptest.NewRequest(http.MethodGet, path, nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	// Assert 200 OK response from handler
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code.\n excpected: %v\n actual: %v",
			http.StatusOK, status)
	}

}

// Expecting the "id" key to be present in the response
func TestHandlerGetPoints_ResponseBodyHasKey(t *testing.T) {
	var database sync.Map

	apiCfg := apiConfig{
		DB: database,
	}

	testReceiptID := "00000000-0000-0000-0000-000000000000"

	testReceipt := Reciept{
		ID:           testReceiptID,
		Retailer:     "Test Retailer",
		PurchaseDate: "2024-12-18",
		PurchaseTime: "12:00",
		Items: []Item{
			{
				ShortDescription: "Test Item", Price: "10.00",
			},
		},
		Total: "10.00",
	}

	apiCfg.DB.Store(testReceiptID, testReceipt)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /reciepts/{id}/points", apiCfg.handlerGetPointsByID)

	path := "/reciepts/" + testReceiptID + "/points"
	req := httptest.NewRequest(http.MethodGet, path, nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	var responseBody map[string]int64
	err := json.NewDecoder(w.Body).Decode(&responseBody)
	if err != nil {
		t.Errorf("issue decoding resposne body: %v", err)
	}

	// Assert "points" key exists
	_, ok := responseBody["points"]
	if !ok {
		t.Errorf("handler returned did not contain expected key response %v, actual: %v", true, ok)
	}

}

// Expecting "points" value in response to equal manually calculated expected value
func TestHandlerGetPoints_ValidatePoints(t *testing.T) {
	var database sync.Map

	apiCfg := apiConfig{
		DB: database,
	}

	testReceiptID := "00000000-0000-0000-0000-000000000000"

	testReceipt := Reciept{
		ID:           testReceiptID,
		Retailer:     "Test Retailer",
		PurchaseDate: "2024-12-18",
		PurchaseTime: "12:00",
		Items: []Item{
			{
				ShortDescription: "Test Item", Price: "10.00",
			},
		},
		Total: "10.00",
	}

	apiCfg.DB.Store(testReceiptID, testReceipt)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /reciepts/{id}/points", apiCfg.handlerGetPointsByID)

	path := "/reciepts/" + testReceiptID + "/points"
	req := httptest.NewRequest(http.MethodGet, path, nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	var responseBody map[string]int64

	err := json.NewDecoder(w.Body).Decode(&responseBody)
	if err != nil {
		t.Errorf("issue decoding resposne body: %v", err)
	}

	expectedPoints := int64(89)

	// Assert response's points value to equal expected value
	actualPoints, _ := responseBody["points"]
	if expectedPoints != actualPoints {
		t.Errorf("handler returned incorrect number of points, %v, expected: %v", actualPoints, expectedPoints)
	}

}

// Expecting 404 NotFound from handler
func TestHandlerGetPoints_NotFound(t *testing.T) {
	var database sync.Map

	apiCfg := apiConfig{
		DB: database,
	}

	testReceiptID := "00000000-0000-0000-0000-000000000000"
	dummyRecieptID := "10000000-2000-3000-4000-500000000000"

	testReceipt := Reciept{
		ID:           testReceiptID,
		Retailer:     "Test Retailer",
		PurchaseDate: "2024-12-18",
		PurchaseTime: "12:00",
		Items: []Item{
			{
				ShortDescription: "Test Item", Price: "10.00",
			},
		},
		Total: "10.00",
	}

	apiCfg.DB.Store(testReceiptID, testReceipt)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /reciepts/{id}/points", apiCfg.handlerGetPointsByID)

	path := "/reciepts/" + dummyRecieptID + "/points"
	req := httptest.NewRequest(http.MethodGet, path, nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	// Assert 404 NotFound response from handler
	if status := w.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code.\n excpected: %v\n actual: %v",
			http.StatusOK, status)
	}

}

// Expecting "description" key in error resposne body
func TestHandlerGetPoints_ErrorResponseKey(t *testing.T) {
	var database sync.Map

	apiCfg := apiConfig{
		DB: database,
	}

	testReceiptID := "00000000-0000-0000-0000-000000000000"
	dummyRecieptID := "10000000-2000-3000-4000-500000000000"

	testReceipt := Reciept{
		ID:           testReceiptID,
		Retailer:     "Test Retailer",
		PurchaseDate: "2024-12-18",
		PurchaseTime: "12:00",
		Items: []Item{
			{
				ShortDescription: "Test Item", Price: "10.00",
			},
		},
		Total: "10.00",
	}

	apiCfg.DB.Store(testReceiptID, testReceipt)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /reciepts/{id}/points", apiCfg.handlerGetPointsByID)

	path := "/reciepts/" + dummyRecieptID + "/points"
	req := httptest.NewRequest(http.MethodGet, path, nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	var responseBody map[string]interface{}
	err := json.NewDecoder(w.Body).Decode(&responseBody)
	if err != nil {
		t.Errorf("issue decoding resposne body: %v", err)
	}

	// Assert that key "description" is present in response body
	_, ok := responseBody["description"]
	if !ok {
		t.Errorf("handler returned did not contain expected key response %v, actual: %v", true, ok)
	}

}

// Expecting "No receipt found for that ID." value in error resposne body
func TestHandlerGetPoints_ErrorResponseValue(t *testing.T) {
	var database sync.Map

	apiCfg := apiConfig{
		DB: database,
	}

	testReceiptID := "00000000-0000-0000-0000-000000000000"
	dummyRecieptID := "10000000-2000-3000-4000-500000000000"

	testReceipt := Reciept{
		ID:           testReceiptID,
		Retailer:     "Test Retailer",
		PurchaseDate: "2024-12-18",
		PurchaseTime: "12:00",
		Items: []Item{
			{
				ShortDescription: "Test Item", Price: "10.00",
			},
		},
		Total: "10.00",
	}

	apiCfg.DB.Store(testReceiptID, testReceipt)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /reciepts/{id}/points", apiCfg.handlerGetPointsByID)

	path := "/reciepts/" + dummyRecieptID + "/points"
	req := httptest.NewRequest(http.MethodGet, path, nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	var responseBody map[string]interface{}
	err := json.NewDecoder(w.Body).Decode(&responseBody)
	if err != nil {
		t.Errorf("issue decoding resposne body: %v", err)
	}

	// Assert value "No receipt found for that ID." to be present in the response body
	errorValue := responseBody["description"]
	expectedValue := "No receipt found for that ID."
	if errorValue != expectedValue {
		t.Errorf("handler returned did not contain expected value\nexpected: %v\nactual: %v", expectedValue, errorValue)
	}

}
