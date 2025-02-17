package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/google/uuid"
)

// Expecting 200 OK response from handler
func TestHandlerProcessrReceipts_Success(t *testing.T) {
	var database sync.Map

	apiCfg := apiConfig{
		DB: database,
	}

	testReceipt := Receipt{
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

	var b bytes.Buffer

	err := json.NewEncoder(&b).Encode(testReceipt)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodPost, "/receipts/process", &b)

	apiCfg.handlerProcessReceipts(w, req)

	// Expecting 200 OK response from handler
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code\n expected: %v\n actual: %v",
			status, http.StatusOK)
	}

}

// Expecting the "id" key to be present in the response body
func TestHandlerProcessReceipts_ResponseBodyHasKey(t *testing.T) {
	var database sync.Map

	apiCfg := apiConfig{
		DB: database,
	}

	testReceipt := Receipt{
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

	var b bytes.Buffer

	err := json.NewEncoder(&b).Encode(testReceipt)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodPost, "/receipts/process", &b)

	apiCfg.handlerProcessReceipts(w, req)

	resp := w.Result()

	var responseBody map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		t.Errorf("issue decoding resposne body: %v", err)
	}

	// Assert "id" key exists
	_, ok := responseBody["id"].(string)
	if !ok {
		t.Errorf("handler returned did not contain expected key: id\nexpected: %v\nactual: %v", true, ok)
	}
}

// Expecting "id" value in response to have valid UUID syntax
func TestHandlerProcessReceipts_ValidateUUID(t *testing.T) {
	var database sync.Map

	apiCfg := apiConfig{
		DB: database,
	}

	testReceipt := Receipt{
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

	var b bytes.Buffer

	err := json.NewEncoder(&b).Encode(testReceipt)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodPost, "/receipts/process", &b)

	apiCfg.handlerProcessReceipts(w, req)

	resp := w.Result()

	var responseBody map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		t.Errorf("issue decoding resposne body: %v", err)
	}

	testUUID, _ := responseBody["id"].(string)

	// Assert valid UUID structure
	_, err = uuid.Parse(testUUID)
	if err != nil {
		t.Errorf("handler returned invalid uuid response structure: %v", err)
	}

}

// Expecting 400 BadRequest from handler
func TestHandlerProcessReceipts_BadRequest(t *testing.T) {
	var database sync.Map

	apiCfg := apiConfig{
		DB: database,
	}

	testReceipts := []Receipt{
		// Malformed Retailer
		{
			Retailer:     "",
			PurchaseDate: "2024-12-18",
			PurchaseTime: "12:00",
			Items: []Item{
				{
					ShortDescription: "Test Item", Price: "10.00",
				},
			},
			Total: "10.00",
		},
		// Malformed PurchaseDate
		{
			Retailer:     "Test Retailer",
			PurchaseDate: "",
			PurchaseTime: "12:00",
			Items: []Item{
				{
					ShortDescription: "Test Item", Price: "10.00",
				},
			},
			Total: "10.00",
		},
		// Malformed PurchaseTime
		{
			Retailer:     "Test Retailer",
			PurchaseDate: "2024-12-18",
			PurchaseTime: "",
			Items: []Item{
				{
					ShortDescription: "Test Item", Price: "10.00",
				},
			},
			Total: "10.00",
		},
		// Malformed Short Description
		{
			Retailer:     "Test Retailer",
			PurchaseDate: "2024-12-18",
			PurchaseTime: "12:00",
			Items: []Item{
				{
					ShortDescription: "", Price: "10.00",
				},
			},
			Total: "10.00",
		},
		// Malformed Price
		{
			Retailer:     "Test Retailer",
			PurchaseDate: "2024-12-18",
			PurchaseTime: "12:00",
			Items: []Item{
				{
					ShortDescription: "Test Item", Price: "",
				},
			},
			Total: "10.00",
		},
		// Malformed Total
		{
			Retailer:     "Test Retailer",
			PurchaseDate: "2024-12-18",
			PurchaseTime: "12:00",
			Items: []Item{
				{
					ShortDescription: "Test Item", Price: "10.00",
				},
			},
			Total: "",
		},
	}

	for _, testReceipt := range testReceipts {
		var b bytes.Buffer

		err := json.NewEncoder(&b).Encode(testReceipt)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodPost, "/receipts/process", &b)

		apiCfg.handlerProcessReceipts(w, req)

		// Assert 400 BadRequest response from handler
		if status := w.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code.\n excpected: %v\n actual: %v",
				http.StatusOK, status)
		}
	}
}

// Expecting "description" key in error resposne
func TestHandlerProcessReceipts_ErrorResponseKey(t *testing.T) {
	var database sync.Map

	apiCfg := apiConfig{
		DB: database,
	}

	testReceipt := Receipt{
		Retailer:     "",
		PurchaseDate: "2024-12-18",
		PurchaseTime: "12:00",
		Items: []Item{
			{
				ShortDescription: "Test Item", Price: "10.00",
			},
		},
		Total: "10.00",
	}

	var b bytes.Buffer

	err := json.NewEncoder(&b).Encode(testReceipt)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodPost, "/receipts/process", &b)

	apiCfg.handlerProcessReceipts(w, req)

	var responseBody map[string]interface{}
	err = json.NewDecoder(w.Body).Decode(&responseBody)
	if err != nil {
		t.Errorf("issue decoding resposne body: %v", err)
	}

	// Assert that key "description" is present in response body
	_, ok := responseBody["description"]
	if !ok {
		t.Errorf("handler returned did not contain expected key response %v, actual: %v", true, ok)
	}

}

// Expecting "The receipt is invalid." value in error resposne body
func TestHandlerProcessReceipts_ErrorResponseValue(t *testing.T) {
	var database sync.Map

	apiCfg := apiConfig{
		DB: database,
	}

	testReceipt := Receipt{
		Retailer:     "",
		PurchaseDate: "2024-12-18",
		PurchaseTime: "12:00",
		Items: []Item{
			{
				ShortDescription: "Test Item", Price: "10.00",
			},
		},
		Total: "10.00",
	}

	var b bytes.Buffer

	err := json.NewEncoder(&b).Encode(testReceipt)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodPost, "/receipts/process", &b)

	apiCfg.handlerProcessReceipts(w, req)

	var responseBody map[string]interface{}
	err = json.NewDecoder(w.Body).Decode(&responseBody)
	if err != nil {
		t.Errorf("issue decoding resposne body: %v", err)
	}

	// Assert value "The receipt is invalid." to be present in the response body
	errorValue := responseBody["description"]
	expectedValue := "The receipt is invalid."
	if errorValue != expectedValue {
		t.Errorf("handler returned did not contain expected value\nexpected: %v\nactual: %v", expectedValue, errorValue)
	}

}
