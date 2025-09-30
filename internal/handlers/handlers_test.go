package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/gregoryAlvim/gobank/internal/services/mocks"
)

func TestAccountHandler_CreateAccount(t *testing.T) {
	mockService := new(mocks.AccountServiceInterface)
	handler := NewAccountHandler(mockService)

	requestBody := map[string]interface{}{
		"monthly_income": 5000,
		"age":            30,
		"full_name":      "John Doe",
		"phone_number":   "123456789",
		"email":          "john.doe@example.com",
		"category":       "standard",
		"balance":        1000,
	}
	body, _ := json.Marshal(requestBody)

	req, err := http.NewRequest("POST", "/account?type=natural", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	mockService.On("CreateAccount", "natural", mock.Anything).Return(nil)

	handler.CreateAccount(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expectedResponse := `{"message":"Account created successfully"}`
	assert.JSONEq(t, expectedResponse, rr.Body.String())

	mockService.AssertExpectations(t)
}

func TestAccountHandler_GetBalance(t *testing.T) {
	mockService := new(mocks.AccountServiceInterface)
	handler := NewAccountHandler(mockService)

	req, err := http.NewRequest("GET", "/account/1/balance?type=natural", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	vars := map[string]string{
		"id": "1",
	}
	req = mux.SetURLVars(req, vars)

	mockService.On("GetBalance", 1, "natural").Return(123.45, nil)

	handler.GetBalance(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expectedResponse := `{"balance":123.45}`
	assert.JSONEq(t, expectedResponse, rr.Body.String())

	mockService.AssertExpectations(t)
}

func TestAccountHandler_Deposit(t *testing.T) {
	mockService := new(mocks.AccountServiceInterface)
	handler := NewAccountHandler(mockService)

	requestBody := map[string]interface{}{"amount": 100}
	body, _ := json.Marshal(requestBody)

	req, err := http.NewRequest("POST", "/account/1/deposit?type=natural", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	vars := map[string]string{
		"id": "1",
	}
	req = mux.SetURLVars(req, vars)

	mockService.On("Deposit", 1, 100.0, "natural").Return(nil)

	handler.Deposit(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expectedResponse := `{"message":"Deposit successful"}`
	assert.JSONEq(t, expectedResponse, rr.Body.String())

	mockService.AssertExpectations(t)
}

func TestAccountHandler_Withdraw(t *testing.T) {
	mockService := new(mocks.AccountServiceInterface)
	handler := NewAccountHandler(mockService)

	requestBody := map[string]interface{}{"amount": 50}
	body, _ := json.Marshal(requestBody)

	req, err := http.NewRequest("POST", "/account/1/withdraw?type=natural", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	vars := map[string]string{
		"id": "1",
	}
	req = mux.SetURLVars(req, vars)

	mockService.On("Withdraw", 1, 50.0, "natural").Return(nil)

	handler.Withdraw(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expectedResponse := `{"message":"Withdrawal successful"}`
	assert.JSONEq(t, expectedResponse, rr.Body.String())

	mockService.AssertExpectations(t)
}

func TestAccountHandler_Transfer(t *testing.T) {
	mockService := new(mocks.AccountServiceInterface)
	handler := NewAccountHandler(mockService)

	requestBody := map[string]interface{}{
		"from_id":   1,
		"to_id":     2,
		"from_type": "natural",
		"to_type":   "legal",
		"amount":    100,
	}
	body, _ := json.Marshal(requestBody)

	req, err := http.NewRequest("POST", "/account/transfer", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	mockService.On("Transfer", 1, 2, 100.0, "natural", "legal").Return(nil)

	handler.Transfer(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expectedResponse := `{"message":"Transfer successful"}`
	assert.JSONEq(t, expectedResponse, rr.Body.String())

	mockService.AssertExpectations(t)
}

func TestAccountHandler_CloseAccount(t *testing.T) {
	mockService := new(mocks.AccountServiceInterface)
	handler := NewAccountHandler(mockService)

	req, err := http.NewRequest("DELETE", "/account/1?type=natural", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	vars := map[string]string{
		"id": "1",
	}
	req = mux.SetURLVars(req, vars)

	mockService.On("CloseAccount", 1, "natural").Return(nil)

	handler.CloseAccount(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expectedResponse := `{"message":"Account closed successfully"}`
	assert.JSONEq(t, expectedResponse, rr.Body.String())

	mockService.AssertExpectations(t)
}
